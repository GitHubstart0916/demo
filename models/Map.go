package models

type Map struct {
	Id          int    `db:"id"`
	Path        string `db:"path"`
	Create_time string `db:"create_time"`
	User_id     int    `db:"user_id"`
	Update_time string `db:"update_time"`
	Count       string `db:"count"`
}
