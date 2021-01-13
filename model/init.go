package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

var DB *gorm.DB

func init() {

	DB, _ = gorm.Open("mysql", "root:123456@tcp(192.168.20.58:3306)/data_model?charset=utf8&parseTime=True&loc=Local")
	DB.LogMode(true)

	//if err != nil {
	//	fmt.Println("连接失败")
	//	fmt.Println(err)
	//}
	fmt.Println(DB)
	fmt.Println("连接成功！！！")
	DB.AutoMigrate(&Product{})

}

func CreateData(c *gin.Context) {
	// createTodo 添加一条数据
	json := Product{}
	c.BindJSON(&json)
	fmt.Println(json)
	DB.Create(&json)
	log.Printf("%v", &json)
	fmt.Println(json)
	c.JSON(http.StatusOK, gin.H{
		"Code":  json.Code,
		"Price": json.Price,
	})
}

func GetData(c *gin.Context) {
	id := c.Param("id")
	ds := DB.First(&Product{}, "code = ?", id)
	fmt.Println(ds)
	c.JSON(http.StatusOK, gin.H{
		"data":   ds,
		"status": http.StatusOK,
	})

}
