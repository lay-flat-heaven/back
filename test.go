package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/",func(c *gin.Context) {
		c.JSON(202,gin.H{
			"hello":"helllo",
		})
	})
	router.POST("/upload", func(c *gin.Context) {
	
		file, _ := c.FormFile("file")

		fmt.Println(file.Filename)

		
		 c.SaveUploadedFile(file, "/home/shinoshina/gocode/src/cback/kk.txt")
		

		c.String(202, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")

}
