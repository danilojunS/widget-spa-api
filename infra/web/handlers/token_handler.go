package handlers

import (
	config "github.com/danilojunS/widgets-spa-api/config"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

// TokenGet handler
// this is a helper for creating tokens (for testing)
// in a production app, the tokens would be generated elsewhere
// or with some credentials logic
func TokenGet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	mySigningKey := []byte(config.Get().TokenSecret)

	oneDayInSeconds := int64(60 * 60 * 24)

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + oneDayInSeconds,
		Issuer:    "secret-issuer",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		InternalError(w, "")
		return
	}

	_, err = w.Write([]byte(ss))
	if err != nil {
		InternalError(w, "")
		return
	}
}
