package main

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(filepath.Dir("/home/vagrant"))

	a := 8 << 2
	fmt.Println(a)

	var aa uint = 60 /* 60 = 0011 1100 */
	var c uint

	c = aa << 2 /* 240 = 1111 0000 */ /* 60*2*2 */
	fmt.Printf("第四行 - c 的值为 %d\n", c)

	router := gin.Default()
	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		// firstname := c.DefaultQuery("firstname", "Guest")
		// lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.SetCookie("gin", "test", 60, "/", "localhost", true, true)
		// c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
		c.JSON(200, gin.H{"status": "you are logged in"})
		// data := make(map[string]string)
		// data["firstname"] = firstname
		// data["lastname"] = lastname

		// json := struct {
		// 	Code      int    `json:"code"`
		// 	Firstname string `json:"firstname"`
		// 	Lastname  string `json:"lastname"`
		// }{http.StatusOK, firstname, lastname}
		// c.AsciiJSON(http.StatusOK, json)
	})

	router.Run(":8000")
}
