package models

type EmailToken struct {
	Token string `db:"token"`
}
