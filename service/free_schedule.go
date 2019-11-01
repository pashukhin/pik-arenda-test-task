package service

import (
	"database/sql"
	"errors"
	"fmt"
	`github.com/pashukhin/pik-arenda-test-task/entity`
	`github.com/jmoiron/sqlx`
	"strings"
	"time"
)

type FreeSchedule struct {
	db *sqlx.DB
}

func NewFreeSchedule(db *sqlx.DB) *FreeSchedule {
	return &FreeSchedule{db}
}

var ErrNoFreeSlot = errors.New("no free slot")

func (s *FreeSchedule) add(tx *sqlx.Tx, from, to time.Time, value int) (err error) {
	// find interval containing from
	cFrom := &entity.FreeSchedule{}
	if e := tx.Get(cFrom, `select * from "free_schedule" where "start" < $1 and "end" > $1`, from); e != nil {
		if e != sql.ErrNoRows {
			err = e
			return
		}
		cFrom = nil
	}
	// find interval containing to
	cTo := &entity.FreeSchedule{}
	if e := tx.Get(cTo, `select * from "free_schedule" where "start" < $1 and "end" > $1`, to); e != nil {
		if e != sql.ErrNoRows {
			err = e
			return
		}
		cTo = nil
	}
	// update all intervals intersects with (from, to)
	sqlUpdate := `update "free_schedule" set "start" = greatest("start", $1), "end" = least("end", $2), "value" = "value" + $3 where "start" < $2 and "end" > $1`
	res, err := tx.Exec(sqlUpdate, from, to, value)
	if err != nil {
		return
	}
	var ra int64
	if ra, err = res.RowsAffected(); err != nil {
		return
	}
	sqlInsert := `INSERT INTO "free_schedule" ("start", "end", "value") VALUES ($1, $2, $3)`
	// nothing updated
	if ra == 0 {
		// insert new interval
		if _, err = tx.Query(sqlInsert, from, to, value); err != nil {
			return
		}

	}
	if cFrom != nil {
		// insert new interval before intersected part
		if _, err = tx.Query(sqlInsert, cFrom.Start, from, cFrom.Value); err != nil {
			return
		}
	}
	if cTo != nil {
		// insert new interval before intersected part
		if _, err = tx.Query(sqlInsert, to, cTo.End, cTo.Value); err != nil {
			return
		}
	}
	return
}

func (s *FreeSchedule) getFreeSlots(from, to *time.Time) (result []*entity.FreeSlot, err error) {
	var intersects []*entity.FreeSchedule
	intersects, err = s.List(from, to)
	if err != nil {
		return
	}
	var f, t *time.Time
	for _, fs := range intersects {
		if fs.Value > 0 {
			if f == nil {
				f = &fs.Start
			}
			t = &fs.End
		} else {
			if f != nil && t != nil {
				result = append(result, &entity.FreeSlot{*f, *t})
			}
			f, t = nil, nil
		}
	}
	if f != nil && t != nil {
		result = append(result, &entity.FreeSlot{*f, *t})
	}
	return
}

func (s * FreeSchedule) check(from, to time.Time) error {
	if slots, err := s.getFreeSlots(&from, &to); err != nil {
		return err
	} else {
		for _, fs := range slots {
			if (fs.Start.Before(from) || fs.Start.Equal(from)) && (to.Before(fs.End) || to.Equal(fs.End)) {
				return nil
			}
		}
		return ErrNoFreeSlot
	}
}

func (s *FreeSchedule) List(from, to *time.Time) (list []*entity.FreeSchedule, err error) {
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
	sql := `select * from "free_schedule"`
	if len(query) > 0 {
		sql += ` where ` + strings.Join(query, ` and `)
	}
	sql += ` order by "start"`
	err = s.db.Select(&list, sql, params...)
	return
}

func (s *FreeSchedule) Create(o *entity.FreeSchedule) (err error) {
	err = errors.New("method not allowed")
	return
}

func (s *FreeSchedule) Get(id int) (o *entity.FreeSchedule, err error) {
	o = &entity.FreeSchedule{}
	err = s.db.Get(o, `select * from "free_schedule" where "id" = $1`, id)
	return
}

func (s *FreeSchedule) Update(o *entity.FreeSchedule) (err error) {
	err = errors.New("method not allowed")
	return
}

func (s *FreeSchedule) Delete(id int) (o *entity.FreeSchedule, err error) {
	err = errors.New("method not allowed")
	return
}
