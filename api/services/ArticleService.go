package services

import (
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/repositories/impl"
	"go-blog-jwt-token/api/utils"
	"gorm.io/gorm"
	"net/http"
)

func GetAllArticle(r *http.Request, db *gorm.DB, pagination *request.Pagination) (interface{}, error) {
	repo := impl.NewArticleRepositoryImpl(db)

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

func FindArticle(db *gorm.DB, IdArticle int64) (entities.Article, error) {
	repo := impl.NewArticleRepositoryImpl(db)

	get, err := repo.FindArticle(IdArticle)

	if err != nil {
		return entities.Article{}, err
	}

	return get, nil
}

func CreateArticle(db *gorm.DB, articles request.ArticleRequest) (bool, error) {
	repo := impl.NewArticleRepositoryImpl(db)

	get, err := repo.CreateArticle(articles)

	if err != nil {
		return false, err
	}

	return get, nil
}

func UpdateArticle(db *gorm.DB, article request.ArticleRequest, IdArticle int64) (bool, error) {
	repo := impl.NewArticleRepositoryImpl(db)

	get, err := repo.UpdateArticle(article, IdArticle)

	if err != nil {
		return false, err
	}

	return get, nil
}

func DeleteArticle(db *gorm.DB, IdArticle int64) (bool, error) {
	repo := impl.NewArticleRepositoryImpl(db)

	get, err := repo.DeleteArticle(IdArticle)

	if err != nil {
		return false, err
	}

	return get, nil
}