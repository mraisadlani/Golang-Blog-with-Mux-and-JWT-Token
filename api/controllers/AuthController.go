package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go-blog-jwt-token/api/configs"
	"go-blog-jwt-token/api/payloads/request"
	"go-blog-jwt-token/api/payloads/response"
	"go-blog-jwt-token/api/securities"
	"go-blog-jwt-token/api/services"
	"net/http"
	"strconv"
)

// @Summary Do Login
// @Description REST API Auth
// @Accept  json
// @Produce  json
// @Tags Auth Controller
// @Param reqBody body request.LoginRequest true "Form Request"
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /AuthAPI/DoLogin [post]
func DoLogin(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	var loginRequest request.LoginRequest
	err = json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = loginRequest.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := services.CheckLogin(db, loginRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if get.ID > 0 {
		hashPwd := get.Password
		pwd := loginRequest.Password

		hash := securities.VerifyPassword(hashPwd, pwd)

		if hash == nil {
			token, err := securities.GenerateToken(get.Username)

			if err != nil {
				response.ResponseError(w, http.StatusInternalServerError, err)
				return
			}

			response.ResponseToken(w, "Login Berhasil", token, get, http.StatusOK)
		} else {
			response.ResponseError(w, http.StatusBadRequest, errors.New("Password tidak sesuai"))
			return
		}
	}
}

// @Summary Forget Password
// @Description REST API Auth
// @Accept  json
// @Produce  json
// @Tags Auth Controller
// @Param id_user path string true "Id User"
// @Param reqBody body request.ForgetPasswordRequest true "Form Request"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /AuthAPI/{id_user}/ForgetPassword [post]
func ForgetPassword(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	vars := mux.Vars(r)
	idUser, _ := strconv.ParseInt(vars["id_user"], 10, 64)

	var forgetRequest request.ForgetPasswordRequest
	err = json.NewDecoder(r.Body).Decode(&forgetRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	get, err := services.FindById(db, idUser)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if get.ID > 0 {
		hashPwd := get.Password
		newPwd := forgetRequest.NewPassword
		oldPwd := forgetRequest.OldPassword

		err = securities.VerifyPassword(hashPwd, oldPwd)

		if err == nil {
			if newPwd != oldPwd {
				hash, err := securities.HashPassword(newPwd)

				if err != nil {
					response.ResponseError(w, http.StatusInternalServerError, err)
					return
				}

				forgetRequest.NewPassword = hash

				set, err := services.ForgetPassword(db, forgetRequest, idUser)

				if err != nil {
					response.ResponseError(w, http.StatusInternalServerError, err)
					return
				}

				response.ResponseMessage(w, "Berhasil mengubah password", set, http.StatusOK)
			} else {
				response.ResponseError(w, http.StatusBadRequest, errors.New("Password baru tidak boleh sama"))
				return
			}
		} else {
			response.ResponseError(w, http.StatusBadRequest, errors.New("Password lama tidak sesuai"))
			return
		}
	}
}