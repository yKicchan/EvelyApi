package middleware

import (
	"EvelyApi/app"
	. "EvelyApi/config"
	. "EvelyApi/models"
	. "EvelyApi/models/collections"
	"context"
	"errors"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"io/ioutil"
	"labix.org/v2/mgo"
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
	// トークンに埋め込まれたユーザーを検査
	id, err := GetLoginID(ctx)
	if err != nil {
		return jwt.ErrJWTError(err)
	}
	session, err := mgo.Dial(DB_HOST)
	if err != nil {
		return jwt.ErrJWTError(fmt.Sprintf("Database initialization failed: %s", err))
	}
	defer session.Close()
	db := NewEvelyDB(session.DB(DB_NAME))
	if !db.Users.Exists(Keys{"id": id}) {
		return jwt.ErrJWTError("You are not registed user")
	}
	return nil
})

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
