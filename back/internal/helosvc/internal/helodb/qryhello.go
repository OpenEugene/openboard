package helodb

import (
	"context"
	"fmt"

	"github.com/champagneabuelo/openboard/back/internal/altr"
	"github.com/champagneabuelo/openboard/back/internal/pb"
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

func (s *HeloDB) upsertHello(ctx cx, sid string, x *pb.AddHelloReq, y *pb.HelloResp) error {
	id, ok := parseOrUID(s.ug, sid)
	if !ok {
		return fmt.Errorf("invalid uid")
	}

	tx := s.db.Begin()
	defer tx.Rollback()

	setStmt := tx.Prepare(`
		SET @id = ?, @sal = ?, @sub = ?
	`)
	defer setStmt.Close()
	setStmt.ExecContext(ctx, &id, x.Salutation, x.Subject)

	upsStmt := tx.Prepare(`
		INSERT INTO hello (hello_id, salutation, subject)
		VALUES (@id, @sal, @sub)
		ON DUPLICATE KEY UPDATE
		hello_id = @id, salutation = @sal, subject = @sub
	`)
	defer upsStmt.Close()
	upsStmt.ExecContext(ctx)

	tx.Commit()
	if err := tx.Err(); err != nil {
		return err
	}

	y.Id = id.String()
	y.Salutation = x.Salutation
	y.Subject = x.Subject

	return nil
}

func (s *HeloDB) deleteHello(ctx cx, sid string) error {
	db := s.db.Scope()

	delStmt := db.Prepare(`
		UPDATE hello
		SET deleted_at = NOW()
		WHERE hello_id = ?
	`)
	defer delStmt.Close()
	delStmt.ExecContext(ctx, sid)

	return db.Err()
}

func (s *HeloDB) findHellos(ctx cx, x *pb.FndHellosReq, y *pb.HellosResp) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	setStmt := tx.Prepare(`
		SET @ids = ?, @sals = ?
	`)
	defer setStmt.Close()
	setStmt.ExecContext(ctx, csvStr(x.Ids), csvStr(x.Salutations))

	selStmt := tx.Prepare(`
		SELECT
			SQL_CALC_FOUND_ROWS hello_id, salutation, subject,
			created_at, updated_at, deleted_at
		FROM hello
		WHERE IF(@ids != "", FIND_IN_SET(hello_id, @ids), 0=0)
		AND IF(@sals != "", FIND_IN_SET(salutation, @sals), 0=0)
		LIMIT ?
		OFFSET ?
	`)
	defer selStmt.Close()
	rows := selStmt.QueryContext(ctx, altr.LimitUint32(x.Limit), x.Lapse)
	defer rows.Close()

	for rows.Next() {
		r := pb.HelloResp{}

		var tc, tu, td mysql.NullTime
		rows.Scan(
			&r.Id, &r.Salutation, &r.Subject, &tc, &tu, &td,
		)

		r.Created = asTS(tc.Time, tc.Valid)
		r.Updated = asTS(tu.Time, tu.Valid)
		r.Deleted = asTS(td.Time, td.Valid)

		y.Items = append(y.Items, &r)
	}

	fndStmt := tx.Prepare(`
		SELECT FOUND_ROWS()
	`)
	defer fndStmt.Close()
	row := fndStmt.QueryRowContext(ctx)
	row.Scan(&y.Total)

	tx.Commit()
	return tx.Err()
}
