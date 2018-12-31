package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/beastv4/core"
	cfg "github.com/sdslabs/beastv4/core/config"
)

// Returns port in use by beast.
// @Summary Returns ports in use by beast by looking in the hack git repository, also returns min and max value of port allowed while specifying in beast challenge config.
// @Description Returns the ports in use by beast, which cannot be used in creating a new challenge..
// @Tags info
// @Accept  json
// @Produce application/json
// @Success 200 {object} api.PortsInUseResp
// @Failure 402 {object} api.PortsInUseResp
// @Router /api/info/ports/used [get]
func usedPortsInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, PortsInUseResp{
		MinPortValue: core.ALLOWED_MIN_PORT_VALUE,
		MaxPortValue: core.ALLOWED_MAX_PORT_VALUE,
		PortsInUse:   cfg.USED_PORTS_LIST,
	})
}

func challengeInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": WIP_TEXT,
	})
}

func availableChallengeInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": WIP_TEXT,
	})
}

// Returns available base images.
// @Summary Gives all the base images that can be used while creating a beast challenge, this is a constant specified in beast global config
// @Description Returns all the available base images  which can be used for challenge creation as the base OS for challenge.
// @Tags info
// @Accept  json
// @Produce application/json
// @Success 200 {object} api.AvailableImagesResp
// @Failure 402 {object} api.AvailableImagesResp
// @Router /api/info/images/available [get]
func availableImagesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, AvailableImagesResp{
		Message: "Available base images are",
		Images:  cfg.Cfg.AllowedBaseImages,
	})
}
