package controllers

import (
	"go-blog-jwt-token/api/payloads/response"
	"net/http"
)

// @Summary Connected to REST API Blog
// @Description Gate REST API
// @Accept  json
// @Produce  json
// @Tags Welcome Controller
// @Success 200 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router / [get]
func Welcome(w http.ResponseWriter, r *http.Request) {
	response.ResponseMessage(w, "Connected to REST Blog", nil, 200)
}
