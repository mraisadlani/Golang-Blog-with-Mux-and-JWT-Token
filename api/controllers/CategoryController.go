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

// @Summary Get All Category
// @Description REST API Category
// @Accept  json
// @Produce  json
// @Tags Category Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /CategoryAPI/ [get]
func GetAllCategory(w http.ResponseWriter, r *http.Request) {
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

	get, err := services.GetAllCategory(r, db, pagination)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
}

// @Summary Find by Category
// @Description REST API Category
// @Accept  json
// @Produce  json
// @Tags Category Controller
// @Param id_category path string true "Id Category"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /CategoryAPI/{id_category}/GetCategory [get]
func GetCategory(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	IdCategory, err := strconv.ParseInt(vars["id_category"], 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	get, err := services.FindCategory(db, IdCategory)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
}

// @Summary Create Category
// @Description REST API Category
// @Accept  json
// @Produce  json
// @Tags Category Controller
// @Param requestBody body request.CategoryRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /CategoryAPI/CreateCategory [post]
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	var category request.CategoryRequest
	err = json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	category.Prepare()
	err = category.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := services.CreateCategory(db, category)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)
}

// @Summary Update Category
// @Description REST API Category
// @Accept  json
// @Produce  json
// @Tags Category Controller
// @Param id_category path string true "Id Category"
// @Param requestBody body request.CategoryRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /CategoryAPI/{id_category}/UpdateCategory [post]
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	IdCategory, err := strconv.ParseInt(vars["id_category"], 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	var category request.CategoryRequest
	err = json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	category.Prepare()
	err = category.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := services.UpdateCategory(db, category, IdCategory)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)
}

// @Summary Delete Category
// @Description REST API Category
// @Accept  json
// @Produce  json
// @Tags Category Controller
// @Param id_category path string true "Id Category"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /CategoryAPI/{id_category}/DeleteCategory [post]
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	IdCategory, err := strconv.ParseInt(vars["id_category"], 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	get, err := services.DeleteCategory(db, IdCategory)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil menghapus data", get, http.StatusOK)
}