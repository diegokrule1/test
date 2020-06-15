package service1

import (
	b64 "encoding/base64"
	"io/ioutil"
	"log"
	"net/http"

	r "github.com/diegokrule1/test/app/repository"
	"github.com/gin-gonic/gin"
)

type User struct {
	Id    string
	Title string
}

func GetUsers(c *gin.Context) {
	var (
		id    string
		title string
	)
	id, title = r.GetUsers()
	u := User{}
	u.Id = id
	u.Title = title

	c.JSON(http.StatusOK, u)
}

func Hello() string {
	return "Hello, world from go."
}

func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func GetRemote(c *gin.Context) {
	resp, err := http.Get("http://localhost:8080/api/mytest")
	if err != nil {
		log.Fatal("Could not invoke service")
		c.JSON(500, "could not process request")
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Could not invoke service")
			c.JSON(500, "could not process request")
		} else {
			sDec, _ := b64.StdEncoding.Decode(body)
			log.Print(string(body))
			c.JSON(200, body)
		}
	}
}
