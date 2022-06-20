package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func main(){


	u := gin.Default()
	fmt.Print(u)
	var a *gorm.DB
	fmt.Print(a)
}