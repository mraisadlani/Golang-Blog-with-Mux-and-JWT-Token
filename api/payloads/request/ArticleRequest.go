package request

import (
	"errors"
	"github.com/gosimple/slug"
	"time"
)

type ArticleRequest struct {
	ID int64 `json:"id_article" gorm:"column:id_article"`
	NamaArticle string `json:"nama_article"`
	Slug string `json:"slug"`
	Icon string `json:"icon"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (e *ArticleRequest) TableName() string {
	return "tb_articles"
}

func (e *ArticleRequest) Prepare() {
	e.Slug = slug.Make(e.NamaArticle)
	e.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
}

func (e *ArticleRequest) Validate() error {
	if e.NamaArticle == "" {
		return errors.New("Nama Article tidak boleh kosong")
	}

	if e.Icon == "" {
		return errors.New("Icon tidak boleh kosong")
	}

	return nil
}