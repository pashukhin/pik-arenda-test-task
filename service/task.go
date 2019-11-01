package service

import (
	"fmt"
	`github.com/pashukhin/pik-arenda-test-task/entity`
	`github.com/jmoiron/sqlx`
	"strings"
	"time"
)

type Task struct {
	db *sqlx.DB
	fs *FreeSchedule
}

func NewTask(db *sqlx.DB, fs *FreeSchedule) *Task {
	return &Task{db, fs}
}

func (s *Task) List(from, to *time.Time) (list []*entity.Task, err error) {
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
	sql := `select * from "task"`
	if len(query) > 0 {
		sql += ` where ` + strings.Join(query, ` and `)
	}
	err = s.db.Select(&list, sql, params...)
	return
}

func (s *Task) Create(o *entity.Task) (err error) {
	if err = s.fs.check(o.Start, o.End); err != nil {
		return err
	}
	tx := s.db.MustBegin()
	defer func(){
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err = tx.Get(o, `INSERT INTO "task" ("start", "end") VALUES ($1, $2) returning *`, o.Start, o.End); err != nil {
		tx.Rollback()
		return
	}
	return s.fs.add(tx, o.Start, o.End, -1)
}

func (s *Task) Get(id int) (o *entity.Task, err error) {
	o = &entity.Task{}
	err = s.db.Get(o, `select * from "task" where "id" = $1`, id)
	return
}

func (s *Task) Update(o *entity.Task) (err error) {
	tx := s.db.MustBegin()
	defer func(){
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	oldObject := &entity.Task{}
	if err = s.db.Get(oldObject, `select * from "task" where "id" = $1`, o.ID); err != nil {
		return
	}
	// check for free slots
	old := &interval{oldObject.Start, oldObject.End}
	intersection, left, right := intersectionAndParts(old, &interval {o.Start, o.End})
	if intersection == nil { // completely moved
		if err = s.fs.check(o.Start, o.End); err != nil { // just check new position
			return err
		}
	} else { // moved with intersection
		if left != nil && left.start.Before(old.start) { // start moved left => left is new part
			if err = s.fs.check(left.start, left.end); err != nil { // just check new position
				return err
			}
		}
		if right != nil && right.end.After(old.end) { // end moved right => right is new part
			if err = s.fs.check(right.start, right.end); err != nil { // just check new position
				return err
			}
		}
	}
	if err = s.fs.add(tx, oldObject.Start, oldObject.End, 1); err != nil {
		return err
	}
	if err = tx.Get(o, `update "task" set "worker_id" = $1, "start" = $2, "end" = $3, "cancelled" = $4 where "id" = $5 returning *`, o.WorkerID, o.Start, o.End, o.Cancelled, o.ID); err != nil {
		return err
	}
	if err = s.fs.add(tx, o.Start, o.End, -1); err != nil {
		return err
	}
	return
}

func (s *Task) Delete(id int) (o *entity.Task, err error) {
	err = s.db.Get(o, `delete from "task" where "id" = $1 returning *`, id)
	return
}
