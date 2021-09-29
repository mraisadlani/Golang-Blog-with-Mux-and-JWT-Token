package utils

import (
	"fmt"
	"go-blog-jwt-token/api/payloads/request"
	"net/http"
	"strconv"
	"strings"
)

func SetupPagination(r *http.Request, set interface{}, pagination *request.Pagination, totalPages int) error {
	var data = set.(*request.Pagination)

	// get url path
	urlPath := r.URL.Path

	// search query
	searchQuery := ""

	for _, search := range pagination.Searchs {
		searchQuery += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	// set total pages
	if data.Limit == 0 {
		data.TotalPages = 0
	} else {
		data.TotalPages = totalPages
	}

	// set first & last page
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 0, pagination.Sort) + searchQuery
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.TotalPages, pagination.Sort) + searchQuery

	if data.Page > 0 {
		// set previous page
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page-1, pagination.Sort) + searchQuery
	}

	if data.Page < totalPages {
		// set next page
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort) + searchQuery
	}

	if data.Page > totalPages {
		// reset previous page
		data.PreviousPage = ""
	}

	return nil
}

func SortPagination(r *http.Request) (*request.Pagination, error) {
	limit := 0
	page := 1
	sort := "create_at asc"

	var searchs []request.Search

	query := r.URL.Query()

	for key, val := range query {
		qVal := val[len(val) - 1]

		switch key {
		case "limit" :
			limit, _ = strconv.Atoi(qVal)
			break
		case "page":
			page, _ = strconv.Atoi(qVal)
			break
		case "sort":
			sort = qVal
			break
		}

		if strings.Contains(key, ".") {
			// split query parameter
			searchKey:= strings.Split(key, ".")

			// create search object
			search := request.Search{Column: searchKey[0], Action: searchKey[1], Query: qVal}

			searchs = append(searchs, search)
		}
	}

	paginations := &request.Pagination{Limit: limit, Page: page, Sort: sort, Searchs: searchs}

	return paginations, nil
}