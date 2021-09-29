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

// @Summary Get All Sub Menu
// @Description REST API Sub Menu
// @Accept  json
// @Produce  json
// @Tags Sub Menu Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /SubMenuAPI/ [get]
func GetAllSubMenu(w http.ResponseWriter, r *http.Request) {
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

	get, err := services.GetAllSubMenu(r, db, pagination)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
}

// @Summary Find by Sub Menu
// @Description REST API Sub Menu
// @Accept  json
// @Produce  json
// @Tags Sub Menu Controller
// @Param nama_menu path string true "Nama Menu"
// @Param nama_sub_menu path string true "Nama Sub Menu"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /SubMenuAPI/{nama_menu}/GetSubMenu/{nama_sub_menu} [get]
func GetSubMenu(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)

	get, err := services.FindSubMenu(db, vars["nama_menu"], vars["nama_sub_menu"])

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
}

// @Summary Create Sub Menu
// @Description REST API Sub Menu
// @Accept  json
// @Produce  json
// @Tags Sub Menu Controller
// @Param id_menu path string true "Id Menu"
// @Param requestBody body []request.SubMenuRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /SubMenuAPI/{id_menu}/CreateSubMenu [post]
func CreateSubMenu(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	idMenu, err := strconv.ParseInt(vars["id_menu"], 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	var subMenus []request.SubMenuRequest
	err = json.NewDecoder(r.Body).Decode(&subMenus)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	for _, val := range subMenus {
		err = val.Validate()

		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err)
			return
		}
	}

	get, err := services.CreateSubMenu(db, idMenu, subMenus)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)
}

// @Summary Update Sub Menu
// @Description REST API Sub Menu
// @Accept  json
// @Produce  json
// @Tags Sub Menu Controller
// @Param id_menu path string true "Id Menu"
// @Param id_sub_menu path string true "Id Sub Menu"
// @Param requestBody body request.SubMenuRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /SubMenuAPI/{id_menu}/UpdateSubMenu/{id_sub_menu} [post]
func UpdateSubMenu(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	idMenu, err := strconv.ParseInt(vars["id_menu"], 10, 64)
	idSubMenu, err := strconv.ParseInt(vars["id_sub_menu"], 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	var subMenus request.SubMenuRequest
	err = json.NewDecoder(r.Body).Decode(&subMenus)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	subMenus.Prepare()
	err = subMenus.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	subMenus.IdMenu = idMenu

	get, err := services.UpdateSubMenu(db, idMenu, idSubMenu, subMenus)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)
}

// @Summary Delete Sub Menu
// @Description REST API Sub Menu
// @Accept  json
// @Produce  json
// @Tags Sub Menu Controller
// @Param id_menu path string true "Id Menu"
// @Param id_sub_menu path string true "Id Sub Menu"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /SubMenuAPI/{id_menu}/DeleteSubMenu/{id_sub_menu} [post]
func DeleteSubMenu(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	idMenu, err := strconv.ParseInt(vars["id_menu"], 10, 64)
	idSubMenu, err := strconv.ParseInt(vars["id_sub_menu"], 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	get, err := services.DeleteSubMenu(db, idMenu, idSubMenu)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil menghapus data", get, http.StatusOK)
}

// @Summary Status Sub Menu
// @Description REST API Sub Menu
// @Accept  json
// @Produce  json
// @Tags Sub Menu Controller
// @Param id_menu path string true "Id Menu"
// @Param id_sub_menu path string true "Id Sub Menu"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /SubMenuAPI/{id_menu}/StatusSubMenu/{id_sub_menu} [post]
func StatusSubMenu(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	IdMenu, _ := strconv.ParseInt(vars["id_menu"], 10, 64)
	IdSubMenu, _ := strconv.ParseInt(vars["id_sub_menu"], 10, 64)

	get, err := services.FindByIdSubMenu(db, IdMenu, IdSubMenu)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	set, err := services.StatusSubMenu(db, get)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah status", set, http.StatusOK)
}