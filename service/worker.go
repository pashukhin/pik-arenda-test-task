package service

import (
	`github.com/jmoiron/sqlx`
	`github.com/pashukhin/pik-arenda-test-task/entity`
)

type Worker struct {
	db *sqlx.DB
}

func NewWorker(db *sqlx.DB) *Worker {
	return &Worker{db}
}

func (s *Worker) List() (list []*entity.Worker, err error) {
	err = s.db.Select(&list, `select * from "worker"`)
	return
}

func (s *Worker) Create(o *entity.Worker) (err error) {
	return s.db.Get(o, `insert into "worker" ("name") values ($1) returning *`, o.Name)
}

func (s *Worker) Get(id int) (o *entity.Worker, err error) {
	o = &entity.Worker{}
	err = s.db.Get(o, `select * from "worker" where "id" = $1`, id)
	return
}

func (s *Worker) Update(o *entity.Worker) (err error) {
	return s.db.Get(o, `update "worker" set "name" = $1 where "id" = $2 returning *`, o.Name, o.ID)
}

func (s *Worker) Delete(id int) (o *entity.Worker, err error) {
	err = s.db.Get(o, `delete from "worker" where "id" = $1 returning *`, id)
	return
}
