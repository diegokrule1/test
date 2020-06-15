package controller

import (
	s "github.com/diegokrule1/test/app/service1"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:name", s.SayHello)

	r.GET("/users/list", s.GetUsers)

	r.GET("/remote", s.GetRemote)

	//r.GET("/speak", s.TalkToMamal)

	//r.POST("/user/:name", s.CreateUser)
	return r
}
