package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// API API
type API struct {
	Service Service
}

// ProvideAPI ProvideAPI
func ProvideAPI(service Service) API {
	return API{Service: service}
}

// Release Release
func (a *API) Release(ctx *gin.Context) {
	_release := a.Service.Release()
	ctx.JSON(http.StatusOK, _release)
}

// Health Health
func (a *API) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "UP"})
}
