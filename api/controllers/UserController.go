package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
	"go-blog-jwt-token/api/configs"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/payloads/response"
	"go-blog-jwt-token/api/securities"
	"go-blog-jwt-token/api/services"
	"go-blog-jwt-token/api/utils"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// @Summary Get All User
// @Description REST API User
// @Accept  json
// @Produce  json
// @Tags User Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /UserAPI/GetAllUsers [get]
func GetAllUser(w http.ResponseWriter, r *http.Request) {
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

	get, err := services.GetAllUser(r, pagination, db)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
	}
}

// @Summary Create User
// @Description REST API User
// @Accept  json
// @Produce  json
// @Tags User Controller
// @Param reqBody body request.UserRequest true "Form Request"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /UserAPI/CreateUser [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	var userReq request.UserRequest
	err = json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	userReq.Prepare()
	err = userReq.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	hash, err := securities.HashPassword(userReq.Password)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	userReq.Password = hash

	get, err := services.InsertUser(db, userReq)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)
	}
}

// @Summary Find User
// @Description REST API User
// @Accept  json
// @Produce  json
// @Tags User Controller
// @Param id_user path string true "Id User"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /UserAPI/{id_user}/FindById [post]
func FindById(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["id_user"], 10, 64)

	get, err := services.FindById(db, userId)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
	}
}

// @Summary Update User
// @Description REST API User
// @Accept  json
// @Produce  json
// @Tags User Controller
// @Param id_user path string true "Id User"
// @Param reqBody body request.UserRequest true "Form Request"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /UserAPI/{id_user}/UpdateUser [post]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["id_user"], 10, 64)

	var userReq request.UserRequest
	err = json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	userReq.Prepare()
	err = userReq.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	hash, err := securities.HashPassword(userReq.Password)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	userReq.Password = hash

	get, err := services.EditUser(db, userId, userReq)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)
	}
}

// @Summary Delete User
// @Description REST API User
// @Accept  json
// @Produce  json
// @Tags User Controller
// @Param id_user path string true "Id User"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /UserAPI/{id_user}/DeleteUser [post]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["id_user"], 10, 64)

	get, err := services.DeleteUser(db, userId)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil menghapus data", get, http.StatusOK)
	}
}

// @Summary Status User
// @Description REST API User
// @Accept  json
// @Produce  json
// @Tags User Controller
// @Param id_user path string true "Id User"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /UserAPI/{id_user}/StatusUser [post]
func StatusUser(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["id_user"], 10, 64)

	get, err := services.FindById(db, userId)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	set, err := services.StatusUser(db, get)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mengubah status", set, http.StatusOK)
	}
}

// @Summary Upload Image User
// @Description REST API User
// @Accept  json
// @Produce  json
// @Tags User Controller
// @Param id_user path string true "Id User"
// @Param photo formData file true "Photo"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /UserAPI/{id_user}/UploadImage [post]
func UploadImage(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["id_user"], 10, 64)

	find, err := services.FindById(db, userId)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// Set Directory
	dir, err := os.Getwd()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	folderLocation := filepath.Join(dir, "images/profiles")

	if _, err := os.Stat(folderLocation); os.IsNotExist(err) {
		os.MkdirAll(folderLocation, 0700)
	}

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("photo")

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer file.Close()

	contentType := handler.Header.Get("Content-Type")

	if !CheckType(contentType) {
		response.ResponseError(w, http.StatusInternalServerError, errors.New("Format file tidak didukung"))
		return
	}

	var img image.Image

	switch (contentType) {
	case "image/jpg" :
		img, err = jpeg.Decode(file)
	case "image/jpeg" :
		img, err = jpeg.Decode(file)
	case "image/png" :
		img, err = png.Decode(file)
	case "image/gif" :
		img, err = gif.Decode(file)
	}

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	size := resize.Resize(300, 300, img, resize.Lanczos3)

	// Retrieve file information
	randomString := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)

	for i := range b {
		b[i] = randomString[rand.Intn(len(randomString))]
	}

	filename := handler.Filename
	filename = fmt.Sprintf("%s%s", string(b), filepath.Ext(handler.Filename))

	fileLocation := ""

	if find.Photo != "" {
		fileLocation = filepath.Join(dir, find.Photo)
	}

	exist, err := os.Stat(fileLocation)
	if exist != nil {
		e := os.Remove(fileLocation)
		if e != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
	}

	out, err := os.Create(folderLocation + `/` + filename)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer out.Close()

	switch (contentType) {
	case "image/jpg" :
		err = jpeg.Encode(out, size, nil)
	case "image/jpeg" :
		err = jpeg.Encode(out, size, nil)
	case "image/png" :
		err = png.Encode(out, size)
	case "image/gif" :
		err = gif.Encode(out, size, nil)
	}

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if _, err := io.Copy(out, file); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	get, err := services.UploadImage(db, fmt.Sprintf("/images/profiles/%s", filename), userId)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil upload image", get, http.StatusOK)
}

func CheckType(contentType string) bool {
	return contentType == "image/png" || contentType == "image/jpeg" || contentType == "image/jpg" || contentType == "image/gif"
}