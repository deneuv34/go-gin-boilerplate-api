package repository

import "fdnBaseAPI/commons/config"

type BaseRepository struct {
	Model interface{}
	TableName string
}

func (r BaseRepository) Init(model interface{}) {
	r.Model = model
}
func (r BaseRepository) Create(data interface{}) (interface{}, error) {
	db := config.GetDB()
	err := db.Model(&r.Model).Create(data).Error

	return data, err
}

func (r BaseRepository) Delete(data interface{}) error {
	db := config.GetDB()
	err := db.Model(&r.Model).Delete(data).Error
	return err
}

func (r BaseRepository) FindOne(data interface{}) (interface{}, error) {
	db := config.GetDB()
	err := db.Model(&r.Model).First(data).Error
	return &r.Model, err
}
