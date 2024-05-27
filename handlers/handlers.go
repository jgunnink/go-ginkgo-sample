package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HelloHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	}
}

func HelloNameHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello, %s!", name),
		})
	}
}

func HelloCalculatorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		num1, err := strconv.Atoi(c.Query("num1"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid num1 parameter"})
			return
		}

		num2, err := strconv.Atoi(c.Query("num2"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid num2 parameter"})
			return
		}

		result := num1 + num2

		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

func HelloThirdPartyApiHandler(httpClient *http.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Query("url") // Get the URL from the query parameter
		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url parameter"})
			return
		}

		response, err := httpClient.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer response.Body.Close()

		responseData := make(map[string]interface{})
		json.NewDecoder(response.Body).Decode(&responseData)

		c.JSON(response.StatusCode, responseData)
	}
}
