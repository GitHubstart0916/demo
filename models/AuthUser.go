package models

type AuthUser struct {
	Id       int    `db:"id"`
	UserName string `db:"userName"`
	Password string `db:"password"`
}
