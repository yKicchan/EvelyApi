package middleware

import (
	"EvelyApi/app"
	. "EvelyApi/models"
	. "EvelyApi/models/collections"
	"context"
	"errors"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

/**
 * JWT認証用のミドルウェアを作成する
 */
func NewJWTMiddleware(db *EvelyDB) (goa.Middleware, error) {
	keys, err := LoadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	return jwt.New(jwt.NewSimpleResolver(keys), ForceFail(db), app.NewJWTSecurity()), nil
}

/**
 * JWT認証用のミドルウェアを生成する
 * 認証有無の両方でアクセスできるエンドポイント専用で使う
 */
func NewOptionalJWTMiddleware(db *EvelyDB) (goa.Middleware, error) {
	keys, err := LoadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	jwtMiddleware := jwt.New(jwt.NewSimpleResolver(keys), ForceFail(db), app.NewJWTSecurity())
	return func(nextHandler goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			err := jwtMiddleware(nextHandler)(ctx, rw, req)
			if err != nil {
				// err が起きた場合 nextHandlerは呼ばれないが、今回は処理を継続する
				return nextHandler(ctx, rw, req)
			}
			return nil
		}
	}, nil

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

// ForceFail is a middleware illustrating the use of validation middleware with JWT auth.  It checks
// for the presence of a "fail" query string and fails validation if set to the value "true".
func ForceFail(db *EvelyDB) goa.Middleware {
	errValidationFailed := goa.NewErrorClass("validation_failed", 401)
	forceFail := func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			id, err := GetLoginID(ctx)
			if err != nil || !db.Users.Exists(Keys{"id": id}) {
				return errValidationFailed("forcing failure to illustrate Validation middleware")
			}
			return h(ctx, rw, req)
		}
	}
	fm, _ := goa.NewMiddleware(forceFail)
	return fm
}

/**
 * 受け取った任意の情報を付加してトークンを生成する
 * @param  claims トークンに埋め込む情報
 * @return string 生成したトークン
 */
func NewToken(claims jwtgo.MapClaims) (string, error) {
	// jwt生成
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	token.Claims = claims
	// 秘密鍵読み込み
	pem, err := ioutil.ReadFile("./keys/id_rsa")
	if err != nil {
		return "", fmt.Errorf("Faild to load file: %s", err)
	}
	privateKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		return "", fmt.Errorf("Faild to parse private key: %s", err)
	}
	// jwtを暗号化
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("Failed to sign token: %s", err)
	}
	return signedToken, nil
}

/**
 * JWTを複合してログイン中のユーザーIDを返す
 * @param  ctx コンテキスト
 * @return id  jwtに埋め込まれたユーザーID
 * @return err 複合中に発生したエラー
 */
func GetLoginID(ctx context.Context) (string, error) {
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return "", errors.New("token is nothing.")
	}
	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !ok {
		return "", errors.New("Unsupported claims shape")
	}
	id, ok := claims["id"].(string)
	if !ok {
		return "", errors.New("decode error")
	}
	return id, nil
}
