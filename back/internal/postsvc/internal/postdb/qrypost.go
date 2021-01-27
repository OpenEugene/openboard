package postdb

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

func (s *PostDB) upsertType(ctx cx, sid string, x *pb.AddTypeReq, y *pb.TypeResp) error {
	id, ok := parseOrUID(s.ug, sid)
	if !ok {
		return fmt.Errorf("invalid uid")
	}

	_, err := s.db.Exec("INSERT INTO `type` (type_id, name) VALUES (?, ?) ON DUPLICATE KEY UPDATE type_id = ?, name = ?", &id, x.Name, &id, x.Name)
	if err != nil {
		return err
	}

	y.Id = id.String()
	e := s.db.QueryRow("SELECT name FROM `type` where type_id = ?", y.Id).Scan(&y.Name)
	if e != nil {
		return e
	}

	return nil
}

func (s *PostDB) findTypes(ctx cx, x *pb.FndTypesReq, y *pb.TypesResp) error {
	selStmt, err := s.db.Prepare("SELECT type_id, name FROM type LIMIT ? OFFSET ?")
	if err != nil {
		return err
	}
	defer selStmt.Close()

	rows, err := selStmt.Query(x.Limit, x.Lapse)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		r := pb.TypeResp{}

		err := rows.Scan(&r.Id, &r.Name)
		if err != nil {
			return err
		}

		y.Items = append(y.Items, &r)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	err = s.db.QueryRow("SELECT COUNT(*) FROM type").Scan(&y.Total)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostDB) upsertPost(ctx cx, sid string, x *pb.AddPostReq, y *pb.PostResp) error {
	id, ok := parseOrUID(s.ug, sid)
	if !ok {
		return fmt.Errorf("invalid uid")
	}

	_, err := s.db.Exec(`
		INSERT INTO post (post_id, type_id, slug, title, body)
		VALUES (?, ?, ?, ?, ?) 
		ON DUPLICATE KEY UPDATE post_id = ?, type_id = ?, slug = ?, title = ?, body = ?`,
		&id, x.TypeId, "", x.Title, x.Body, &id, x.TypeId, "", x.Title, x.Body)

	if err != nil {
		return err
	}

	y.Id = id.String()
	s.db.QueryRow("SELECT type_id, title, body FROM post WHERE post_id = ?", y.Id).Scan(&y.TypeId, &y.Title, &y.Body)

	return nil
}

// TODO: make it such that if given a list of multiple keywords, we can search the
// title and body for those keywords.
func (s *PostDB) findPosts(ctx cx, x *pb.FndPostsReq, y *pb.PostsResp) error {
	selStmt, err := s.db.Prepare(`
		SELECT post_id, type_id, slug, title, body, 
			created_at, updated_at, deleted_at 
		FROM post WHERE title LIKE ? OR body LIKE ?
		LIMIT ? OFFSET ?
	`)
	if err != nil {
		return err
	}
	defer selStmt.Close()

	rows, err := selStmt.Query(
		"%"+x.Keywords[0]+"%",
		"%"+x.Keywords[0]+"%",
		x.Limit,
		x.Lapse,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		r := pb.PostResp{}
		var tc, tu, td, tb mysql.NullTime

		err := rows.Scan(&r.Id, &r.TypeId, &r.Slug, &r.Title, &r.Body, &tc, &tu, &td)
		if err != nil {
			return err
		}

		r.Created = asTS(tc.Time, tc.Valid)
		r.Updated = asTS(tu.Time, tu.Valid)
		r.Deleted = asTS(td.Time, td.Valid)
		r.Blocked = asTS(tb.Time, tb.Valid)

		y.Posts = append(y.Posts, &r)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	err = s.db.QueryRow(
		"SELECT COUNT(*) FROM post WHERE title LIKE ? OR body LIKE ?",
		"%"+x.Keywords[0]+"%",
		"%"+x.Keywords[0]+"%",
	).Scan(&y.Total)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostDB) deletePost(ctx cx, sid string) error {
	_, err := s.db.Exec(
		"UPDATE post SET deleted_at = ? WHERE post_id = ?",
		time.Now(),
		sid,
	)
	if err != nil {
		return err
	}

	return nil
}
