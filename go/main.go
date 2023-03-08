package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const URL = "https://dummyjson.com/todo/1"

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "World!")
	})

	r.GET("/performcall", func(c *gin.Context) {
		response, err := http.Get(URL)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error getting data from external service",
			})
			return
		}

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error getting data from external service",
			})
			return
		}

		jsonMap := make(map[string](interface{}))
		err = json.Unmarshal([]byte(responseData), &jsonMap)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error getting data from external service",
			})
			return
		}

		c.JSON(http.StatusOK, jsonMap)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
