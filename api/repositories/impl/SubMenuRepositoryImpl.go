package impl

import (
	"errors"
	"fmt"
	"github.com/gosimple/slug"
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"gorm.io/gorm"
	"math"
	"strings"
	"time"
)

type SubMenuRepositoryImpl struct {
	db *gorm.DB
}

func NewSubMenuRepositoryImpl(db *gorm.DB) *SubMenuRepositoryImpl {
	return &SubMenuRepositoryImpl{db:db}
}

func (r *SubMenuRepositoryImpl) GetAll(pagination *request.Pagination) (interface{}, error, int) {
	var subMenus []entities.SubMenu

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

	find = find.Find(&subMenus)

	if find.Error != nil {
		return nil, find.Error, totalPages
	}

	pagination.Rows = subMenus

	counting := int64(totalRows)

	// count all data
	err := r.db.Model(&entities.SubMenu{}).Count(&counting).Error

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

func (r *SubMenuRepositoryImpl) FindBySubMenu(namaMenu string, namaSubMenu string) (entities.SubMenu, error) {
	var menu entities.Menu
	var subMenu entities.SubMenu

	r.db.Where("nama_menu=?", namaMenu).Take(&menu)

	if menu.NamaMenu == "" {
		return entities.SubMenu{}, errors.New("Nama menu tidak ditemukan")
	}

	r.db.Where("id_menu=? AND nama_sub_menu=?", menu.ID, namaSubMenu).Take(&subMenu)

	if subMenu.NamaSubMenu == "" {
		return entities.SubMenu{}, errors.New("Nama sub menu tidak ditemukan")
	}

	return subMenu, nil
}

func (r *SubMenuRepositoryImpl) CreateSubMenu(idMenu int64, subMenus []request.SubMenuRequest) (bool, error) {
	var menu entities.Menu

	r.db.Where("id_menu=?", idMenu).Take(&menu)

	if menu.NamaMenu == "" {
		return false, errors.New("Id menu tidak ditemukan")
	}

	for _, val := range subMenus {
		var subMenu entities.SubMenu

		r.db.Where("nama_sub_menu=?", val.NamaSubMenu).Take(&subMenu)

		if subMenu.NamaSubMenu != "" {
			return false, errors.New("Nama sub menu sudah ada")
		}

		val.IdMenu = menu.ID
		val.Slug = slug.Make(val.NamaSubMenu)
		val.CreateAt = time.Now().Format("2006-01-02 15:04:05")
		val.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

		row := r.db.Omit("update_date").Create(&val)

		if row.Error != nil {
			return false, row.Error
		}
	}

	return true, nil
}

func (r *SubMenuRepositoryImpl) UpdateSubMenu(idMenu int64, idSubMenu int64, subMenus request.SubMenuRequest) (bool, error) {
	var menu entities.Menu

	r.db.Where("id_menu=?", idMenu).Take(&menu)

	if menu.ID == 0 {
		return false, errors.New("Id Menu tidak ditemukan")
	}

	var subMenu entities.SubMenu

	r.db.Where("id_sub_menu=?", idSubMenu).Take(&subMenu)

	if subMenu.NamaSubMenu == "" {
		return false, errors.New("Id Sub Menu tidak ditemukan")
	}

	row := r.db.Omit("create_date").Model(&subMenus).Where("id_sub_menu=?", subMenu.ID).Updates(&subMenus)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *SubMenuRepositoryImpl) DeleteSubMenu(idMenu int64, idSubMenu int64) (bool, error) {
	var menu entities.Menu

	r.db.Where("id_menu=?", idMenu).Take(&menu)

	if menu.ID == 0 {
		return false, errors.New("Id Menu tidak ditemukan")
	}

	var subMenu entities.SubMenu

	r.db.Where("id_sub_menu=?", idSubMenu).Take(&subMenu)

	if subMenu.NamaSubMenu == "" {
		return false, errors.New("Id Sub Menu tidak ditemukan")
	}

	row := r.db.Model(&entities.SubMenu{}).Where("id_sub_menu=?", idSubMenu).Delete(subMenu)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *SubMenuRepositoryImpl) FindById(idMenu int64, idSubMenu int64) (entities.SubMenu, error) {
	var menu entities.Menu
	var subMenu entities.SubMenu

	r.db.Where("id_menu=?", idMenu).Take(&menu)

	if menu.ID == 0 {
		return entities.SubMenu{}, errors.New("Id menu tidak ditemukan")
	}

	r.db.Where("id_menu=? AND id_sub_menu=?", menu.ID, idSubMenu).Take(&subMenu)

	if subMenu.ID == 0 {
		return entities.SubMenu{}, errors.New("Id sub menu tidak ditemukan")
	}

	return subMenu, nil
}

func (r *SubMenuRepositoryImpl) Enable(submenuId int64) (bool, error) {
	var subMenu entities.SubMenu

	r.db.Where("id_sub_menu=?", submenuId).Take(&subMenu)

	if subMenu.ID == 0 {
		return false, errors.New("Id sub menu tidak ditemukan")
	}

	row := r.db.Model(&entities.SubMenu{}).Select("status").Where("id_sub_menu=?", subMenu.ID).Update("status", true)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *SubMenuRepositoryImpl) Disable(submenuId int64) (bool, error) {
	var subMenu entities.SubMenu

	r.db.Where("id_sub_menu=?", submenuId).Take(&subMenu)

	if subMenu.ID == 0 {
		return false, errors.New("Id sub menu tidak ditemukan")
	}

	row := r.db.Model(&entities.SubMenu{}).Select("status").Where("id_sub_menu=?", subMenu.ID).Update("status", false)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}