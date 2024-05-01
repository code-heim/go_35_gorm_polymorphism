package main

import (
	"gin_blogs/app_config"
	"gin_blogs/controllers"
	"gin_blogs/middlewares"
	"gin_blogs/models"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()

	// Init session store
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("blogs", store))

	r.Use(gin.Logger())

	r.Static("/public", "./public")

	r.LoadHTMLGlob("templates/**/*")

	app_config.ReadConfig()
	models.ConnectDatabase()
	models.DBMigrate()

	blogs := r.Group("/blogs", middlewares.AuthMiddleware())
	{
		blogs.GET("/", controllers.BlogsIndex)
		blogs.GET("/:id/comments", controllers.CommentsList)
		blogs.POST("/:id/comments", controllers.CommentsCreate)
		blogs.GET("/:id", controllers.BlogsShow)
	}

	photos := r.Group("/photos", middlewares.AuthMiddleware())
	{
		photos.GET("/", controllers.PhotosIndex)
		photos.GET("/:id/comments", controllers.CommentsList)
		photos.POST("/:id/comments", controllers.CommentsCreate)
		photos.GET("/:id", controllers.PhotosShow)
	}

	r.GET("/users/signup", controllers.SignupPage)
	r.GET("/users/login", controllers.LoginPage)
	r.POST("/users/signup", controllers.Signup)
	r.POST("/users/login", controllers.Login)
	r.DELETE("/users/logout", controllers.Logout)

	r.GET("/test", controllers.Test)

	log.Println("Server started!")
	r.Run() // Default Port 8080
}
