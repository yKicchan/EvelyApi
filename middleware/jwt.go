package middleware

import (
	"EvelyApi/app"
	"context"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

/**
 * JWT認証用のミドルウェアを作成する
 */
func NewJWTMiddleware() goa.Middleware {
	keys, err := LoadJWTPublicKeys()
	if err != nil {
		log.Fatalf("failed to load file: %s", err)
	}
	return jwt.New(jwt.NewSimpleResolver(keys), validationHandler, app.NewJWTSecurity())
}

/**
 * 公開鍵を読み込む
 */
func LoadJWTPublicKeys() ([]jwt.Key, error) {
	keyFiles, err := filepath.Glob("./keys/*.pub")
	if err != nil {
		return nil, err
	}
	keys := make([]jwt.Key, len(keyFiles))
	for i, keyFile := range keyFiles {
		pem, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return nil, err
		}
		key, err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, fmt.Errorf("failed to load key %s: %s", keyFile, err)
		}
		keys[i] = key
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("couldn't load public keys for JWT security")
	}

	return keys, nil
}

/**
 * JWTをチェックする
 */
var validationHandler, _ = goa.NewMiddleware(func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// token取得
	token := jwt.ContextJWT(ctx)
	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !ok {
		return jwt.ErrJWTError("unsupported claims shape")
	}
	// claimsを検査
	if val, ok := claims["id"].(string); !ok || val != "" {
		return jwt.ErrJWTError("you are not user")
	}
	return nil
})
