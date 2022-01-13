package test

import (
	"fmt"
	"github.com/yufuid/sso-go-sdk/auth"
	"os"
	"testing"
)

var idpAuth auth.YufuAuth

var redirectUrl = "https://demo-idp.cig.tencentcs.com/cidp/custom/ai-1b093de96646414a91652e6530b99c01"

func init() {
	//自定义参数示例
	var customClaims = make(map[string]interface{})
	customClaims["xxx"] = "xxx"

	builder := &auth.YufuAuthBuilder{
		YufuAuth: &auth.YufuAuth{},
	}

	yufuAuth, err := builder.
		Subject("siyuanzhang@yufuid.com").
		PrivateKeyPath("testPrivateKey.pem").
		Issuer("yufu").
		SDKRole(auth.IDP).
		//Claims(customClaims).
		Build()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	idpAuth = *yufuAuth
}

func Test_IdpGenerateToken(t *testing.T) {
	token, err := auth.GenerateToken(&idpAuth)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(token)
}

func Test_IdpGenerateUrl(t *testing.T) {
	url, err := auth.GenerateIDPRedirectUrl(redirectUrl, &idpAuth)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(url)
}
