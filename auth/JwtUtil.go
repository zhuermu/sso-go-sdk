package auth

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"os"
	"time"
)

type JwtClaims struct {
	jwt.StandardClaims
	Claims map[string]string
}

func GetPrivateKey(path string) (*rsa.PrivateKey, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	pkBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)

	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetPublicKey(path string) (*rsa.PublicKey, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	pkBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(pkBytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// GenerateToken /**
func GenerateToken(auth *YufuAuth) (tokenString string, err error) {

	if auth.Audience == "" {
		auth.Audience = AUDIENCE_YUFU
	}

	if auth.TenantId != "" {
		if auth.Claims == nil {
			auth.Claims = make(map[string]interface{})
		}
		auth.Claims["tnt_id"] = auth.TenantId
	}

	claims := jwt.MapClaims{
		"aud": auth.Audience,
		"exp": time.Now().Add(time.Millisecond * TOKEN_EXPIRE_TIME_IN_MS).Unix(),
		"iss": auth.Issuer,
		"iat": time.Now().Unix(),
		"sub": auth.Subject,
	}

	for k, v := range auth.Claims {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	privateKey, err := GetPrivateKey(auth.PrivateKeyPath)

	if err != nil {
		return "", err
	}

	return token.SignedString(privateKey)
}

func GenerateIDPRedirectUrl(redirectUrl string, auth *YufuAuth) (url string, err error) {
	token, err := GenerateToken(auth)

	return redirectUrl + "?id_token=" + token, err
}

func VerifyToken(tokenStr string, auth *YufuAuth) (claims map[string]interface{}, err error) {

	if err != nil {
		return nil, err
	}

	if tokenStr == "" {
		return nil, errors.New("token must exist and could not be empty")
	}

	if auth.PublicKeyPath == "" {
		return nil, errors.New("PublicKey must exist and could not be empty")
	}

	key, err := GetPublicKey(auth.PublicKeyPath)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unexpected signing method:" + token.Header["alg"].(string))
		}

		return key, nil
	})

	if err != nil {
		return nil, err
	}

	mapClaims := token.Claims.(jwt.MapClaims)

	if auth.TenantId == "" || auth.TenantId != mapClaims[TENANT_ID_KEY] {
		return nil, errors.New("tenant_id is invalid")
	}

	if !mapClaims.VerifyIssuer(auth.Issuer, true) {
		return nil, errors.New("issuer is invalid")
	}

	if !mapClaims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, errors.New("token is expired")
	}

	return mapClaims, nil
}
