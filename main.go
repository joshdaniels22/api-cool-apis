package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/pokemon", func(c *gin.Context) {
		pokemon := "pikachu"
		requestURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
		res, err := http.Get(requestURL)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			os.Exit(1)
		}

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("client: could not read response body: %s\n", err)
			os.Exit(1)
		}

		var result map[string]interface{}

		json.Unmarshal(resBody, &result)

		c.JSON(http.StatusOK, result)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
