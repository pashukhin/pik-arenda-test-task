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
	old := &interval{o.Start, o.End}
	if err = tx.Get(o, `update "schedule" set "worker_id" = $1, "start" = $2, "end" = $3 where "id" = $4 returning *`, o.WorkerID, o.Start, o.End, o.ID); err != nil {
		return err
	}
	_, left, right := intersectionAndParts(old, &interval {o.Start, o.End})
	if left != nil {
		if left.start == old.start { // moved right
			if err = s.fs.add(tx, left.start, left.end, -1); err != nil {
				return err
			}
		} else { // moved left
			if err = s.fs.add(tx, left.start, left.end, 1); err != nil {
				return err
			}
		}
	}
	if right != nil {
		if right.end == old.end { // moved left
			if err = s.fs.add(tx, right.start, right.end, -1); err != nil {
				return err
			}
		} else { // moved right
			if err = s.fs.add(tx, right.start, right.end, 1); err != nil {
				return err
			}
		}
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
