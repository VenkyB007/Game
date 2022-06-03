package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//main data structure
type Data struct {
	Name      string      `json:"name"`
	Character []Character `json:"character"`
}

//character data structure
type Character struct {
	Name      string `json:"name"`
	Max_Power int    `json:"max_power"`
}

func main() {
	r := gin.Default()
	var data []Data
	go r.GET("/avengers", func(c *gin.Context) {

		var avengers Data

		resp, _ := http.Get("http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b")
		body, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(body, &avengers)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, avengers)
		c.JSON(http.StatusOK, avengers)
	})

	go r.GET("/mutants", func(c *gin.Context) {

		var mutants Data

		resp, _ := http.Get("http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e")
		body, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(body, &mutants)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, mutants)
		c.JSON(http.StatusOK, mutants)
	})
	go r.GET("/get-all", func(c *gin.Context) {
		c.JSON(http.StatusOK, data)

	})

	r.Run(":9009")
}
