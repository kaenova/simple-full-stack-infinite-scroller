package models

import (
	"backend/db"
	"backend/entity"
)

func PostGetPage(page int, pageSize int) (interface{}, error) {
	var (
		mulObj []entity.Post
	)

	db := db.GetDB()
	offset := (page - 1) * pageSize
	result := db.Offset(offset).Limit(pageSize).Find(&mulObj)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return mulObj, nil
}
