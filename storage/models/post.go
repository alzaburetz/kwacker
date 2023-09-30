package storage_models

type Post struct {
	tableName struct{}    `pg:"posts"`
	Id        int64       `pg:"id,pk"`
	HeadLine  string      `pg:"headline"`
	Likes     int         `pg:"likes"`
	Blocks    []PostBlock `pg:"rel:has-many,join_fk:trackable_,polymorphic"`
	Author    *User       `pg:"rel:has-one"`
	AuthorId  int         `pg:"author_id"`
}
