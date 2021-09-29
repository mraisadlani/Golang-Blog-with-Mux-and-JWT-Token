package impl

import (
	"errors"
	"go-blog-jwt-token/api/entities"
	"go-blog-jwt-token/api/payloads/request"
	"gorm.io/gorm"
	"time"
)

type PermissionRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissionRepositoryImpl(db *gorm.DB) *PermissionRepositoryImpl {
	return &PermissionRepositoryImpl{db:db}
}

func (r *PermissionRepositoryImpl) FindPermission(IdUser int64, IdMenu int64) (entities.Permission, error) {
	var user entities.User

	r.db.Where("id_user=?", IdUser).Take(&user)

	if user.Username == "" {
		return entities.Permission{}, errors.New("Id User tidak ditemukan")
	}

	var menu entities.Menu

	r.db.Where("id_menu=?", IdMenu).Take(&menu)

	if menu.NamaMenu == "" {
		return entities.Permission{}, errors.New("Id Menu tidak ditemukan")
	}

	var permission entities.Permission

	r.db.Preload("Menu").Preload("User").Where("id_user=? AND id_menu=?", user.ID, menu.ID).Take(&permission)

	if permission.ID == 0 || permission.ID < 1 {
		return entities.Permission{}, errors.New("Data tidak ditemukan")
	}

	return permission, nil
}

func (r *PermissionRepositoryImpl) CreatePermission(permission []request.PermissionRequest) (bool, error) {
	for _, val := range permission {
		var user entities.User

		r.db.Where("id_user=?", val.IdUser).Take(&user)

		if user.ID == 0 || user.ID < 1 {
			return false, errors.New("Id User tidak ditemukan")
		}

		var menu entities.Menu

		r.db.Where("id_menu=?", val.IdMenu).Take(&menu)

		if menu.ID == 0 {
			return false, errors.New("Id Menu tidak ditemukan")
		}

		var permis entities.Permission

		r.db.Where("id_user=? AND id_menu=?", user.ID, menu.ID).Take(&permis)

		if permis.ID != 0 {
			return false, errors.New("data sudah terdaftar")
		}

		val.CreateAt = time.Now().Format("2006-01-02 15:04:05")
		row := r.db.Create(&val)

		if row.Error != nil {
			return false, row.Error
		}
	}

	return true, nil
}

func (r *PermissionRepositoryImpl) UpdatePermission(permission []request.PermissionRequest) (bool, error) {
	for _, val := range permission {
		var permis entities.Permission

		r.db.Where("id_permission=?", val.ID).Take(&permis)

		if permis.ID == 0 {
			return false, errors.New("Id permission tidak ditemukan")
		}

		var user entities.User

		r.db.Where("id_user=?", val.IdUser).Take(&user)

		if user.ID == 0 || user.ID < 1 {
			return false, errors.New("Id User tidak ditemukan")
		}

		var menu entities.Menu

		r.db.Where("id_menu=?", val.IdMenu).Take(&menu)

		if menu.ID == 0 {
			return false, errors.New("Id Menu tidak ditemukan")
		}

		row := r.db.Exec("UPDATE tb_has_permission SET f_create=?, f_read=?, f_update=?, f_delete=?, f_publish=? WHERE id_permission=? AND id_menu=? AND id_user=?", val.FCreate, val.FRead, val.FUpdate, val.FDelete, val.FPublish, val.ID, val.IdMenu, val.IdUser)

		if row.Error != nil {
			return false, row.Error
		}
	}

	return true, nil
}

func (r *PermissionRepositoryImpl) DeletePermission(IdPermission int64) (bool, error) {
	var permission entities.Permission

	r.db.Where("id_permission=?", IdPermission).Take(&permission)

	if permission.ID == 0 {
		return false, errors.New("Id Permission tidak ditemukan")
	}

	row := r.db.Model(&entities.Permission{}).Where("id_permission=?", IdPermission).Delete(permission)

	if row.Error != nil {
		return false, row.Error
	}
	return true, nil
}