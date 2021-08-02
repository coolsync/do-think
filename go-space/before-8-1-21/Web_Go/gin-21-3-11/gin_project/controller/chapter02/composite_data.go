package chapter02

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


// struct render
func UserInfoStruct(c *gin.Context) {

	// u1 := UserInfo{Id: 1, Name: "bob", Age: 12, Addr: "xxxx"}

	var u2 UserInfo
	u2.Id = 2
	u2.Name = "paul"
	u2.Age = 22
	u2.Addr = "xxx"

	// m := make(map[string]interface{})
	// m["u2"] = u2;

	c.HTML(http.StatusOK, "chapter02/struct.html", u2)
}

// array render
func ArrayHandler(c *gin.Context) {

	sli := []int{1, 2, 3, 4, 5}

	c.HTML(http.StatusOK, "chapter02/array.html", sli)
}

// array and struct render
func ArrayAndStructHandler(c *gin.Context) {

	sliStr := []UserInfo{
		{Id: 1, Name: "bob", Age: 30, Addr: "xxx"},
		{Id: 2, Name: "bob2", Age: 31, Addr: "xxx2"},
		{Id: 3, Name: "bob3", Age: 32, Addr: "xxx3"},
	}

	c.HTML(http.StatusOK, "chapter02/array_struct.html", sliStr)
}

// map render
func MapHandler(c *gin.Context) {

	m1 := map[string]string{"name": "alice", "age": "18"}

	m2 := map[string]int{"id": 1}

	m3 := map[string]interface{}{"m1": m1, "m2": m2}

	c.HTML(http.StatusOK, "chapter02/map.html", m3)
}

// map and struct render
func MapAndStructHandler(c *gin.Context) {

	m1 := map[string]UserInfo{
		"user": {Id: 1, Name: "alice", Age: 18, Addr: "XXXX"},
	}

	c.HTML(http.StatusOK, "chapter02/map_struct.html", m1)
}

// route param
func Param1Handler(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "hello, %s", id)
}
func Param2Handler(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "hello, %s", id)
}
