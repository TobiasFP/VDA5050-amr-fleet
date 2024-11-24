package auth

import (
	"TobiasFP/BotNana/config"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var whitelistedAccountKeys = []string{"D0Fok1CkYG6uxd6Spgl84Q==", "sHCnjdPl9cU|2qwD1Xvzrw=="}

// Auth is A simple struct to handle authentication configuration
type Auth struct {
	Config oauth2.Config
}

// Login is the login function, that redirects the user to keycloak/OAuth
func (auth Auth) Login(ctx *gin.Context) {
	state, err := randString(16)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.SetCookie("state", state, int(time.Hour.Seconds()), "", "", false, true)
	ctx.Redirect(http.StatusTemporaryRedirect, auth.Config.AuthCodeURL(state))
}

// Callback handles the callback from keycloak/OIDC
func (auth *Auth) Callback(ctx *gin.Context) {
	conf := config.GetConfig()

	state, err := ctx.Cookie("state")
	if err != nil {
		ctx.Error(err)
		return
	}
	if ctx.Query("state") != state {
		ctx.Error(errors.New("state did not match"))
		return
	}

	authToken, err := auth.Config.Exchange(ctx, ctx.Query("code"))

	if err != nil {
		ctx.Error(errors.New("Failed to exchange token: " + err.Error()))
		return
	}

	ctx.Redirect(
		http.StatusPermanentRedirect,
		conf.GetString("appUrl")+"/"+
			"callback/"+
			authToken.AccessToken+"/"+
			authToken.RefreshToken+"/"+
			strconv.FormatInt(authToken.Expiry.Unix(), 10),
	)
}

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func filter(haystack []string, key string, intermediary []string) []string {
	if len(haystack) == 0 {
		return intermediary
	}

	if haystack[0] == key {
		appendedIntermediary := append(intermediary, key)
		return filter(haystack[1:], key, appendedIntermediary)
	}
	return filter(haystack[1:], key, intermediary)
}
