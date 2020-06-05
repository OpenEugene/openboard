package userdb

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenEugene/openboard/back/internal/altr"
	"github.com/OpenEugene/openboard/back/internal/pb"
	"github.com/codemodus/uidgen"
	"github.com/go-sql-driver/mysql"
)

type cx = context.Context

var (
	csvStr = altr.CSVFromStrings
	lim    = altr.LimitUint32
	asTS   = altr.Timestamp
)

func parseOrUID(ug *uidgen.UIDGen, sid string) (uidgen.UID, bool) {
	if sid == "" {
		return ug.UID(), true
	}
	return ug.Parse(sid)
}

func (s *UserDB) upsertUser(ctx cx, sid string, x *pb.AddUserReq, y *pb.UserResp) error {
	id, ok := parseOrUID(s.ug, sid)
	if !ok {
		return fmt.Errorf("invalid uid")
	}

	// todo: be able to link roleIDs to users.
	stmt, err := s.db.Prepare("INSERT INTO user (user_id, username, email, email_hold, altmail, altmail_hold, full_name, avatar, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE user_id = ?, username = ?, email = ?, email_hold = ?, altmail = ?, altmail_hold = ?, full_name = ?, avatar = ?, password = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		&id,
		x.Username,
		x.Email,
		x.EmailHold,
		x.Altmail,
		x.AltmailHold,
		x.FullName,
		x.Avatar,
		x.Password,
		&id,
		x.Username,
		x.Email,
		x.EmailHold,
		x.Altmail,
		x.AltmailHold,
		x.FullName,
		x.Avatar,
		x.Password,
	)
	if err != nil {
		return err
	}

	r := pb.User{}

	r.Id = id.String()
	r.Username = x.Username
	r.Email = x.Email
	r.EmailHold = x.EmailHold
	r.Altmail = x.Altmail
	r.AltmailHold = x.AltmailHold
	r.FullName = x.FullName
	r.Avatar = x.Avatar

	y.Item = &r
	// todo: respond with user roles

	return nil
}

func (s *UserDB) deleteUser(ctx cx, sid string) error {
	_, err := s.db.Exec(
		"UPDATE user SET deleted_at = ? WHERE user_id = ?",
		time.Now(),
		sid,
	)
	if err != nil {
		return err
	}

	return nil
}

type userAndRole struct {
	rid, uid, username, email, altmail, fullName, avatar, rolename string
	emailHold, altmailHold                                         bool
	rids                                                           []string
	tl, tc, tu, td, tb                                             mysql.NullTime
}

func (s *UserDB) findUsers(ctx cx, x *pb.FndUsersReq, y *pb.UsersResp) error {
	qry := "SELECT u.user_id, u.username, u.email, u.email_hold, u.altmail, "
	qry += "u.altmail_hold, u.full_name, u.avatar, r.role_id, r.role_name, u.last_login, u.created_at, "
	qry += "u.updated_at, u.deleted_at, u.blocked_at "
	qry += "FROM user u "
	qry += "LEFT JOIN user_role ur ON u.user_id = ur.user_id "
	qry += "LEFT JOIN role r ON r.role_id = ur.role_id "
	qry += "WHERE u.email = ? AND u.email_hold = ? LIMIT ? OFFSET ?"

	rows, err := s.db.Query(qry, x.Email, x.EmailHold, x.Limit, x.Lapse)
	if err != nil {
		return err
	}
	defer rows.Close()

	var tl, tc, tu, td, tb mysql.NullTime
	urs := []userAndRole{}

	for rows.Next() {
		ur := userAndRole{}

		err := rows.Scan(
			&ur.uid,
			&ur.username,
			&ur.email,
			&ur.emailHold,
			&ur.altmail,
			&ur.altmailHold,
			&ur.fullName,
			&ur.avatar,
			&ur.rid,
			&ur.rolename,
			&tl,
			&tc,
			&tu,
			&td,
			&tb,
		)
		if err != nil {
			return err
		}

		urs = append(urs, ur)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	userRespMap := condenseUsersAndRoles(urs)

	for _, resp := range userRespMap {
		y.Items = append(y.Items, resp)
	}

	err = s.db.QueryRow(
		"SELECT COUNT(*) FROM user WHERE email = ? AND email_hold = ?",
		x.Email,
		x.EmailHold,
	).Scan(&y.Total)
	if err != nil {
		return err
	}

	return nil
}

type userRespMap map[string]*pb.User

// concenseUsersAndRoles takes a list of repeated users (with different roles)
// and condenses them into a map, putting their associated roles into a slice in a struct.
func condenseUsersAndRoles(urs []userAndRole) userRespMap {
	urm := make(userRespMap)
	user := pb.User{}
	role := pb.RoleResp{}

	// Collect roles for each user
	for _, ur := range urs {
		// This set of data only needed for new user keys in map.
		if _, ok := urm[user.Id]; !ok {
			urm[user.Id].Id = ur.uid
			urm[user.Id].Username = ur.username
			urm[user.Id].Email = ur.email
			urm[user.Id].EmailHold = ur.emailHold
			urm[user.Id].Altmail = ur.altmail
			urm[user.Id].AltmailHold = ur.altmailHold
			urm[user.Id].FullName = ur.fullName
			urm[user.Id].Avatar = ur.avatar
			urm[user.Id].LastLogin = asTS(ur.tl.Time, ur.tl.Valid)
			urm[user.Id].Created = asTS(ur.tc.Time, ur.tc.Valid)
			urm[user.Id].Updated = asTS(ur.tu.Time, ur.tu.Valid)
			urm[user.Id].Deleted = asTS(ur.td.Time, ur.td.Valid)
			urm[user.Id].Blocked = asTS(ur.tb.Time, ur.tb.Valid)
		}

		// The role data can only be set if user is a key in the map.
		if _, ok := urm[user.Id]; ok && role.Id != "" {
			role.Id = ur.rid
			role.Name = ur.rolename
			urm[user.Id].Roles = append(urm[user.Id].Roles, &role)
		}
	}

	return urm
}

func (s *UserDB) upsertRole(ctx cx, sid string, x *pb.AddRoleReq, y *pb.RoleResp) error {
	id, ok := parseOrUID(s.ug, sid)
	if !ok {
		return fmt.Errorf("invalid uid")
	}

	stmt, err := s.db.Prepare("INSERT INTO role (role_id, role_name) VALUES (?, ?) ON DUPLICATE KEY UPDATE role_id = ?, role_name = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&id, x.Name, &id, x.Name)
	if err != nil {
		return err
	}

	y.Id = id.String()
	y.Name = x.Name

	return nil
}

func (s *UserDB) findRoles(ctx cx, x *pb.FndRolesReq, y *pb.RolesResp) error {
	selStmt, err := s.db.Prepare("SELECT role_id, role_name FROM role WHERE role_id = ? OR role_name = ? LIMIT ? OFFSET ?")
	if err != nil {
		return err
	}
	defer selStmt.Close()

	var roleIDs, roleNames string

	// TODO: enable search of more than one role ID
	if len(x.RoleIds) > 0 {
		roleIDs = x.RoleIds[0]
	}
	// TODO: enable search of more than one role name
	if len(x.RoleNames) > 0 {
		roleNames = x.RoleNames[0]
	}

	rows, err := selStmt.Query(roleIDs, roleNames, x.Limit, x.Lapse)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		r := pb.RoleResp{}

		err := rows.Scan(
			&r.Id, &r.Name,
		)
		if err != nil {
			return err
		}

		y.Items = append(y.Items, &r)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	err = s.db.QueryRow(
		"SELECT COUNT(*) FROM role WHERE role_id = ? OR role_name = ? LIMIT ? OFFSET ?",
		roleIDs,
		roleNames,
		x.Limit,
		x.Lapse,
	).Scan(&y.Total)
	if err != nil {
		return err
	}

	return nil
}
