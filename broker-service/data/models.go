package data

import "time"

type User struct {
	ID             string `gorm:"column:id;primary_key" json:"id"`
	Username       string `gorm:"column:name" json:"username"`
	HashedPassword []byte `gorm:"column:hashed_password" json:"-"`
}

type ImagesDescription struct {
	ID       uint   `gorm:"primaryKey;"`
	FileName string `gorm:"column:file_name;unique"`
	Tags     []*Tag `gorm:"many2many:image_tags; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TimedEntity
}

type Tag struct {
	ID    uint   `gorm:"primaryKey;"`
	Title string `gorm:"unique"`
	TimedEntity
}

type TimedEntity struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tabler interface {
	TableName() string
}

func (ImagesDescription) TableName() string {
	return "images"

}
func (Tag) TableName() string {
	return "tags"
}
