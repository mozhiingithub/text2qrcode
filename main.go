package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("template/*")

	// 注册图片handler
	r.GET("/img", imgHandler)

	// 注册首页handler
	r.GET("/", indexHandler)

	// 注册接收文本handler
	r.POST("/", uploadHandler)

	// 运行
	r.Run()
}

// 根据url的query，生成相应文本的二维码图片
func imgHandler(c *gin.Context) {

	// 解析url，获取text
	text := c.Query("text")
	log.Printf("text:%s\n", text)

	// 调用go-qrcode包，生成二维码的字节流
	img, err := qrcode.Encode(text, qrcode.Medium, 256)
	if nil != err {
		log.Println(err)
	}

	// 将图片写入context
	c.Writer.Write(img)
}

// 主页handler
func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// 接收文本
func uploadHandler(c *gin.Context) {

	// 通过POST方法获取文本
	text := c.PostForm("text")

	// 访问imgHandler的网址
	imgURL := "img/?text=" + text

	// html模板各项数据的map
	m := gin.H{
		"text":   text,
		"imgURL": imgURL,
	}

	// 转至返回界面
	c.HTML(http.StatusOK, "result.html", m)
}
