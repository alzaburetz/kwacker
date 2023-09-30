package storage_models

type BlockType uint8

const (
	None BlockType = iota
	Text
	Image
)

type PostBlock struct {
	tableName     struct{} `pg:"post_blocks"`
	TrackableID   int
	TrackableType string
	Id            int64     `pg:"id,pk"`
	Type          BlockType `pg:"type"`
	Content       string    `pg:"content"`
	Style         string    `pg:"style"`
	PostId        int64     `pg:"post_id"`
}
