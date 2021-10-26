package models

import (
	"backend/db"
	"backend/entity"
)

func PostGetAll() (interface{}, error) {
	var (
		mulObj []entity.Post
	)

	db := db.GetDB()

	result := db.Find(&mulObj).Limit(30)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return "data kosong", nil
	}

	return mulObj, nil
}

func PostSearchID(id string) (interface{}, error) {
	var (
		obj entity.Post
	)

	db := db.GetDB()

	result := db.Where("post_id = ?", id).First(&obj)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return "data kosong", nil
	}

	return obj, nil
}
