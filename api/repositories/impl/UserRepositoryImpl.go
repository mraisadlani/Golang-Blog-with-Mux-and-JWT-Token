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

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) GetAll(pagination *request.Pagination) (interface{}, error, int) {
	var user []entities.User

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
		for _, val := range searchs {
			column := val.Column
			action := val.Action
			query := val.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
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

	find = find.Find(&user)

	if find.Error != nil {
		return nil, find.Error, totalPages
	}

	pagination.Rows = user

	counting := int64(totalRows)

	// count all
	err := r.db.Model(&entities.User{}).Count(&counting).Error

	if err != nil {
		return nil, find.Error, totalPages
	}

	totalRows = int(counting)
	pagination.TotalRows = totalRows

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > totalRows {
		// set to row with total rows
		toRow = totalRows
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return pagination, nil, totalPages
}

func (r *UserRepositoryImpl) FindById(idUser int64) (entities.User, error) {
	var user entities.User

	r.db.Where("id_user=?", idUser).Preload("Role").Take(&user)

	if user.ID <= 0 {
		return entities.User{}, errors.New("Id User tidak ditemukan")
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (entities.User, error) {
	var user entities.User

	r.db.Where("username=?", username).Take(&user)

	if user.ID <= 0 {
		return entities.User{}, errors.New("Username tidak ditemukan")
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (entities.User, error) {
	var user entities.User

	r.db.Where("email=?", email).Take(&user)

	if user.ID <= 0 {
		return entities.User{}, errors.New("Email tidak ditemukan")
	}

	return user, nil
}

func (r *UserRepositoryImpl) Insert(reqUser request.UserRequest) (bool, error) {
	var user entities.User

	r.db.Where("email=? OR username=?", reqUser.Email, reqUser.Username).Take(&user)

	if user.Username != "" {
		return false, errors.New("Username sudah terdaftar")
	} else if user.Email != "" {
		return false, errors.New("Email sudah terdaftar")
	} else if user.Username != "" && user.Email != "" {
		return false, errors.New("Username dan Email sudah terdaftar")
	}

	err := r.db.Exec("INSERT INTO tb_users (first_name, last_name, username, password, email, no_telp, photo, id_role, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		reqUser.FirstName,
		reqUser.LastName,
		reqUser.Username,
		reqUser.Password,
		reqUser.Email,
		reqUser.NoTelp,
		reqUser.Photo,
		reqUser.IdRole,
		reqUser.Status,
	).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepositoryImpl) Update(reqUser request.UserRequest, idUser int64) (bool, error) {
	get, err := r.FindById(idUser)

	if err != nil {
		return false, err
	}

	row := r.db.Omit("create_at, photo, username").Model(&reqUser).Where("id_user=?", get.ID).Updates(&reqUser)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *UserRepositoryImpl) Delete(idUser int64) (bool, error) {
	get, err := r.FindById(idUser)

	if err != nil {
		return false, err
	}

	row := r.db.Model(&entities.User{}).Where("id_user=?", get.ID).Delete(get)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *UserRepositoryImpl) Enable(idUser int64) (bool, error) {
	get, err := r.FindById(idUser)

	if err != nil {
		return false, err
	}

	row := r.db.Model(&entities.User{}).Select("status").Where("id_user=?", get.ID).Update("status", true)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *UserRepositoryImpl) Disable(idUser int64) (bool, error) {
	get, err := r.FindById(idUser)

	if err != nil {
		return false, err
	}

	row := r.db.Model(&entities.User{}).Select("status").Where("id_user=?", get.ID).Update("status", false)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *UserRepositoryImpl) UploadImage(filename string, userId int64) (bool, error) {
	get, err := r.FindById(userId)

	if err != nil {
		return false, err
	}

	get.Photo = filename

	row := r.db.Model(&get).Select("photo").Where("id_user=?", get.ID).Updates(&get)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}