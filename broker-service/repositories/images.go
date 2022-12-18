package repositories

import (
	"broker/data"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ImageRepository struct {
	db *gorm.DB
}

func CreateImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{db: db}
}

func (r *ImageRepository) Insert(img *data.ImagesDescription) error {
	return r.db.Create(img).Error
}

func (r *ImageRepository) Get(id string) (*data.ImagesDescription, error) {
	image := new(data.ImagesDescription)
	err := r.db.Where("id = ?", id).Find(image).Error
	return image, err
}
func (r *ImageRepository) GetByFilename(filename string) (*data.ImagesDescription, error) {
	image := new(data.ImagesDescription)
	err := r.db.Where("file_name = ?", filename).Preload("Tags").Find(image).Error
	return image, err
}

func (r *ImageRepository) Delete(id string) error {
	result := r.db.Select(clause.Associations).Unscoped().Delete(&data.ImagesDescription{}, id)
	return result.Error
}

func (r *ImageRepository) DeleteByFileName(filename string) error {
	argFilename := "%" + filename + "%"
	result := r.db.Where("file_name LIKE ?", argFilename).Delete(&data.ImagesDescription{})
	return result.Error
}

func (r *ImageRepository) GetAll() ([]data.ImagesDescription, error) {
	var imgs []data.ImagesDescription
	err := r.db.Order("created_at desc").Model(&data.ImagesDescription{}).Preload("Tags").Find(&imgs).Error
	return imgs, err
}

func (r *ImageRepository) GetAllW() ([]data.ImagesDescription, error) {
	var imgs []data.ImagesDescription
	err := r.db.Find(&imgs).Error
	return imgs, err
}

func (r *ImageRepository) UpdateTags(img_id string, tags []*data.Tag) error {
	image, err := r.Get(img_id)
	if err != nil {
		return err
	} else {
		r.db.Model(&image).Association("Tags").Replace(tags)
	}
	return nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
