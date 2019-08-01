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

	y.id = id
	y.username = x.username
	y.email = x.email
	y.emailHold = x.emailHold
	y.altmail = x.altmail
	y.altmailHold = x.altmailHold
	y.firstName = x.firstName
	y.lastName = x.lastName
	y.avatar = x.avatar
	y.roles = x.roles

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
	stmt, err := db.Prepare("SELECT user_id, username, email, email_hold, 
	altmail, altmail_hold, first_name, last_name, avatar, last_login, created_at,
	updated_at, deleted_at, blocked_at WHERE ")

	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {

	}
	if err = rows.Err(); err != nil {
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
