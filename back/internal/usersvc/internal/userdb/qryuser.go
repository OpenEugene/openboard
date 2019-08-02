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

	stmt, err := s.db.Prepare("INSERT INTO users(user_id, username, email, email_hold, altmail, altmail_hold, first_name, last_name, avatar, password) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE user_id = ?, username = ?, email = ?, email_hold = ?, altmail = ?, altmail_hold = ?, first_name = ?, last_name = ?, avatar = ?, password = ?")

	if err != nil {
		return err
	}

	_, err := stmt.Exec(id, x.username, x.email, x.emailHold, x.altmail, x.altmailHold, x.firstName, x.lastName, x.avatar, x.password)

	if err != nil {
		return err
	}

	y.Id = id
	y.Username = x.username
	y.Email = x.email
	y.emailHold = x.emailHold
	y.Altmail = x.altmail
	y.AltmailHold = x.altmailHold
	y.FirstName = x.firstName
	y.LastName = x.lastName
	y.Avatar = x.avatar
	y.Roles = x.roles

	return nil
}

func (s *UserDB) deleteUser(ctx cx, sid string) error {
	stmt, err := s.db.Prepare("DELETE FROM users WHERE user_id = ?")
	if err != nil {
		return err
	}

	_, err := stmt.Exec(sid)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserDB) findUsers(ctx cx, x *pb.FndUsersReq, y *pb.UsersResp) error {
	selStmt, err := db.Prepare("SELECT user_id, username, email, email_hold, 
		altmail, altmail_hold, first_name, last_name, avatar, last_login, created_at,
		updated_at, deleted_at, blocked_at FROM user WHERE email = ? OR email_hold = ?
		OR altmail = ? OR altmail_hold = ?")

	if err != nil {
		return err
	}
	defer selStmt.Close()
	rows, err := selStmt.Query(x.email, x.emailHold, x.altmail, x.altmailHold)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		r := pb.UserResp{}

		var tl, tc, tu, td, tb mysql.NullTime
		err := rows.Scan(
			&r.Id, &r.Username, &r.Email, &r.EmailHold, &r.Altmail, &r.AltmailHold,
			&r.FirstName, &r.LastName, &r.Avatar, &tl, &tc, &tu, &td, &tb,
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

	err = db.QueryRow("SELECT COUNT(*) FROM user WHERE email = ? OR email_hold = ?
		OR altmail = ? OR altmail_hold = ?").Scan(&y.Total)

	if err != nil {
		return err
	}
}

func (s *UserDB) upsertRole(ctx cx, sid string, x *pb.AddRoleReq, y *pb.RoleResp) error {
	id, ok := parseOrUID(s.ug, sid)

	if !ok {
		return fmt.Errorf("invalid uid")
	}
}

func (s *UserDB) findRole(ctx cx, x *pb.FndRolesReq, y *pb.RolesResp) error {

}
