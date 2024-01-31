package main

import (
	"ks/handler"

	"github.com/gin-gonic/gin"
)
func main()  {
	router:=gin.Default()
	router.GET("/hi",handler.Hi)
	router.POST("/signup",handler.SignUp)
	router.POST("/login",handler.Login)
	router.POST("/product",handler.AddProduct)
	router.GET("/product",handler.ListProducts)
	router.DELETE("/product",handler.DeleteProduct)
	router.Run(":8080")
}