package jwk

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/Eric-GreenComb/contrib/uuid"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	"github.com/sirupsen/logrus"

	"sme-jwk/internal/jwt"
)

// Service Service
type Service struct {
	ChainCache *cache.ChainCache
}

// ProvideService ProvideService
func ProvideService(chainCache *cache.ChainCache) Service {
	return Service{ChainCache: chainCache}
}

// GetJwk GetJwk
func (s *Service) GetJwk(kid string) (map[string]interface{}, error) {

	_keys, err := s.ChainCache.Get("sme.jwks." + kid)
	if err != nil {
		logrus.Println("ChainCache Get", err.Error())
		return nil, err
	}

	var _map map[string]interface{}
	err = json.Unmarshal([]byte(_keys.(string)), &_map)
	if err != nil {
		logrus.Println("JsonToMap err: ", err)
		return nil, err
	}

	return _map, nil
}

// CreateJwk CreateJwk
func (s *Service) CreateJwk(jwk jwt.JWK) error {

	err := s.ChainCache.Set("sme.jwks."+jwk.KID, jwk.Keys, &store.Options{})
	if err != nil {
		logrus.Println("ChainCache Set", err.Error())
		return err
	}

	return nil
}

// Signin Signin
func (s *Service) Signin(user jwt.User) (string, error) {

	_jwt := jwt.JWT{}

	_cacheKey, err := s.ChainCache.Get("sme.jwks." + user.KID)
	if err != nil {
		logrus.Println(err)
		return "", err
	}
	var _keys jwt.JWKKeys
	err = json.Unmarshal([]byte(_cacheKey.(string)), &_keys)
	if err != nil {
		logrus.Println(err)
		return "", err
	}

	_decoded, err := base64.RawURLEncoding.DecodeString(_keys.Keys[0].K)
	if err != nil {
		logrus.Println(err)
		return "", err
	}
	logrus.Println("decode base64:", string(_decoded))
	_jwt.SigningKey = _decoded

	_claims := jwt.CustomClaims{}
	_claims.Issuer = "gopper.in"
	_claims.UID = "13810167616"
	_claims.Name = "eric"
	_claims.IssuedAt = time.Now().Unix()
	// _claims.ExpiresAt = time.Now().Add(8 * time.Hour).Unix()
	_claims.ExpiresAt = time.Now().Add(365 * 24 * time.Hour).Unix()
	_claims.Subject = "13810167616"
	_claims.Jti = uuid.UUID()

	_jwtToken, err := _jwt.CreateToken(user.KID, _claims)
	if err != nil {
		return "", err
	}

	return _jwtToken, nil
}

// VerifyToken VerifyToken
func (s *Service) VerifyToken(user jwt.User) (bool, error) {

	_jwt := jwt.JWT{}
	err := _jwt.VerifyToken(user.Token, user.Key)
	if err != nil {
		logrus.Println(err)
		return false, err
	}

	return true, nil
}
