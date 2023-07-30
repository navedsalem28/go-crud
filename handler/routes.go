package handler

import (
	"go-crud/common"
	"net/http"
)

type Route struct {
	Name    string
	Path    string
	Method  string
	Handler http.HandlerFunc
}

var Routes = []Route{
	{
		"Home Page",
		"/",
		"GET",
		common.Home,
	}, {
		"Users Home Page",
		"/user",
		"GET",
		common.User,
	}, {
		"Article Home Page",
		"/article",
		"GET",
		common.Article,
	}, {
		"Category Home Page",
		"/category",
		"GET",
		common.Category,
	},

	// All Users CRUD APIs Route
	{
		"Get All Users",
		"/Users",
		"GET",
		GetAllUser,
	}, {
		"Get Single User",
		"/Users/{id}",
		"GET",
		GetSingleUser,
	}, {
		"Create User",
		"/Users",
		"POST",
		CreateUser,
	}, {
		"Update User",
		"/Users/{id}",
		"PUT",
		UpdateUser,
	}, {
		"Delete User",
		"/Users/{id}",
		"DELETE",
		DeleteUser,
	},
	// All Articles CRUD APIs Route
	{
		"Get All Articles",
		"/Articles",
		"GET",
		GetAllArticles,
	}, {
		"Get Single Article",
		"/Articles/{id}",
		"GET",
		GetSingleArticle,
	}, {
		"Create Article",
		"/Articles",
		"POST",
		CreateArticle,
	}, {
		"Update Article",
		"/Articles/{id}",
		"PUT",
		UpdateArticle,
	}, {
		"Delete Article",
		"/Articles/{id}",
		"DELETE",
		DeleteArticle,
	},
	// All categories CRUD APIs Route
	{
		"Get All Categories",
		"/Categories",
		"GET",
		GetAllCategories,
	}, {
		"Get Single Category",
		"/Categories/{id}",
		"GET",
		GetSingleCategory,
	}, {
		"Create Category",
		"/Categories",
		"POST",
		CreateCategory,
	}, {
		"Update Category",
		"/Categories/{id}",
		"PUT",
		UpdateCategory,
	}, {
		"Delete Category",
		"/Categories/{id}",
		"DELETE",
		DeleteCategory,
	},
}
