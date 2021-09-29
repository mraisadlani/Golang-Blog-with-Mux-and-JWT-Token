package request

import (
	"errors"
	"github.com/gosimple/slug"
	"time"
)

type PostRequest struct {
	ID int64 `json:"id_post" gorm:"column:id_post"`
	NamaPost string `json:"nama_post"`
	Slug string `json:"slug"`
	Description string `json:"description"`
	Published bool `json:"published"`
	IdArticle int64 `json:"id_article" gorm:"column:id_article"`
	CreateBy string `json:"create_by"`
	CreateAt string `json:"create_at"`
	UpdateBy string `json:"update_by"`
	UpdateAt string `json:"update_at"`
}

func (e *PostRequest) TableName() string {
	return "tb_posts"
}

func (e *PostRequest) Prepare() {
	e.Slug = slug.Make(e.NamaPost)
	e.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
}

func (e *PostRequest) Validate() error {
	if e.NamaPost == "" {
		return errors.New("Nama Post tidak boleh kosong")
	}

	if e.Description == "" {
		return errors.New("Description tidak boleh kosong")
	}

	return nil
}