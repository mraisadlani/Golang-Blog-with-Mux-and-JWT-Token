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

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{db}
}

func (r *CategoryRepositoryImpl) GetAll(pagination *request.Pagination) (interface{}, error, int) {
	var category []entities.Category

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

	find = find.Find(&category)

	if find.Error != nil {
		return nil, find.Error, totalPages
	}

	pagination.Rows = category

	counting := int64(totalRows)

	// count all data
	err := r.db.Model(&entities.Category{}).Count(&counting).Error

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

func (r *CategoryRepositoryImpl) FindCategory(idCategory int64) (entities.Category, error) {
	var category entities.Category

	r.db.Where("id_category=?", idCategory).Take(&category)

	if category.ID == 0 {
		return entities.Category{}, errors.New("Id category tidak ditemukan")
	}

	return category, nil
}

func (r *CategoryRepositoryImpl) CreateCategory(categoryRequest request.CategoryRequest) (bool, error) {
	var category entities.Category

	r.db.Where("nama_category=?", categoryRequest.NamaCategory).Take(&category)

	if category.ID > 0 {
		return false, errors.New("Nama category sudah ada")
	}

	row := r.db.Omit("update_at").Create(&categoryRequest)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *CategoryRepositoryImpl) UpdateCategory(categoryRequest request.CategoryRequest, idCategory int64) (bool, error) {
	var category entities.Category

	r.db.Where("id_category=?", idCategory).Take(&category)

	if category.ID == 0 {
		return false, errors.New("Id category tidak ditemukan")
	}

	row := r.db.Omit("create_at").Model(&categoryRequest).Where("id_category=?", category.ID).Updates(&categoryRequest)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *CategoryRepositoryImpl) DeleteCategory(idCategory int64) (bool, error) {
	var category entities.Category

	r.db.Where("id_category=?", idCategory).Take(&category)

	if category.ID == 0 {
		return false, errors.New("Id category tidak ditemukan")
	}

	row := r.db.Model(&entities.Category{}).Where("id_category=?", idCategory).Delete(category)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}