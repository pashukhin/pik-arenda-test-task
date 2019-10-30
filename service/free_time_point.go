package service

import (
	"errors"
	"fmt"
	`github.com/pashukhin/pik-arenda-test-task/entity`
	`github.com/jmoiron/sqlx`
	"strings"
	"time"
)

type FreeTimePoint struct {
	db *sqlx.DB
}

func NewFreeTimePoint(db *sqlx.DB) *FreeTimePoint {
	return &FreeTimePoint{db}
}

func (s *FreeTimePoint) List(from, to *time.Time) (list []*entity.FreeTimePoint, err error) {
	query := make([]string, 0)
	params := make([]interface{}, 0)
	if from != nil {
		params = append(params, *from)
		query = append(query, fmt.Sprintf(`"point" >= $%d`, len(params)))
	}
	if to != nil {
		params = append(params, *to)
		query = append(query, fmt.Sprintf(`"point" <= $%d`, len(params)))
	}
	sql := `select * from "free_time_point"`
	if len(query) > 0 {
		sql += ` where ` + strings.Join(query, ` and `)
	}
	err = s.db.Select(&list, sql, params...)
	return
}

func (s *FreeTimePoint) Create(o *entity.FreeTimePoint) (err error) {
	err = errors.New("method not allowed")
	return
}

func (s *FreeTimePoint) Get(id int) (o *entity.FreeTimePoint, err error) {
	o = &entity.FreeTimePoint{}
	err = s.db.Get(o, `select * from "free_time_point" where "id" = $1`, id)
	return
}

func (s *FreeTimePoint) Update(o *entity.FreeTimePoint) (err error) {
	err = errors.New("method not allowed")
	return
}

func (s *FreeTimePoint) Delete(id int) (o *entity.FreeTimePoint, err error) {
	err = errors.New("method not allowed")
	return
}
