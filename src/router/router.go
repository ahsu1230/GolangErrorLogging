
package router

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/ahsu1230/golangwebservertutorial/src/entities"
	"github.com/ahsu1230/golangwebservertutorial/src/logger"
	"github.com/ahsu1230/golangwebservertutorial/src/services"
)

func Setup() *gin.Engine {
	engine := gin.Default()
	// Write to a file
	// file, _ := os.Create("request.log")
	// gin.DefaultWriter = io.MultiWriter(file)
	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	
	engine.Use(AppRequestHandler())
	engine.NoRoute(func(c *gin.Context) {
		logger.Info("API endpoint unrecognized by handler", logger.Fields{})
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Unrecognized path",
		})
	})
	engine.GET("/ping1", services.GetPing1)
	engine.GET("/ping2", services.GetPing2)
	engine.GET("/ping3", services.GetPing3)
	engine.GET("/ping4", services.GetPing4)

	engine.GET("/pong1", services.GetPong1)
	engine.GET("/pong2", services.GetPong2)
	engine.GET("/pong3", services.GetPong3)
	engine.GET("/pong4", services.GetPong4)

	return engine
}


func AppRequestHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
		requestUuid := uuid.New()

		logger.Info("Received Request", logger.Fields{
			"requestUuid": requestUuid,
			"requestMethod": c.Request.Method,
			"requestURL": c.Request.URL,
			// "requestHost": c.Request.Host,
			// "requestHeader": c.Request.Header,
			// "requestBody": c.Request.Form,
		})
		c.Set("requestUuid", requestUuid)
		c.Writer.Header().Set("X-Request-Id", requestUuid.String())

		c.Next()

		if (len(c.Errors) > 0) {
			appError := createAppErrorFromResponseErrors(c)
			logger.Error(appError.Message, appError.Error, logger.Fields{
				"code": appError.Code,
				"requestUuid": requestUuid,
			})
			c.AbortWithStatusJSON(appError.Code, &appError)
			return
		}

		if (!c.IsAborted()) {
			logger.Info("Succesfully completed request!", logger.Fields{
				"requestUuid": requestUuid,
				"fullPath": c.FullPath(),
				"status": c.Writer.Status(),
			})
		}
	}
}

func createAppErrorFromResponseErrors(c *gin.Context) entities.AppError {
	// For now, assume only one Error in response

	wrappedErr := c.Errors[0].Err
	logger.Info("Handling err...", logger.Fields{
		"wrapped": wrappedErr,
		"cause": errors.Cause(wrappedErr),
		"isSQL": errors.Is(wrappedErr, entities.ErrSQL),
		"isRepo": errors.Is(wrappedErr, entities.ErrRepo),
		"isCtrl": errors.Is(wrappedErr, entities.ErrCtrl),
	})

	message := "An unknown error occured"
	if (errors.Is(wrappedErr, entities.ErrSQL)) {
		message = "A database-related error occured"
	} else if (errors.Is(wrappedErr, entities.ErrRepo)) {
		message = "A repo error occured"
	} else if (errors.Is(wrappedErr, entities.ErrCtrl)) {
		message = "A controller error occured"
	}

	return entities.AppError {
		http.StatusInternalServerError,
		message,
		wrappedErr,
	}
}