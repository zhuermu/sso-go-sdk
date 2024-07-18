玉符单点登录 SDK
======
玉符SDK集成了签署和验证JWT令牌的功能，使得身份提供者（IDP）和服务提供者（SP）只需要用很少的工作量就可以快速将玉符提供的单点登录等功能集成到现有的服务中。

## 单点登录SDK简介

* 作为服务提供者（SP）,可以使用玉符SDK验证JWT令牌的有效性（包括有效期、签名等），验证成功后可取出token中字段进行相应的鉴权。
* 作为身份提供者（IDP）,可以使用玉符SDK灵活进行参数配置，并根据传入的回调地址生成带有token的跳转url，进行单点登录功能。

## 使用SDK
**Installing**

使用 `go get`将SDK添加到您的`GOPATH` workspace或Go module dependencies 
```go
    go get github.com/zhuermu/sso-go-sdk
```

如果需要更新SDK，请使用`go get -u`检索最新版本的SDK
```go   
    go get -u github.com/zhuermu/sso-go-sdk
```

**服务提供者(SP)**
1. 使用必要信息初始化SDK实例。注：开发阶段可以使用测试公钥`test/testPublicyKey.pem`
```go
        builder := &auth.YufuAuthBuilder{
		YufuAuth: &auth.YufuAuth{},
	}

	spAuth, err := builder.
		TenantId("tn-yufu").
		PublicKeyPath("testPublicKey.pem").
		Issuer("yufuiss").
		SDKRole(auth.SP).
		Build()

```

2. 调用第1步生成的实例`verify`验证单点登录url里的token（自动验证有效期、Issuer、Audience、签名等），如通过，说明该令牌来受信任的有效租户(企业/组织)，样例
```go
        mapClaims, err := auth.VerifyToken(token, spAuth)    //token校验 如果成功会返回包含用户信息的对象，如果失败，错误信息会包含在err中
	tenantId := mapClaims["tnt_id"]   //租户ID
	username := mapClaims["sub"]      //用户名称
	...
```

3. 根据第2步获取的用户信息，服务提供商(SP)在token验证通过后，取出token中用户名称等必要信息，进行相应登录鉴权，否则提示用户登录失败

**身份提供者（IDP)**

1. 使用必要信息初始化SDK（必要参数在玉符初始化后获取）
```go
        var customClaims = make(map[string]interface{})
	customClaims["xxx"] = "xxx" //自定义参数示例

	builder := &auth.YufuAuthBuilder{
		YufuAuth: &auth.YufuAuth{},
	}

	yufuAuth, err := builder.
		Subject("testSub").
		PrivateKeyPath("testPrivateKey.pem").
		Issuer("testIdp").
		SDKRole(auth.IDP).
		Claims(customClaims). //自定义参数，不需要可以不加
		Build()
```
2. 传入必要和自定义参数，生成token/跳转url
```go
 token, err := auth.GenerateToken(&yufuAuth)
 
 url, err := auth.GenerateIDPRedirectUrl(redirectUrl, &yufuAuth)
```
3. 成功生成url后，进行302重定向跳转，若以上参数配置无误，则可单点登录至相应的应用。

详细代码可参考git地址中test文件
