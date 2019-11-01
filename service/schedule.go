package service

import (
	"fmt"
	`github.com/pashukhin/pik-arenda-test-task/entity`
	`github.com/jmoiron/sqlx`
	"strings"
	"time"
)

type Schedule struct {
	db *sqlx.DB
	fs *FreeSchedule
}

func NewSchedule(db *sqlx.DB, fs *FreeSchedule) *Schedule {
	return &Schedule{db, fs}
}

func (s *Schedule) List(from, to *time.Time) (list []*entity.Schedule, err error) {
	query := make([]string, 0)
	params := make([]interface{}, 0)
	if from != nil {
		params = append(params, *from)
		query = append(query, fmt.Sprintf(`"end" > $%d`, len(params)))
	}
	if to != nil {
		params = append(params, *to)
		query = append(query, fmt.Sprintf(`"start" < $%d`, len(params)))
	}
	sql := `select * from "schedule"`
	if len(query) > 0 {
		sql += ` where ` + strings.Join(query, ` and `)
	}
	err = s.db.Select(&list, sql, params...)
	return
}

func (s *Schedule) Create(o *entity.Schedule) (err error) {
	tx := s.db.MustBegin()
	defer func(){
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err = tx.Get(o, `INSERT INTO "schedule" ("worker_id", "start", "end") VALUES ($1, $2, $3) returning *`, o.WorkerID, o.Start, o.End); err != nil {
		return
	}
	return s.fs.add(tx, o.Start, o.End, 1)
}

func (s *Schedule) Get(id int) (o *entity.Schedule, err error) {
	o = &entity.Schedule{}
	err = s.db.Get(o, `select * from "schedule" where "id" = $1`, id)
	return
}

func (s *Schedule) Update(o *entity.Schedule) (err error) {
	tx := s.db.MustBegin()
	defer func(){
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	oldObject := &entity.Schedule{}
	if err = s.db.Get(oldObject, `select * from "schedule" where "id" = $1`, o.ID); err != nil {
		return
	}
	// old := &interval{oldObject.Start, oldObject.End}
	//intersection, left, right := intersectionAndParts(old, &interval {o.Start, o.End})
	//if intersection == nil { // completely moved
	//	if err = s.fs.check(o.Start, o.End); err != nil { // just check new position
	//		return err
	//	}
	//} else { // moved with intersection
	//	if left != nil && left.start.Before(old.start) { // start moved left => left is new part
	//		if err = s.fs.check(left.start, left.end); err != nil { // just check new position
	//			return err
	//		}
	//	}
	//	if right != nil && right.end.After(old.end) { // end moved right => right is new part
	//		if err = s.fs.check(right.start, right.end); err != nil { // just check new position
	//			return err
	//		}
	//	}
	//}
	if err = s.fs.add(tx, oldObject.Start, oldObject.End, -1); err != nil {
		return err
	}
	if err = tx.Get(o, `update "schedule" set "worker_id" = $1, "start" = $2, "end" = $3 where "id" = $4 returning *`, o.WorkerID, o.Start, o.End, o.ID); err != nil {
		return err
	}
	if err = s.fs.add(tx, o.Start, o.End, 1); err != nil {
		return err
	}
	return
}

func (s *Schedule) Delete(id int) (o *entity.Schedule, err error) {
	tx := s.db.MustBegin()
	defer func(){
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err = tx.Get(o, `delete from "schedule" where "id" = $1 returning *`, id); err != nil {
		return
	}
	err = s.fs.add(tx, o.Start, o.End, -1)
	return
}
