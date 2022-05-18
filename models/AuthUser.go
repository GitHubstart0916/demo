package models

import "database/sql"

type AuthUser struct {
	Id       int            `db:"id"`
	UserName string         `db:"userName"`
	Password string         `db:"password"`
	Email    sql.NullString `db:"email"`
}
