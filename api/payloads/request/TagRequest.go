package request

import (
	"errors"
	"github.com/gosimple/slug"
	"time"
)

type TagRequest struct {
	ID int64 `json:"id_tag" gorm:"column:id_tag"`
	NamaTag string `json:"nama_tag"`
	Slug string `json:"slug"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (e *TagRequest) TableName() string {
	return "tb_tags"
}

func (e *TagRequest) Prepare() {
	e.Slug = slug.Make(e.NamaTag)
	e.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
}

func (e *TagRequest) Validate() error {
	if e.NamaTag == "" {
		return errors.New("Nama tag tidak boleh kosong")
	}

	return nil
}