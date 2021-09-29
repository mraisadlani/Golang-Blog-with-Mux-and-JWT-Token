package services

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/repositories/impl"
	"go-blog-jwt-token/api/utils"
	"gorm.io/gorm"
	"net/http"
)

func GetAllPost(r *http.Request, db *gorm.DB, pagination *request.Pagination) (interface{}, error) {
	repo := impl.NewPostRepositoryImpl(db)

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

func CreatePost(db *gorm.DB, postRequest request.PostRequest, urlImage string, categories []string, tags []string) (bool, error) {
	repo := impl.NewPostRepositoryImpl(db)

	get, err := repo.CreatePost(postRequest, urlImage, categories, tags)

	if err != nil {
		return false, err
	}

	return get, nil
}

func FindPost(db *gorm.DB, IdArticle int64, IdPost int64) (entities.Post, error) {
	repo := impl.NewPostRepositoryImpl(db)

	get, err := repo.FindPost(IdArticle, IdPost)

	if err != nil {
		return entities.Post{}, err
	}

	return get, nil
}

func UpdatePost(db *gorm.DB, postRequest request.PostRequest, urlImage string, categories []string, tags []string) (bool, error) {
	repo := impl.NewPostRepositoryImpl(db)

	get, err := repo.UpdatePost(postRequest, categories, urlImage, tags)

	if err != nil {
		return false, err
	}

	return get, nil
}

func DeletePost(db *gorm.DB, IdArticle int64, IdPost int64) (bool, error) {
	repo := impl.NewPostRepositoryImpl(db)

	get, err := repo.DeletePost(IdArticle, IdPost)

	if err != nil {
		return false, err
	}

	return get, nil
}

func PublishPost(db *gorm.DB, IdArticle int64, IdPost int64) (bool, error) {
	repo := impl.NewPostRepositoryImpl(db)

	get, err := repo.PublishPost(IdArticle, IdPost)

	if err != nil {
		return false, err
	}

	return get, nil
}

func CancelPost(db *gorm.DB, IdArticle int64, IdPost int64) (bool, error) {
	repo := impl.NewPostRepositoryImpl(db)

	get, err := repo.CancelPost(IdArticle, IdPost)

	if err != nil {
		return false, err
	}

	return get, nil
}