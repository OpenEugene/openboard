package userdb

import (
	"context"
	"fmt"

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
	stmt, err := s.db.Prepare("INSERT INTO user (user_id, username, email, email_hold, altmail, altmail_hold, full_name, avatar, password) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE user_id = ?, username = ?, email = ?, email_hold = ?, altmail = ?, altmail_hold = ?, full_name = ?, avatar = ?, password = ?")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(
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
	); err != nil {
		return err
	}

	item := pb.User{
		Id:          123,
		Username:    x.Username,
		Email:       x.Email,
		EmailHold:   x.EmailHold,
		Altmail:     x.Altmail,
		AltmailHold: x.AltmailHold,
		FullName:    x.FullName,
		Avatar:      x.Avatar,
	}
	y.Item = &item
	// todo: respond with user roles

	return nil
}

func (s *UserDB) deleteUser(ctx cx, sid string) error {
	stmt, err := s.db.Prepare("DELETE FROM user WHERE user_id = ?")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(sid); err != nil {
		return err
	}

	return nil
}

func (s *UserDB) findUsers(ctx cx, x *pb.FndUsersReq, y *pb.UsersResp) error {
	selStmt, err := s.db.Prepare("SELECT user_id, username, email, email_hold, altmail, altmail_hold, full_name, avatar, last_login, created_at, updated_at, deleted_at, blocked_at FROM user WHERE (email = ? AND email_hold = ?) OR (altmail = ? AND altmail_hold = ?)")
	if err != nil {
		return err
	}
	defer selStmt.Close()

	rows, err := selStmt.Query(
		x.Email,
		x.EmailHold,
		x.Altmail,
		x.AltmailHold,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		r := pb.User{}

		var tl, tc, tu, td, tb mysql.NullTime

		if err := rows.Scan(
			&r.Id, &r.Username, &r.Email, &r.EmailHold, &r.Altmail, &r.AltmailHold, &r.FullName, &r.Avatar, &tl, &tc, &tu, &td, &tb,
		); err != nil {
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

	stmt, err := s.db.Prepare("INSERT INTO role (role_id, role_name) VALUES(?, ?) ON DUPLICATE KEY UPDATE role_id = ?, role_name = ?")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(&id, x.Name, &id, x.Name); err != nil {
		return err
	}

	y.Id = 456
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

	if err = s.db.QueryRow("SELECT COUNT(*) FROM role WHERE role_id in ? OR role_name in = ? LIMIT ? OFFSET ?").Scan(&y.Total); err != nil {
		return err
	}

	return nil
}
