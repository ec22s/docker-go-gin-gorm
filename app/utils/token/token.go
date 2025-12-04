package token

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  "time"

  "github.com/golang-jwt/jwt"
  "github.com/gin-gonic/gin"
)

// 指定されたユーザーIDに基づいてJWTトークンを生成する
func GenerateToken(id uint) (string, error) {
  tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
  if err != nil {
    return "", err
  }
  token := jwt.New(jwt.SigningMethodRS512)
  claims := token.Claims.(jwt.MapClaims)
  claims["authorized"] = true
  claims["user_id"] = id
  claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	privateKeyData, err := os.ReadFile("/tmp/private_key.pem")
  if err != nil {
    return "", err
  }
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
  if err != nil {
    return "", err
  }
  return token.SignedString(key)
}

func extractTokenString(c *gin.Context) string {
  bearToken := c.Request.Header.Get("Authorization")
  strArr := strings.Split(bearToken, " ")
  if len(strArr) == 2 {
    return strArr[1]
  }
  return ""
}

func parseToken(tokenString string) (*jwt.Token, error) {
	publicKeyData, err := os.ReadFile("/tmp/public_key.pem")
  if err != nil {
    return nil, err
  }
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
    key, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
    if err != nil {
      return "", err
    }
    return key, nil
  })
  if err != nil {
    return nil, err
  }
  return token, nil
}

// トークンが有効かどうかを検証
func TokenValid(c *gin.Context) error {
  tokenString := extractTokenString(c)
  token, err := parseToken(tokenString)
  if err != nil {
    return err
  }
  if !token.Valid {
    return fmt.Errorf("Invalid token")
  }
  return nil
}

// トークンからユーザーIDを取得
func ExtractTokenId(c *gin.Context) (uint, error) {
  tokenString := extractTokenString(c)
  token, err := parseToken(tokenString)
  if err != nil {
    return 0, err
  }

  claims, ok := token.Claims.(jwt.MapClaims)
  if ok && token.Valid {
    userId, ok := claims["user_id"].(float64)
    if !ok {
      return 0, nil
    }
    return uint(userId), nil
  }
  return 0, nil
}
