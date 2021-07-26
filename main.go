/**
 * @Author: zhangc
 * @Date: 2021/7/19 13:48
 * @LastEditors: zhangc
 * @FilePath: /main.go
 * @Description:
 * @Contactme: zhangchun34582@hundsun.com
**/

package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	ip := flag.String("ip", "", "resp")

	flag.Parse()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"ip":      *ip,
		})
	})
	r.Run(":8080")
}
