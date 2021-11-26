package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

// Claims defines jwt claims
type Claims struct {
    UserID string `json:"email"`
    jwt.StandardClaims
}

// GenerateToken handles generation of a jwt code
// @returns string -> token and error -> err
func GenerateToken(userID string) (string, error) {
    // Define token expiration time
    expirationTime := time.Now().Add(1440 * time.Minute)
    // Define the payload and exp time
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    // Generate token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign token with secret key encoding
    tokenString, err := token.SignedString(jwtKey)

    return tokenString, err
}

// DecodeToken handles decoding a jwt token
func DecodeToken(tkStr string) (string, error) {
    claims := &Claims{}

    tkn, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            return "", err
        }
        return "", err
    }

    if !tkn.Valid {
        return "", err
    }

    return claims.UserID, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, isvalid := token.Method.(*jwt.SigningMethodHMAC)
		if !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		fmt.Println(isvalid)
		return []byte(jwtKey), nil
	})
}