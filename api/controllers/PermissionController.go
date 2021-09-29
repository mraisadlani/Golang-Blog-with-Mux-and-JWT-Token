package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-blog-jwt-token/api/configs"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/payloads/response"
	"go-blog-jwt-token/api/services"
	"net/http"
	"strconv"
)

// @Summary Find Permission
// @Description REST API Permission
// @Accept  json
// @Produce  json
// @Tags Permission Controller
// @Param id_user path string true "Id User"
// @Param id_menu path string true "Id Menu"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /PermissionAPI/{id_user}/GetPermission/{id_menu} [get]
func GetPermission(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)

	IdUser, err := strconv.ParseInt(vars["id_user"], 10, 64)
	IdMenu, err := strconv.ParseInt(vars["id_menu"], 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	get, err := services.FindPermission(db, IdUser, IdMenu)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
}

// @Summary Create Permission
// @Description REST API Permission
// @Accept  json
// @Produce  json
// @Tags Permission Controller
// @Param requestBody body []request.PermissionRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /PermissionAPI/CreatePermission [post]
func CreatePermission(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	var permission []request.PermissionRequest
	err = json.NewDecoder(r.Body).Decode(&permission)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	get, err := services.CreatePermission(db, permission)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)
}

// @Summary Update Permission
// @Description REST API Permission
// @Accept  json
// @Produce  json
// @Tags Permission Controller
// @Param requestBody body []request.PermissionRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /PermissionAPI/UpdatePermission [post]
func UpdatePermission(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	var permission []request.PermissionRequest
	err = json.NewDecoder(r.Body).Decode(&permission)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	get, err := services.UpdatePermission(db, permission)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)
}

// @Summary Delete Permission
// @Description REST API Permission
// @Accept  json
// @Produce  json
// @Tags Permission Controller
// @Param id_permission path string true "Id Permission"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /PermissionAPI/{id_permission}/DeletePermission [post]
func DeletePermission(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)

	IdPermission, err := strconv.ParseInt(vars["id_permission"], 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	get, err := services.DeletePermission(db, IdPermission)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil menghapus data", get, http.StatusOK)
}