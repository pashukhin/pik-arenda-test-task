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
}

func NewTask(db *sqlx.DB) *Task {
	return &Task{db}
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
	tx := s.db.MustBegin()
	if err = tx.Get(o, `INSERT INTO "task" ("start", "end") VALUES ($1, $2) returning *`, o.Start, o.End); err != nil {
		tx.Rollback()
		return
	}
	// select all free_schedule included in task
	// update it to value++
	// select one free_schedule A including task.start
	// insert new free_schedule with start = task.start, end = A.end, value = A.value+1
	// update A to end = task.start
	if _, err = tx.NamedExec(`insert into "free_time_point" ("task_id", "point", "value") values (:id, :start, -1), (:id, :end, 1)`, o); err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (s *Task) Get(id int) (o *entity.Task, err error) {
	o = &entity.Task{}
	err = s.db.Get(o, `select * from "task" where "id" = $1`, id)
	return
}

func (s *Task) Update(o *entity.Task) (err error) {
	tx := s.db.MustBegin()
	if err = tx.Get(o, `update "task" set "worker_id" = $1, "start" = $2, "end" = $3, "cancelled" = $4 where "id" = $5 returning *`, o.WorkerID, o.Start, o.End, o.Cancelled, o.ID); err != nil {
		return err
	}
	if _, err = tx.NamedExec(`delete from "free_time_point" where "task_id" = :id`, o); err != nil {
		tx.Rollback()
		return err
	}
	if !o.Cancelled {
		if _, err = tx.NamedExec(`insert into "free_time_point" ("task_id", "point", "value") values (:id, :start, -1), (:id, :end, 1)`, o); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return
}

func (s *Task) Delete(id int) (o *entity.Task, err error) {
	err = s.db.Get(o, `delete from "task" where "id" = $1 returning *`, id)
	return
}
