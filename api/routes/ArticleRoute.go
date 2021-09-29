package routes

import (
	"go-blog-jwt-token/api/controllers"
	"net/http"
)

var articleRoute = []Route{
	Route{
		URI: "/ArticleAPI/",
		Method: http.MethodGet,
		Handler: controllers.GetAllArticle,
		Auth: true,
	},
	Route{
		URI: "/ArticleAPI/{id_article}/GetArticle",
		Method: http.MethodGet,
		Handler: controllers.GetArticle,
		Auth: true,
	},
	Route{
		URI: "/ArticleAPI/CreateArticle",
		Method: http.MethodPost,
		Handler: controllers.CreateArticle,
		Auth: true,
	},
	Route{
		URI: "/ArticleAPI/{id_article}/UpdateArticle",
		Method: http.MethodPost,
		Handler: controllers.UpdateArticle,
		Auth: true,
	},
	Route{
		URI: "/ArticleAPI/{id_article}/DeleteArticle",
		Method: http.MethodPost,
		Handler: controllers.DeleteArticle,
		Auth: true,
	},
}