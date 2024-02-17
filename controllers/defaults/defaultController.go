package defaults

import (
	"fmt"
	"gormnote/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (con DefaultController) Index(ctx *gin.Context) {
	// ctx.String(http.StatusOK, "首页")
	userList := []models.User{}
	if err := models.DB.Find(&userList).Error; err != nil {
		fmt.Printf("err: %v\n", err)
	}

	ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		"userList": userList,
	})
}

func (con DefaultController) Add(ctx *gin.Context) {
	// 增加数据
	// 实例化一个结构体并将它添加到数据库里。
	user := models.User{
		Username: "hlry",
		Age:      20,
		Email:    "hlry@gmail.com",
		AddTime:  1708060539,
	}
	if err := models.DB.Create(&user).Error; err != nil {
		fmt.Println("数据添加失败！")
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("数据添加成功！")
	}

	// 查询数据，验证一下确实是添加进去了。
	userList := []models.User{}
	if err := models.DB.Find(&userList).Error; err != nil {
		fmt.Printf("err: %v\n", err)
	}

	ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		"userList": userList,
	})
}

func (con DefaultController) Delete(ctx *gin.Context) {
	// 删除数据
	// user := models.User{Id: 2} //指定删除的数据的id。
	// models.DB.Delete(&user)

	user := models.User{}
	models.DB.Where("id = ?", 2).Delete(&user) //查询到id=2的数据并删除。

	userList := []models.User{}
	if err := models.DB.Find(&userList).Error; err != nil {
		fmt.Printf("err: %v\n", err)
	}

	ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		"userList": userList,
	})
}

func (con DefaultController) Query(ctx *gin.Context) {
	userList := []models.User{}
	// 查询id>2的数据。
	// models.DB.Where("id > ?", 2).Find(&userList)

	// 查询id=3,4,6的数据，?是占位符，用切片打包要查询的数据的id。
	models.DB.Where("id in ?", []int{3, 4, 6}).Find(&userList)

	ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		"userList": userList,
	})
}

func (con DefaultController) Update(ctx *gin.Context) {
	user := models.User{}                 //先实例化一个空结构体。
	models.DB.Where("id = 4").Find(&user) //找到要修改的数据。
	//修改数据。
	user.Username = "kzh"
	user.Age = 19
	user.Email = "kzh@outlook.com"
	models.DB.Save(&user)

	userList := []models.User{}
	if err := models.DB.Find(&userList).Error; err != nil {
		fmt.Printf("err: %v\n", err)
	}

	ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		"userList": userList,
	})
}

func (con DefaultController) Query_bt(ctx *gin.Context) {

	articleList := []models.Article{}
	models.DB.Preload("ArticleCategary").Find(&articleList)

	ctx.HTML(http.StatusOK, "default/article.html", gin.H{
		"results": articleList,
	})
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"result": articleList,
	// })
}

func (con DefaultController) Query_hm(ctx *gin.Context) {

	articleCategaryList := []models.ArticleCategary{}
	models.DB.Preload("Article").Find(&articleCategaryList)

	// ctx.HTML(http.StatusOK, "default/article.html", gin.H{
	// 	"results": articleCategaryList,
	// })
	ctx.JSON(http.StatusOK, gin.H{
		"results": articleCategaryList,
	})
}

func (con DefaultController) Query_mtm(ctx *gin.Context) {
	articleList := []models.Article{}
	models.DB.Preload("Tags").Find(&articleList)
	// ctx.HTML(http.StatusOK, "default/article.html", gin.H{
	// 	"results": articleList,
	// })
	ctx.JSON(http.StatusOK, gin.H{
		"articleList": articleList,
	})

	tagsList := []models.Tags{}
	models.DB.Preload("Article").Find(&tagsList)
	// ctx.HTML(http.StatusOK, "default/article.html", gin.H{
	// 	"results": tagsList,
	// })
	ctx.JSON(http.StatusOK, gin.H{
		"tagsList": tagsList,
	})
}
