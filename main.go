package main

import (
	"github.com/MatthewTran22/ProjectPlanner/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", controllers.Get)
	r.GET("/data", controllers.GetData)

	r.Run(":8000")

}
