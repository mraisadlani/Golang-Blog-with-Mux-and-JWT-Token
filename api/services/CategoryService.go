package services

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/repositories/impl"
	"go-blog-jwt-token/api/utils"
	"gorm.io/gorm"
	"net/http"
)

func GetAllCategory(r *http.Request, db *gorm.DB, pagination *request.Pagination) (interface{}, error) {
	repo := impl.NewCategoryRepositoryImpl(db)

	get, err, totalPages := repo.GetAll(pagination)

	if err != nil {
		return nil, err
	}

	err = utils.SetupPagination(r, get, pagination, totalPages)

	if err != nil {
		return nil, err
	}

	return get, nil
}

func FindCategory(db *gorm.DB, IdCategory int64) (entities.Category, error) {
	repo := impl.NewCategoryRepositoryImpl(db)

	get, err := repo.FindCategory(IdCategory)

	if err != nil {
		return entities.Category{}, err
	}

	return get, nil
}

func CreateCategory(db *gorm.DB, category request.CategoryRequest) (bool, error) {
	repo := impl.NewCategoryRepositoryImpl(db)

	get, err := repo.CreateCategory(category)

	if err != nil {
		return false, err
	}

	return get, nil
}

func UpdateCategory(db *gorm.DB, category request.CategoryRequest, IdCategory int64) (bool, error) {
	repo := impl.NewCategoryRepositoryImpl(db)

	get, err := repo.UpdateCategory(category, IdCategory)

	if err != nil {
		return false, err
	}

	return get, nil
}

func DeleteCategory(db *gorm.DB, IdCategory int64) (bool, error) {
	repo := impl.NewCategoryRepositoryImpl(db)

	get, err := repo.DeleteCategory(IdCategory)

	if err != nil {
		return false, err
	}

	return get, nil
}