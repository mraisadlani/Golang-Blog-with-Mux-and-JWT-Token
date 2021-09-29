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

// @Summary Get All Tag
// @Description REST API Tag
// @Accept  json
// @Produce  json
// @Tags Tag Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /TagAPI/ [get]
func GetAllTag(w http.ResponseWriter, r *http.Request) {
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

	get, err := services.GetAllTag(r, db, pagination)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
}

// @Summary Find by Tag
// @Description REST API Tag
// @Accept  json
// @Produce  json
// @Tags Tag Controller
// @Param id_tag path string true "Id Tag"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /TagAPI/{id_tag}/GetTag [get]
func GetTag(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	IdTag, err := strconv.ParseInt(vars["id_tag"], 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	get, err := services.FindTag(db, IdTag)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
}

// @Summary Create Tag
// @Description REST API Tag
// @Accept  json
// @Produce  json
// @Tags Tag Controller
// @Param requestBody body request.TagRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /TagAPI/CreateTag [post]
func CreateTag(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	var tags request.TagRequest
	err = json.NewDecoder(r.Body).Decode(&tags)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	tags.Prepare()
	err = tags.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := services.CreateTag(db, tags)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)
}

// @Summary Update Tag
// @Description REST API Tag
// @Accept  json
// @Produce  json
// @Tags Tag Controller
// @Param id_tag path string true "Id Tag"
// @Param requestBody body request.TagRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /TagAPI/{id_tag}/UpdateTag [post]
func UpdateTag(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	IdTag, err := strconv.ParseInt(vars["id_tag"], 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	var tag request.TagRequest
	err = json.NewDecoder(r.Body).Decode(&tag)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	tag.Prepare()
	err = tag.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := services.UpdateTag(db, tag, IdTag)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)
}

// @Summary Delete Tag
// @Description REST API Tag
// @Accept  json
// @Produce  json
// @Tags Tag Controller
// @Param id_tag path string true "Id Tag"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /TagAPI/{id_tag}/DeleteTag [post]
func DeleteTag(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	IdTag, err := strconv.ParseInt(vars["id_tag"], 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	get, err := services.DeleteTag(db, IdTag)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil menghapus data", get, http.StatusOK)
}