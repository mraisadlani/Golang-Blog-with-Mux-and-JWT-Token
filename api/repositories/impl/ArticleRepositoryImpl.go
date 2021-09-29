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

type ArticleRepositoryImpl struct {
	db *gorm.DB
}

func NewArticleRepositoryImpl(db *gorm.DB) *ArticleRepositoryImpl {
	return &ArticleRepositoryImpl{db:db}
}

func (r *ArticleRepositoryImpl) GetAll(pagination *request.Pagination) (interface{}, error, int) {
	var articles []entities.Article

	totalRows := 0
	totalPages := 0
	fromRow := 0
	toRow := 0

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset & order
	find := r.db.Preload("Post").Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

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

	find = find.Find(&articles)

	if find.Error != nil {
		return nil, find.Error, totalPages
	}

	pagination.Rows = articles

	counting := int64(totalRows)

	// count all data
	err := r.db.Model(&entities.Article{}).Count(&counting).Error

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

func (r *ArticleRepositoryImpl) FindArticle(articleId int64) (entities.Article, error) {
	var article entities.Article

	r.db.Preload("Post").Where("id_article=?", articleId).Take(&article)

	if article.NamaArticle == "" {
		return entities.Article{}, errors.New("Nama article tidak ditemukan")
	}

	return article, nil
}

func (r *ArticleRepositoryImpl) CreateArticle(articleRequest request.ArticleRequest) (bool, error) {
	var article entities.Article

	r.db.Where("nama_article=?", articleRequest.NamaArticle).Take(&article)

	if article.ID > 0 {
		return false, errors.New("Nama article sudah ada")
	}

	row := r.db.Omit("update_at").Create(&articleRequest)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *ArticleRepositoryImpl) UpdateArticle(articleRequest request.ArticleRequest, articleId int64) (bool, error) {
	var article entities.Article

	r.db.Where("id_article=?", articleId).Take(&article)

	if article.ID == 0 {
		return false, errors.New("Id Article tidak ditemukan")
	}

	row := r.db.Omit("create_at").Model(&articleRequest).Where("id_article=?", article.ID).Updates(&articleRequest)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *ArticleRepositoryImpl) DeleteArticle(articleId int64) (bool, error) {
	var article entities.Article

	r.db.Where("id_article=?", articleId).Take(&article)

	if article.ID == 0 {
		return false, errors.New("Id article tidak ditemukan")
	}

	row := r.db.Model(&entities.Article{}).Where("id_article=?", articleId).Delete(article)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}