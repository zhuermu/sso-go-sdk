package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/zhuermu/sso-go-sdk/auth"
)

var token = "eyJraWQiOiJhaS00NTg1ZmQ2OGEzYTc0ODhkYWNlODEyNDViNjU4NzM1Zjpzc28iLCJhbGciOiJSUzI1NiJ9.eyJhcHBJbnN0YW5jZUlkIjoiYWktNDU4NWZkNjhhM2E3NDg4ZGFjZTgxMjQ1YjY1ODczNWYiLCJhdWQiOiJhaS00NTg1ZmQ2OGEzYTc0ODhkYWNlODEyNDViNjU4NzM1ZiIsInN1YiI6InNpeXVhbnpoYW5nQHl1ZnVpZC5jb20iLCJ0bnRfaWQiOiJ0bi0yN2VkMWUyYzEzZDQ0NzQ1YTZmODUwYzdhMjE5NmJkNSIsInNjb3BlIjoic3NvIiwiaXNzIjoieXVmdWlkLmNvbVwvYWktNDU4NWZkNjhhM2E3NDg4ZGFjZTgxMjQ1YjY1ODczNWYiLCJleHAiOjE2NDE5NzQ1ODgsImlhdCI6MTY0MTk3Mzk4OH0.K8SYc3wIds-DedD3Gfwi7fjTRpYm3oDOgx_t6FL-QCeT-O1FWorvO5eqY-NRfSMVyiIOZ3ciELf25oXyEmMsc1umGRwBkZGVJsjXOsPbbWSRDwa3NisHeIJuGpX8EbJUILkU2yA5t_mKcFe5rB-8-GJTznWsxI3nW183g_0hG3SBBzebMPsSTJP6ipmi3AIb5dg-tY9XmNUhiJ8tg9HF4XHAzOAWKELxza5mXHm0ZVxcLW1CiG4cuSXmE3aaIGHLU0-usxhpKM7mGUg39S5Sw2aP8jq52BXsQb1XQPeObnV9HvUTxJomRZEzGcrQVK8hg3nuoFPK_t4cLKwxSAYucA"

func Test_SpVerifyToken(t *testing.T) {
	builder := &auth.YufuAuthBuilder{
		YufuAuth: &auth.YufuAuth{},
	}

	spAuth, err := builder.
		TenantId("tn-yufu").
		PublicKeyPath("testPublicKey.pem").
		Issuer("yufuid.com/ai-xxx").
		SDKRole(auth.SP).
		Build()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	mapClaims, err := auth.VerifyToken(token, spAuth)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	tenantId := mapClaims["tnt_id"]
	username := mapClaims["sub"]

	fmt.Println(tenantId.(string) + "," + username.(string))
}
