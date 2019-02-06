package datastore

import (
	"fmt"
	"strings"

	"github.com/jmoiron/modl"

	"github.com/jacobkaufmann/hallpass"
)

type usersStore struct{ *Datastore }

func (s *usersStore) Get(id int64) (*hallpass.User, error) {
	var users []*hallpass.User
	if err := s.dbh.Select(&users, `SELECT * FROM "user" WHERE user_id=$1;`, id); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, hallpass.ErrUserNotFound
	}
	return users[0], nil
}

func (s *usersStore) List(opt *hallpass.UserListOptions) ([]*hallpass.User, error) {
	if opt == nil {
		opt = &hallpass.UserListOptions{}
	}
	var sql = `SELECT * FROM "user"`

	var conds []string
	if opt.Email != "" {
		conds = append(conds, "email LIKE '%"+opt.Email+"%'")
	}
	if opt.Org != "" {
		conds = append(conds, "org='"+opt.Org+"'")
	}

	if len(conds) > 0 {
		sql += " WHERE (" + strings.Join(conds, ") AND (") + ")"
	}

	sql += " ORDER BY email DESC;"

	var users []*hallpass.User
	err := s.dbh.Select(&users, sql)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *usersStore) Create(u *hallpass.User) (bool, error) {
	var created bool

	err := transact(s.dbh, func(tx modl.SqlExecutor) error {
		var existing []*hallpass.User
		if err := tx.Select(&existing, `SELECT * FROM "user" WHERE email=$1 LIMIT 1;`, u.Email); err != nil {
			return err
		}
		if len(existing) > 0 {
			*u = *existing[0]
			return nil
		}

		if err := tx.Insert(u); err != nil {
			return err
		}

		created = true
		return nil
	})

	return created, err
}

func (s *usersStore) Delete(id int64) (bool, error) {
	var deleted bool

	u, err := s.Get(id)
	if err != nil {
		return deleted, err
	}

	err = transact(s.dbh, func(tx modl.SqlExecutor) error {
		var count int64
		if count, err = tx.Delete(u); err != nil {
			return err
		}
		if count != 1 {
			return fmt.Errorf("failed to delete user with id: %d", id)
		}

		return nil
	})

	return deleted, nil
}

func (s *usersStore) Update(u *hallpass.User) (bool, error) {
	var updated bool

	err := transact(s.dbh, func(tx modl.SqlExecutor) error {
		var count int64
		var err error
		if count, err = tx.Update(u); err != nil {
			return err
		}
		if count != 1 {
			return fmt.Errorf("failed to update user with id: %d", u.ID)
		}

		return nil
	})

	return updated, err
}
