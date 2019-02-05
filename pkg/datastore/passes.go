package datastore

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jacobkaufmann/hallpass"
	"github.com/jmoiron/modl"
)

type passesStore struct{ *Datastore }

func (s *passesStore) Get(id int64) (*hallpass.Pass, error) {
	var passes []*hallpass.Pass
	if err := s.dbh.Select(&passes, `SELECT * FROM pass WHERE pass_id=$1;`, id); err != nil {
		return nil, err
	}
	if len(passes) == 0 {
		return nil, hallpass.ErrPassNotFound
	}
	return passes[0], nil
}

func (s *passesStore) List(opt *hallpass.PassListOptions) ([]*hallpass.Pass, error) {
	if opt == nil {
		opt = &hallpass.PassListOptions{}
	}
	var sql = `SELECT * FROM pass`

	var conds []string
	if opt.UserID >= 0 {
		conds = append(conds, "user_id="+strconv.FormatInt(opt.UserID, 10)+"")
	}
	if opt.Student != "" {
		conds = append(conds, "student='"+opt.Student+"'")
	}

	if len(conds) > 0 {
		sql += " WHERE (" + strings.Join(conds, ") AND (") + ")"
	}

	sql += " ORDER BY created_date DESC;"

	var passes []*hallpass.Pass
	err := s.dbh.Select(&passes, sql)
	if err != nil {
		return nil, err
	}
	return passes, nil
}

func (s *passesStore) Create(p *hallpass.Pass) (bool, error) {
	var created bool

	err := transact(s.dbh, func(tx modl.SqlExecutor) error {
		if err := tx.Insert(p); err != nil {
			return err
		}

		created = true
		return nil
	})

	return created, err
}

func (s *passesStore) Update(p *hallpass.Pass) (bool, error) {
	var updated bool

	err := transact(s.dbh, func(tx modl.SqlExecutor) error {
		var count int64
		var err error
		if count, err = tx.Update(p); err != nil {
			return err
		}
		if count != 1 {
			return fmt.Errorf("failed to update pass with id: %d", p.ID)
		}

		updated = true
		return nil
	})

	return updated, err
}
