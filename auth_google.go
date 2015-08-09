package devastator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nbusy/neptulon/jsonrpc"
)

// Response from GET https://www.googleapis.com/plus/v1/people/me?access_token=...
// has the following structure with denoted fields of interest (rest is left out):
type gProfile struct {
	Emails      []gEmail
	DisplayName string
	Image       gImage
}

type gEmail struct {
	Value string
}

type gImage struct {
	URL string
}

// googleAuth authenticates a user with Google+ using provided OAuth 2.0 access token.
// If authenticated successfully, user profile is retrieved from Google+ and user is given a TLS client-certificate in return.
func googleAuth(ctx *jsonrpc.ReqContext, db DB, cm *CertMgr) {
	t := ctx.Req.Params.(map[string]interface{})["accessToken"]
	p, i, err := getGProfile(t.(string))
	if err != nil {
		ctx.ResErr = &jsonrpc.ResError{Code: 666, Message: "Failed to authenticated user with Google+ OAuth access token."}
		log.Printf("Errored during Google+ profile call using provided access token: %v with error: %v", t, err)
	}

	var key []byte
	if key == nil {
	}

	// user is authenticated at this point so check if this is a first-time registration
	if user, ok := db.GetByMail(p.Emails[0].Value); ok {
		if user.Cert == nil {
			// todo: add CertMgr
			if user.Cert, key, err = cm.GenClientCert(string(user.ID)); err != nil {
				log.Fatal("Failed to generate client certificate for user:", err)
			}
			db.SaveUser(user)
		}

		ctx.Res = user.Cert
		ctx.Conn.Session.Set("userid", user.ID)
		return
	}

	// first-time login so generate create user
	u := User{Email: p.Emails[0].Value, Name: p.DisplayName, Picture: i, Cert: make([]byte, 555)}
	db.SaveUser(&u)
	ctx.Res = u.Cert
	ctx.Conn.Session.Set("userid", u.ID)
	return
}

// getGProfile retrieves user info (display name, e-mail, profile pic) using an access token that has 'profile' and 'email' scopes.
// Also retrieves user profile image via profile image URL provided the response.
func getGProfile(token string) (profile *gProfile, profilePic []byte, err error) {
	// retrieve profile info from Google
	uri := fmt.Sprintf("https://www.googleapis.com/plus/v1/people/me?access_token=%s", token)
	res, err := http.Get(uri)
	if err != nil {
		return
	}

	resBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return
	}

	if err = json.Unmarshal(resBody, profile); err != nil {
		return
	}

	// retrieve profile image
	uri = profile.Image.URL
	res, err = http.Get(uri)
	if err != nil {
		return
	}

	profilePic, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return
	}

	return
}
