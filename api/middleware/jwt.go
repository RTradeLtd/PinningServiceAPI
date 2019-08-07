package middleware

import (
	"errors"
	"fmt"
	"time"

	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

// Verifier is used to handle verification of JWT's
type Verifier struct {
	key []byte
}

// NewJWT is used to generate a new JWT token
func (v *Verifier) NewJWT(sub string) (string, error) {
	// generate a jwt with claims to verify email
	verificationJWT := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":    sub,
		"expire": time.Now().Add(time.Hour * 24).UTC().String(),
	})
	// return a signed version of the jwt
	return verificationJWT.SignedString(v.key)
}

// Verify is used to verify a jwt against a username (sub)
func (v *Verifier) Verify(jwtString, username string) error {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unable to validate signing method: %v", token.Header["alg"])
		} else if method != jwt.SigningMethodHS512 {
			return nil, errors.New("expect hs512 signing method")
		}
		// return byte version of signing key
		return []byte(v.key), nil
	})
	// verify jwt was parsed properly
	if err != nil {
		return err
	}
	// verify that the token is valid
	if !token.Valid {
		return errors.New("failed to validate token")
	}
	// extract claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("failed to parse claims")
	}
	// verify the sub parameter of the claim matches the expected one
	if claims["sub"] != username {
		return fmt.Errorf("username from claim does not match expected user of %s", username)
	}
	return nil
}
