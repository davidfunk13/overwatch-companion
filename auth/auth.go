package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

// Jwks Keys struct
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// JSONWebKeys struct
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// JWTMiddleware handler
var JWTMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {

		// Verify 'aud' claim
		aud := os.Getenv("AUTH0_AUDIENCE")

		//isolate array of available audiences
		audienceArray := token.Claims.(jwt.MapClaims)["aud"].([]interface{})

		// create a map set to lookup if audience exists.
		set := make(map[string]bool)

		// create set
		for _, a := range audienceArray {
			set[a.(string)] = true
		}

		//boolean value representing whether the set contains the audience of this api
		contains := set[aud]

		//return err if false
		if !contains {
			return token, errors.New("invalid audience")
		}

		// Verify 'iss' claim
		iss := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"

		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)

		if !checkIss {
			return token, errors.New("invalid issuer")
		}

		cert, err := getPemCert(token)

		if err != nil {
			panic(err.Error())
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))

		return result, nil
	},
	SigningMethod: jwt.SigningMethodRS256,
})

// CustomClaims struct defining user scope
type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

// CheckScope checks user scope and permissions
func CheckScope(scope string, tokenString string) bool {
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		cert, err := getPemCert(token)
		if err != nil {
			return nil, err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	})

	claims, ok := token.Claims.(*CustomClaims)

	hasScope := false
	if ok && token.Valid {
		result := strings.Split(claims.Scope, " ")
		for i := range result {
			if result[i] == scope {
				hasScope = true
			}
		}
	}

	return hasScope
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://" + os.Getenv("AUTH0_DOMAIN") + "/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}

	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}
