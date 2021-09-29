package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-blog-jwt-token/api/configs"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/payloads/response"
	"go-blog-jwt-token/api/services"
	"go-blog-jwt-token/api/utils"
	"net/http"
	"strconv"
)

// @Summary Get All Role
// @Description REST API Role
// @Accept  json
// @Produce  json
// @Tags Role Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /RoleAPI/GetAllRoles [get]
func GetAllRole(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	pagination, err := utils.SortPagination(r)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	get, err := services.GetAllRole(r, pagination, db)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
	}
}

// @Summary Create Role
// @Description REST API Role
// @Accept  json
// @Produce  json
// @Tags Role Controller
// @Param reqBody body request.RoleRequest true "Form Request"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /RoleAPI/CreateRole [post]
func CreateRole(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	var roleReq request.RoleRequest
	err = json.NewDecoder(r.Body).Decode(&roleReq)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	roleReq.Prepare()
	err = roleReq.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := services.InsertRole(db, roleReq)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)
	}
}

// @Summary Find Role
// @Description REST API Role
// @Accept  json
// @Produce  json
// @Tags Role Controller
// @Param role_name path string true "Role name"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /RoleAPI/{role_name}/FindByRoleName [post]
func FindByRoleName(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)

	get, err := services.FindByRoleName(db, vars["role_name"])

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
	}
}

// @Summary Update Role
// @Description REST API Role
// @Accept  json
// @Produce  json
// @Tags Role Controller
// @Param id_role path string true "Id Role"
// @Param reqBody body request.RoleRequest true "Form Request"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /RoleAPI/{id_role}/UpdateRole [post]
func UpdateRole(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	roleId, _ := strconv.ParseInt(vars["id_role"], 10, 64)

	var roleReq request.RoleRequest
	err = json.NewDecoder(r.Body).Decode(&roleReq)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	roleReq.Prepare()
	err = roleReq.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := services.EditRole(db, roleId, roleReq)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)
	}
}

// @Summary Delete Role
// @Description REST API Role
// @Accept  json
// @Produce  json
// @Tags Role Controller
// @Param id_role path string true "Id Role"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /RoleAPI/{id_role}/DeleteRole [post]
func DeleteRole(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	roleId, _ := strconv.ParseInt(vars["id_role"], 10, 64)

	get, err := services.DeleteRole(db, roleId)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil menghapus data", get, http.StatusOK)
	}
}

// @Summary Status Role
// @Description REST API Role
// @Accept  json
// @Produce  json
// @Tags Role Controller
// @Param id_role path string true "Id Role"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /RoleAPI/{id_role}/StatusRole [post]
func StatusRole(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	roleId, _ := strconv.ParseInt(vars["id_role"], 10, 64)

	get, err := services.FindByIdRole(db, roleId)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	set, err := services.StatusRole(db, get)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mengubah status", set, http.StatusOK)
	}
}