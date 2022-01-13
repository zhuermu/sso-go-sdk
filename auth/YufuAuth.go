package auth

import (
	"errors"
	"strings"
)

const (
	TENANT_ID_KEY           = "tnt_id"
	AUDIENCE_YUFU           = "yufu"
	TOKEN_EXPIRE_TIME_IN_MS = 300000
)

const (
	// SP 服务提供商
	SP = "SP"

	// IDP 身份提供商
	IDP = "IDP"
)

type YufuAuth struct {
	Issuer         string
	PublicKeyPath  string
	PrivateKeyPath string
	KeyFingerPrint string
	TenantId       string
	SDKRole        string
	Audience       string
	Subject        string
	Claims         map[string]interface{}
}

type YufuAuthBuilder struct {
	YufuAuth *YufuAuth
}

func (auth *YufuAuthBuilder) Issuer(issuer string) *YufuAuthBuilder {
	auth.YufuAuth.Issuer = issuer
	return auth
}

func (auth *YufuAuthBuilder) PublicKeyPath(publicKeyPath string) *YufuAuthBuilder {
	auth.YufuAuth.PublicKeyPath = publicKeyPath
	return auth
}

func (auth *YufuAuthBuilder) PrivateKeyPath(privateKeyPath string) *YufuAuthBuilder {
	auth.YufuAuth.PrivateKeyPath = privateKeyPath
	return auth
}

func (auth *YufuAuthBuilder) KeyFingerPrint(keyFingerPrint string) *YufuAuthBuilder {
	auth.YufuAuth.KeyFingerPrint = keyFingerPrint
	return auth
}

func (auth *YufuAuthBuilder) TenantId(tenantId string) *YufuAuthBuilder {
	auth.YufuAuth.TenantId = tenantId
	return auth
}

func (auth *YufuAuthBuilder) SDKRole(sdkRole string) *YufuAuthBuilder {
	auth.YufuAuth.SDKRole = sdkRole
	return auth
}

func (auth *YufuAuthBuilder) Audience(audience string) *YufuAuthBuilder {
	auth.YufuAuth.Audience = audience
	return auth
}

func (auth *YufuAuthBuilder) Claims(claims map[string]interface{}) *YufuAuthBuilder {
	auth.YufuAuth.Claims = claims
	return auth
}

func (auth *YufuAuthBuilder) Subject(subject string) *YufuAuthBuilder {
	auth.YufuAuth.Subject = subject
	return auth
}
func (auth *YufuAuthBuilder) Build() (*YufuAuth, error) {

	result := *auth.YufuAuth

	if 0 == strings.Compare(result.SDKRole, SP) {
		if result.TenantId == "" {
			return nil, errors.New("tenant must be set with SP Role")
		}

		if result.Issuer == "" {
			return nil, errors.New("issuer must be set with SP Role")
		}

		if result.PublicKeyPath == "" {
			return nil, errors.New("public Key must set with SP Role")
		}

		return &result, nil

	} else if 0 == strings.Compare(result.SDKRole, IDP) {
		if result.PrivateKeyPath == "" {
			return nil, errors.New("private Key must be set with IDP Role")
		}

		return &result, nil
	}

	return nil, errors.New("must select a type in the SP or IDP")

}
