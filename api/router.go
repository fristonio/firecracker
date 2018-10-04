package api

import "github.com/gin-gonic/gin"

func initGinRouter() *gin.Engine {
	router := gin.New()

	// API routes group
	apiGroup := router.Group("/api")
	{
		// Deploy route group
		deployGroup := apiGroup.Group("/deploy")
		{
			deployGroup.POST("/all/:action", deployAllHandler)
			deployGroup.POST("/challenge/:id/:action", deployChallengeHandler)
			deployGroup.POST("/local/", deployLocalChallengeHandler)
		}

		// Status route group
		statusGroup := apiGroup.Group("/status")
		{
			statusGroup.GET("/challenge/:id", challengeStatusHandler)
			statusGroup.GET("/all/", statusHandler)
		}

		// Info route group
		infoGroup := apiGroup.Group("/info")
		{
			infoGroup.GET("/challenge/:id", challengeInfoHandler)
			infoGroup.GET("/available/", availableChallengeInfoHandler)
		}
	}

	return router
}