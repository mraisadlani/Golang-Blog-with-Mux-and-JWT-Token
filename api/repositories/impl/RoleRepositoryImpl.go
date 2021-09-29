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

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepositoryImpl(db *gorm.DB) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{db}
}

func (r *RoleRepositoryImpl) GetAll(pagination *request.Pagination) (interface{}, error, int) {
	var roles []entities.Role

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

	find = find.Find(&roles)

	if find.Error != nil {
		return nil, find.Error, totalPages
	}

	pagination.Rows = roles

	counting := int64(totalRows)

	// count all data
	err := r.db.Model(&entities.Role{}).Count(&counting).Error

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

func (r *RoleRepositoryImpl) FindByNamaRole(rolename string) (entities.Role, error) {
	var role entities.Role

	r.db.Where("nama_role=?", rolename).Take(&role)

	if role.NamaRole == "" {
		return entities.Role{}, errors.New("Nama role tidak ditemukan")
	}

	return role, nil
}

func (r *RoleRepositoryImpl) FindById(idRole int64) (entities.Role, error) {
	var role entities.Role

	r.db.Where("id_role=?", idRole).Take(&role)

	if role.ID <= 0 {
		return entities.Role{}, errors.New("Id Role tidak ditemukan")
	}

	return role, nil
}

func (r *RoleRepositoryImpl) Insert(roleReq request.RoleRequest) (bool, error) {
	var role entities.Role

	r.db.Where("nama_role=?", roleReq.NamaRole).Take(&role)

	if role.ID > 0 {
		return false, errors.New("Nama role sudah ada")
	}
	row := r.db.Omit("update_at").Create(&roleReq)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *RoleRepositoryImpl) Update(roleReq request.RoleRequest, roleId int64) (bool, error) {
	var role entities.Role

	r.db.Where("id_role=?", roleId).Take(&role)

	if role.ID == 0 {
		return false, errors.New("Id Role tidak ditemukan")
	}

	row := r.db.Omit("create_at").Model(&roleReq).Where("id_role=?", role.ID).Updates(&roleReq)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *RoleRepositoryImpl) Delete(roleId int64) (bool, error) {
	var role entities.Role

	r.db.Where("id_role=?", roleId).Take(&role)

	if role.ID == 0 {
		return false, errors.New("Id role tidak ditemukan")
	}

	row := r.db.Model(&entities.Role{}).Where("id_role=?", roleId).Delete(role)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *RoleRepositoryImpl) Enable(roleId int64) (bool, error) {
	var role entities.Role

	r.db.Where("id_role=?", roleId).Take(&role)

	if role.ID == 0 {
		return false, errors.New("Id role tidak ditemukan")
	}

	row := r.db.Model(&entities.Role{}).Select("status").Where("id_role=?", role.ID).Update("status", true)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *RoleRepositoryImpl) Disable(roleId int64) (bool, error) {
	var role entities.Role

	r.db.Where("id_role=?", roleId).Take(&role)

	if role.ID == 0 {
		return false, errors.New("Id role tidak ditemukan")
	}

	row := r.db.Model(&entities.Role{}).Select("status").Where("id_role=?", role.ID).Update("status", false)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}