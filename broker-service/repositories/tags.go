package repositories

import (
	"broker/data"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TagRepository struct {
	db *gorm.DB
}

func CreateTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) Insert(tag *data.Tag) error {
	return r.db.Create(tag).Error
}

func (r *TagRepository) UpdateTag(tag *data.Tag) error {
	return r.db.Save(tag).Error
}

func (r *TagRepository) Get(id string) (*data.Tag, error) {
	tag := new(data.Tag)
	err := r.db.Where("id = ?", id).Find(tag).Error
	return tag, err
}
func (r *TagRepository) GetOrCreate(title string) (*data.Tag, error) {
	tag := new(data.Tag)
	err := r.db.Where("title = ?", title).Find(tag).Error

	if err != nil || tag.ID == 0 {
		tag.Title = title
		r.db.Create(tag)
	}
	return tag, err
}

func (r *TagRepository) Delete(id string) error {
	result := r.db.Select(clause.Associations).Unscoped().Delete(&data.Tag{}, id)
	return result.Error
}

func (r *TagRepository) GetAll() ([]data.Tag, error) {
	var tags []data.Tag
	err := r.db.Order("title").Find(&tags).Error
	//err := r.db.Order("created_at desc").Find(&tags).Error
	return tags, err
}
