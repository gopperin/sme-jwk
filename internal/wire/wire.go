//+build wireinject

package wire

import (
	"github.com/eko/gocache/cache"
	"github.com/google/wire"

	"sme-jwk/internal/domain/base"
	"sme-jwk/internal/domain/jwk"
)

// InitBaseAPI init base api wire
func InitBaseAPI() base.API {
	wire.Build(base.ProvideAPI, base.ProvideService)
	return base.API{}
}

// InitJWKAPI init jwk api wire
func InitJWKAPI(chaincache *cache.ChainCache) jwk.API {
	wire.Build(jwk.ProvideAPI, jwk.ProvideService)
	return jwk.API{}
}
