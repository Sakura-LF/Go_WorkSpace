package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("Gin/HTML/index.html")

	//r.GET("/", func(context *gin.Context) {
	//	context.HTML(200, "index.html", nil)
	//})

	r.Any("/")

	r.POST("/Upload", Savefile)

	r.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "没有此路由")
	})
	r.Run()
}

func Savefile(context *gin.Context) {
	//获取前端传入的文件: 有两个返回值File,FileError
	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	fmt.Println(file.Filename)

	//加入时间戳
	time_int := time.Now().Unix()
	time_str := strconv.FormatInt(time_int, 10)

	//将文件保存到本地:在文件名后面拼接时间戳

	if err := context.SaveUploadedFile(file, "Gin/img/"+time_str+file.Filename); err != nil {
		return
	}

}
