package domain

import "github.com/jmoiron/sqlx"

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) {

}
