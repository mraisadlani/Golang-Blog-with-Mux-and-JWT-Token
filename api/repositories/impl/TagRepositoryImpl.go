package impl

import (
	"errors"
	"fmt"
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"gorm.io/gorm"
	"math"
	"strings"
)

type TagRepositoryImpl struct {
	db *gorm.DB
}

func NewTagRepositoryImpl(db *gorm.DB) *TagRepositoryImpl {
	return &TagRepositoryImpl{db}
}

func (r *TagRepositoryImpl) GetAll(pagination *request.Pagination) (interface{}, error, int) {
	var tags []entities.Tag

	totalRows := 0
	totalPages := 0
	fromRow := 0
	toRow := 0

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset & order
	find := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	// generate where query
	searchs := pagination.Searchs

	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s=?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break
			}
		}
	}

	find = find.Find(&tags)

	if find.Error != nil {
		return nil, find.Error, totalPages
	}

	pagination.Rows = tags

	counting := int64(totalRows)

	// count all data
	err := r.db.Model(&entities.Tag{}).Count(&counting).Error

	if err != nil {
		return nil, err, totalPages
	}

	totalRows = int(counting)

	pagination.TotalRows = totalRows

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page * pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > totalRows {
		toRow = totalRows
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return pagination, nil, totalPages
}

func (r *TagRepositoryImpl) FindTag(idTag int64) (entities.Tag, error) {
	var tag entities.Tag

	r.db.Where("id_tag=?", idTag).Take(&tag)

	if tag.ID == 0 {
		return entities.Tag{}, errors.New("Id tag tidak ditemukan")
	}

	return tag, nil
}

func (r *TagRepositoryImpl) CreateTag(tagRequest request.TagRequest) (bool, error) {
	var tag entities.Tag

	r.db.Where("nama_tag=?", tagRequest.NamaTag).Take(&tag)

	if tag.ID > 0 {
		return false, errors.New("Nama tag sudah ada")
	}

	row := r.db.Omit("update_at").Create(&tagRequest)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *TagRepositoryImpl) UpdateTag(tagRequest request.TagRequest, idTag int64) (bool, error) {
	var tag entities.Tag

	r.db.Where("id_tag=?", idTag).Take(&tag)

	if tag.ID == 0 {
		return false, errors.New("Id tag tidak ditemukan")
	}

	row := r.db.Omit("create_at").Model(&tagRequest).Where("id_tag=?", tag.ID).Updates(&tagRequest)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *TagRepositoryImpl) DeleteTag(idTag int64) (bool, error) {
	var tag entities.Tag

	r.db.Where("id_tag=?", idTag).Take(&tag)

	if tag.ID == 0 {
		return false, errors.New("Id tag tidak ditemukan")
	}

	row := r.db.Model(&entities.Tag{}).Where("id_tag=?", idTag).Delete(tag)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}