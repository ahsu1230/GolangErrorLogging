package main

import (
	"errors"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type hero struct {
    Name     string    `json:"name`
    HeroName  string   `json:"heroName"`
}

type appError struct {
    Code     int    `json:"code"`
    Message  string `json:"message"`
}

func JSONAppErrorReporter() gin.HandlerFunc {
    return jsonAppErrorReporterT(gin.ErrorTypeAny)
}

func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        detectedErrors := c.Errors.ByType(errType)
		log.Println("Handler:", detectedErrors)
		return
    }
}

func main() {
    router := gin.Default()
    router.Use(JSONAppErrorReporter())
	router.GET("/ping1", doPing1)
	router.GET("/ping2", doPing2)
	router.GET("/ping3", doPing3)

	router.GET("/pong1", doPongGet1)
	router.GET("/pong2", doPongGet2)
	router.POST("/pong1", doPongPost1)
    router.Run(":3000")
}

// Success pings
func doPing1(c *gin.Context) {
	c.Status(http.StatusOK)
}

func doPing2(c *gin.Context) {
	c.String(http.StatusOK, "success")
}

func doPing3(c *gin.Context) {
	hero := hero { "Steve", "Cap" }
	c.JSON(http.StatusOK, &hero)
}

// Failure pongs
func doPongGet1(c *gin.Context) {
	err := errors.New("Error -> String")
	c.Error(err)
	c.String(http.StatusInternalServerError, err.Error())
}

func doPongGet2(c *gin.Context) {
	err := errors.New("Error -> Abort")
	c.AbortWithError(http.StatusInternalServerError, err)
}

func doPongPost1(c *gin.Context) {
	err := errors.New("Error -> Abort")
	c.AbortWithError(http.StatusInternalServerError, err)
}