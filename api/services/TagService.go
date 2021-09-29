package services

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/repositories/impl"
	"go-blog-jwt-token/api/utils"
	"gorm.io/gorm"
	"net/http"
)

func GetAllTag(r *http.Request, db *gorm.DB, pagination *request.Pagination) (interface{}, error) {
	repo := impl.NewTagRepositoryImpl(db)

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

func FindTag(db *gorm.DB, IdTag int64) (entities.Tag, error) {
	repo := impl.NewTagRepositoryImpl(db)

	get, err := repo.FindTag(IdTag)

	if err != nil {
		return entities.Tag{}, err
	}

	return get, nil
}

func CreateTag(db *gorm.DB, tags request.TagRequest) (bool, error) {
	repo := impl.NewTagRepositoryImpl(db)

	get, err := repo.CreateTag(tags)

	if err != nil {
		return false, err
	}

	return get, nil
}

func UpdateTag(db *gorm.DB, tag request.TagRequest, IdTag int64) (bool, error) {
	repo := impl.NewTagRepositoryImpl(db)

	get, err := repo.UpdateTag(tag, IdTag)

	if err != nil {
		return false, err
	}

	return get, nil
}

func DeleteTag(db *gorm.DB, IdTag int64) (bool, error) {
	repo := impl.NewTagRepositoryImpl(db)

	get, err := repo.DeleteTag(IdTag)

	if err != nil {
		return false, err
	}

	return get, nil
}