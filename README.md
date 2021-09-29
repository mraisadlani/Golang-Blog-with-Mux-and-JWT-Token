# Go Blog with Mux and MySQL
Membuat sebuah studi kasus pembuatan Blog menggunakan Golang dan MySQL sebagai media penyimpanannya

# How to using swagger
1. Install generator swagger menggunakan perintah :
```sh
go get -u github.com/swaggo/swag/cmd/swag
```
2. tambahkan module swagger menggunakan perintah :
```sh
go get -u github.com/swaggo/http-swagger
go get -u github.com/alecthomas/template
```
3. Kemudian, tambahkan General API pada file main.go
 ```sh
// @title Go Blog
// @version 1.0
// @description Service Blog
// @termsOfService http://swagger.io/terms/

// @contact.name Muhammad Rais Adlani
// @contact.url https://gitlab.com/mraisadlani
// @contact.email mraisadlani@gmail.com

// @license.name MIT
// @license.url https://gitlab.com/mraisadlani/golang-blog-with-mux-and-jwt-token/-/blob/main/LICENSE

// @securityDefinitions.apikey BearerAuth
// @in Header
// @name Authorization

// @host localhost:9876
// @BasePath /api
 ```
atau ingin menggunakan dynamic general API dengan cara seperti ini:
```sh
docs.SwaggerInfo.Title = "Go Blog"
docs.SwaggerInfo.Description = "Service Blog"
docs.SwaggerInfo.Version = "1.0"
docs.SwaggerInfo.Host = fmt.Sprintf("%s:%v", "localhost", configs.Config.Server.PORT)
docs.SwaggerInfo.BasePath = "/api"
docs.SwaggerInfo.Schemes = []string{"http", "https"}
```

4. Tambahkan API Operation didalam kode controller seperti ini :
# Untuk Penggunaan Method GET
```sh
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
```

# Untuk Penggunaan Method POST
```sh
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
```

# Untuk Penggunaan Method PUT
```sh
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
// @Router /UserAPI/{id_user}/UpdateUser [put]
```

# Untuk Penggunaan Method DELETE
```sh
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
// @Router /UserAPI/{id_user}/DeleteUser [delete]
```

4. Untuk menjalankan generator gin-swagger menggunakan perintah :
```sh
swag init -g main.go
```

5. Jalankan kembali aplikasi menggunakan perintah :
```sh
go run main.go
```

6. Kemudian bisa diakses melalui url :
```sh
http://localhost:9000/swagger/index.html
```

# How to run Application
clone the repository :
```sh
https://gitlab.com/mraisadlani/golang-blog-with-mux-and-jwt-token.git
```

Run application:
```sh
go run main.go
```

# Todolist Application Blog
- [x] REST API User
  - [x] Create User
  - [x] Get All Users (Pagination)
  - [x] Find By Id User
  - [x] Update User
  - [x] Delete User
  - [x] Upload Image
  - [x] Change Status
- [x] REST API Role
  - [x] Get All Role (Pagination)
  - [x] Create Role
  - [x] Find by Rolename
  - [x] Update Role
  - [x] Delete Role
  - [x] Change Status
- [x] REST API Permission
  - [x] Create Permission
  - [x] Update Permission
  - [x] Delete Permission
  - [x] Get Permission user
- [x] REST API Menu and Sub Menu
  - [x] Create Menu & Submenu
  - [x] Update Menu & Submenu
  - [x] Delete Menu & Submenu
  - [x] Get All Menu & Submenu (Pagination)
  - [x] Find Menu & Submenu
  - [x] Change Status Menu & Submenu
- [x] REST API Article
  - [x] Create Article
  - [x] Update Article
  - [x] Delete Article
  - [x] Get All Article (Pagination)
  - [x] Get By Id Article
- [x] REST API Auth
  - [x] Do Login
  - [x] Forget Password
- [x] REST API Category
  - [x] Create Category
  - [x] Update Category
  - [x] Delete Category
  - [x] Get All Category (Pagination)
  - [x] Get By Id Category
- [x] REST API Tag
  - [x] Create Tag
  - [x] Update Tag
  - [x] Delete Tag
  - [x] Get All Tag (Pagination)
  - [x] Get By Id Tag
- [x] REST API Post
  - [x] Find Post
  - [x] Get All Post (Pagination)
  - [x] Create Post
  - [x] Update Post
  - [x] Delete Post
  - [x] Publish Post
  - [x] Cancel Post

# Library
- Mux : https://github.com/gorilla/mux
- Resize Image : https://github.com/nfnt/resize
- Environment : https://github.com/joho/godotenv
- Generate Slug : https://github.com/gosimple/slug
- JWT Token : https://github.com/dgrijalva/jwt-go
- Swagger : https://github.com/swaggo/gin-swagger
- Gorm & MySQL Driver : https://gorm.io/
- Validation Format Email : https://github.com/badoux/checkmail

# Contribute
Support saya agar lebih banyak berkontribusi dalam membuat sebuah project sederhana menggunakan bahasa pemrograman golang
- Saweria : https://saweria.co/mraisadlani
