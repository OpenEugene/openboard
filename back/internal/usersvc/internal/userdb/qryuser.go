package userdb

import (
	"context"
	"fmt"
	"strconv"

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
	stmt, err := s.db.Prepare("INSERT INTO users(user_id, username, email, email_hold, altmail, altmail_hold, full_name, avatar, password) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE user_id = ?, username = ?, email = ?, email_hold = ?, altmail = ?, altmail_hold = ?, full_name = ?, avatar = ?, password = ?")

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
	)

	if err != nil {
		return err
	}

	var intID int
	intID, err = strconv.Atoi(id.String())
	if err != nil {
		return err
	}

	y.Item.Id = uint32(intID)
	y.Item.Username = x.Username
	y.Item.Email = x.Email
	y.Item.EmailHold = x.EmailHold
	y.Item.Altmail = x.Altmail
	y.Item.AltmailHold = x.AltmailHold
	y.Item.FullName = x.FullName
	y.Item.Avatar = x.Avatar
	// todo: respond with user roles

	return nil
}

func (s *UserDB) deleteUser(ctx cx, sid string) error {
	stmt, err := s.db.Prepare("DELETE FROM users WHERE user_id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(sid)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserDB) findUsers(ctx cx, x *pb.FndUsersReq, y *pb.UsersResp) error {
	selStmt, err := s.db.Prepare("SELECT user_id, username, email, email_hold, altmail, altmail_hold, full_name, avatar, last_login, created_at, updated_at, deleted_at, blocked_at FROM user WHERE email IN ? OR email_hold = ? OR altmail in ? OR altmail_hold = ? LIMIT ? OFFSET ?")

	if err != nil {
		return err
	}
	defer selStmt.Close()
	rows, err := selStmt.Query(
		x.Emails,
		x.EmailHold,
		x.Altmails,
		x.AltmailHold,
		x.Limit,
		x.Lapse,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		r := pb.User{}

		var tl, tc, tu, td, tb mysql.NullTime
		err := rows.Scan(
			&r.Id, &r.Username, &r.Email, &r.EmailHold, &r.Altmail, &r.AltmailHold, &r.FullName, &r.Avatar, &tl, &tc, &tu, &td, &tb,
		)

		if err != nil {
			return err
		}

		// TODO: retrieve users by roleIDs, which doesn't have a table yet.
		r.LastLogin = asTS(tl.Time, tl.Valid)
		r.Created = asTS(tc.Time, tc.Valid)
		r.Updated = asTS(tu.Time, tu.Valid)
		r.Deleted = asTS(td.Time, td.Valid)
		r.Blocked = asTS(tb.Time, tb.Valid)

		y.Items = append(y.Items, &r)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	err = s.db.QueryRow("SELECT COUNT(*) FROM user WHERE email = ? OR email_hold = ? OR altmail = ? OR altmail_hold = ?").Scan(&y.Total)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserDB) upsertRole(ctx cx, sid string, x *pb.AddRoleReq, y *pb.RoleResp) error {
	id, ok := parseOrUID(s.ug, sid)

	if !ok {
		return fmt.Errorf("invalid uid")
	}

	stmt, err := s.db.Prepare("INSERT INTO roles(role_id, role_name) VALUES(?, ?) ON DUPLICATE KEY UPDATE role_id = ?, role_name = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, x.Name)

	if err != nil {
		return err
	}

	var intID int
	intID, err = strconv.Atoi(id.String())
	if err != nil {
		return err
	}
	y.Id = uint32(intID)
	y.Name = x.Name

	return nil
}

func (s *UserDB) findRoles(ctx cx, x *pb.FndRolesReq, y *pb.RolesResp) error {
	selStmt, err := s.db.Prepare("SELECT role_id, role_name FROM role WHERE role_id in ? OR role_name in = ? LIMIT ? OFFSET ?")

	if err != nil {
		return err
	}
	defer selStmt.Close()
	rows, err := selStmt.Query(x.RoleIds, x.RoleNames, x.Limit, x.Lapse)
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

	err = s.db.QueryRow("SELECT COUNT(*) FROM role WHERE role_id in ? OR role_name in = ? LIMIT ? OFFSET ?").Scan(&y.Total)

	if err != nil {
		return err
	}

	return nil
}
