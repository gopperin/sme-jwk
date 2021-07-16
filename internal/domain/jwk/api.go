package jwk

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sme-jwk/internal/jwt"
)

// API API
type API struct {
	Service Service
}

// ProvideAPI ProvideAPI
func ProvideAPI(service Service) API {
	return API{Service: service}
}

// GetJwk GetJwk
func (a *API) GetJwk(ctx *gin.Context) {

	_kid := ctx.Param("kid")
	_map, _ := a.Service.GetJwk(_kid)
	ctx.JSON(http.StatusOK, _map)
	return

}

// CreateJwk CreateJwk
func (a *API) CreateJwk(ctx *gin.Context) {
	var _jwk jwt.JWK
	ctx.ShouldBindJSON(&_jwk)

	err := a.Service.CreateJwk(_jwk)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 500, "data": _jwk})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": 200, "data": _jwk})
}

// Signin Signin
func (a *API) Signin(ctx *gin.Context) {
	var _user jwt.User
	ctx.ShouldBindJSON(&_user)

	_jwt, err := a.Service.Signin(_user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 500, "data": ""})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": 200, "data": _jwt})
}

// VerifyToken VerifyToken
func (a *API) VerifyToken(ctx *gin.Context) {
	var _user jwt.User
	ctx.ShouldBindJSON(&_user)

	_bool, err := a.Service.VerifyToken(_user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 500, "data": _bool})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": 200, "data": _bool})
}
