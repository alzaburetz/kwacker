package storage_models

import "time"

type User struct {
	tableName       struct{}  `pg:users`
	Id              int64     `pg:"id,pk"`
	Name            string    `pg:"name,type:varchar(255)"`
	Surname         string    `pg:"surname,type:varchar(255)"`
	Created         time.Time `pg:"created_time"`
	Edited          time.Time `pg:"edited_time"`
	ProfilePicture  string    `pg:"profile_picture"`
	Email           string    `pg:"email,type:varchar(100)"`
	NickNameHandler string    `pg:"nickname,type:varchar(100)"`
	Posts           int       `pg:"posts"`
	Status          string    `pg:"status"`
	Bio             string    `pg:"bio"`
	Likes           int64     `pg:"likes"`
}
