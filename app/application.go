package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shakilbd009/go-utils-lib/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {

	mapUrls()
	logger.Info("about to start the application...")
	router.Run(":8081")
}
