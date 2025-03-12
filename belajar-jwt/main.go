package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("makan bang")

func generateToken(username string) (string, error) {

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
		Issuer:    "test",
		Subject:   username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		return " ", err
	}

	return ss, nil

}

func validatedToken(tokenString string) (*jwt.RegisteredClaims, error) {

	tokenValidation, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("inexpected signing method %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Konversi claims menjadi `*jwt.RegisteredClaims`
	if claims, ok := tokenValidation.Claims.(*jwt.RegisteredClaims); ok && tokenValidation.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func main() {

	fmt.Println("Proses generate token ")

	token, err := generateToken("admin")
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated token:", token)

	fmt.Println("Proses validasi token ")
	validasiToken, err := validatedToken(token)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Token is valid")
	fmt.Println("Expires At:", validasiToken.ExpiresAt)
	fmt.Println("Issuer:", validasiToken.Issuer)
	fmt.Println("Subject (Username):", validasiToken.Subject)

}
