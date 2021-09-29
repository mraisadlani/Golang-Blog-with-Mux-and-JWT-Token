package controllers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
	"github.com/nfnt/resize"
	"go-blog-jwt-token/api/configs"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/payloads/response"
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
	"strings"
	"time"
)

// @Summary Get All Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /PostAPI/GetAllPost [get]
func GetAllPost(w http.ResponseWriter, r *http.Request) {
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

	get, err := services.GetAllPost(r, db, pagination)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
}

// @Summary Create Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param requestBody formData request.PostRequest true "Form"
// @Param photo formData file true "Photo"
// @Param categories formData string true "Category"
// @Param tags formData string true "Tag"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /PostAPI/CreatePost [post]
func CreatePost(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	dir, err := os.Getwd()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	folderLocation := filepath.Join(dir, "images/posts")

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

	size := resize.Resize(600, 600, img, resize.Lanczos3)

	// Retrieve file information
	randomString := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)

	for i := range b {
		b[i] = randomString[rand.Intn(len(randomString))]
	}

	filename := handler.Filename
	filename = fmt.Sprintf("%s%s", string(b), filepath.Ext(handler.Filename))

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

	publish, _ := strconv.ParseBool(r.FormValue("published"))
	idArticle, _ := strconv.ParseInt(r.FormValue("id_article"), 10, 64)

	posts := request.PostRequest{
		NamaPost: r.FormValue("nama_post"),
		Description: r.FormValue("description"),
		Published: publish,
		IdArticle: idArticle,
		CreateBy: r.FormValue("create_by"),
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateBy: r.FormValue("update_by"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	urlImage := fmt.Sprintf("/images/posts/%s", filename)
	posts.Slug = slug.Make(posts.NamaPost)

	categories := strings.Split(r.FormValue("categories"), ",")
	tags := strings.Split(r.FormValue("tags"), ",")

	get, err := services.CreatePost(db, posts, urlImage, categories, tags)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)
}

// @Summary Find Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param id_article path string true "Id Article"
// @Param id_post path string true "Id Post"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /PostAPI/{id_article}/FindPost/{id_post} [get]
func FindPost(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	IdArticle, _ := strconv.ParseInt(vars["id_article"], 10, 64)
	IdPost, _ := strconv.ParseInt(vars["id_post"], 10, 64)

	get, err := services.FindPost(db, IdArticle, IdPost)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
}

// @Summary Update Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param requestBody formData request.PostRequest true "Form"
// @Param photo formData file false "Photo"
// @Param categories formData string true "Category"
// @Param tags formData string true "Tag"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /PostAPI/UpdatePost [post]
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	publish, _ := strconv.ParseBool(r.FormValue("published"))
	idPost, _ := strconv.ParseInt(r.FormValue("id_post"), 10, 64)
	idArticle, _ := strconv.ParseInt(r.FormValue("id_article"), 10, 64)

	posts := request.PostRequest{
		ID: idPost,
		NamaPost: r.FormValue("nama_post"),
		Description: r.FormValue("description"),
		Published: publish,
		IdArticle: idArticle,
		CreateBy: r.FormValue("create_by"),
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateBy: r.FormValue("update_by"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	find, err := services.FindPost(db, idArticle, idPost)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	file, handler, err := r.FormFile("photo")
	urlImage := ""

	if err == nil {


		dir, err := os.Getwd()
		folderLocation := filepath.Join(dir, "images/posts")

		if _, err := os.Stat(folderLocation); os.IsNotExist(err) {
			os.MkdirAll(folderLocation, 0700)
		}

		r.ParseMultipartForm(10 << 20)

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

		size := resize.Resize(600, 600, img, resize.Lanczos3)

		// Retrieve file information
		randomString := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b := make([]rune, 10)

		for i := range b {
			b[i] = randomString[rand.Intn(len(randomString))]
		}

		filename := handler.Filename
		filename = fmt.Sprintf("%s%s", string(b), filepath.Ext(handler.Filename))

		fileLocation := ""

		if find.Image != "" {
			fileLocation = filepath.Join(dir, find.Image)
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

		urlImage = fmt.Sprintf("/images/posts/%s", filename)
	}

	fmt.Println(urlImage)

	if urlImage == "" {
		fmt.Println(find.Image)
		urlImage = find.Image
	}

	posts.Slug = slug.Make(posts.NamaPost)

	categories := strings.Split(r.FormValue("categories"), ",")
	tags := strings.Split(r.FormValue("tags"), ",")

	get, err := services.UpdatePost(db, posts, urlImage, categories, tags)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)
}

// @Summary Delete Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param id_article path string true "Id Article"
// @Param id_post path string true "Id Post"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /PostAPI/{id_article}/DeletePost/{id_post} [post]
func DeletePost(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	IdArticle, _ := strconv.ParseInt(vars["id_article"], 10, 64)
	IdPost, _ := strconv.ParseInt(vars["id_post"], 10, 64)

	find, err := services.FindPost(db, IdArticle, IdPost)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	dir, err := os.Getwd()
	fileLocation := ""

	if find.Image != "" {
		fileLocation = filepath.Join(dir, find.Image)
	}

	exist, err := os.Stat(fileLocation)
	if exist != nil {
		e := os.Remove(fileLocation)
		if e != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
	}

	get, err := services.DeletePost(db, IdArticle, IdPost)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil menghapus data", get, http.StatusOK)
}

// @Summary Publish Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param id_article path string true "Id Article"
// @Param id_post path string true "Id Post"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /PostAPI/{id_article}/PublishPost/{id_post} [post]
func PublishPost(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	IdArticle, _ := strconv.ParseInt(vars["id_article"], 10, 64)
	IdPost, _ := strconv.ParseInt(vars["id_post"], 10, 64)

	get, err := services.PublishPost(db, IdArticle, IdPost)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

		response.ResponseMessage(w, "Berhasil mempublish data", get, http.StatusOK)
}

// @Summary Cancel Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param id_article path string true "Id Article"
// @Param id_post path string true "Id Post"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /PostAPI/{id_article}/CancelPost/{id_post} [post]
func CancelPost(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	IdArticle, _ := strconv.ParseInt(vars["id_article"], 10, 64)
	IdPost, _ := strconv.ParseInt(vars["id_post"], 10, 64)

	get, err := services.CancelPost(db, IdArticle, IdPost)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membatalkan data", get, http.StatusOK)
}