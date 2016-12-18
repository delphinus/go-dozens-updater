package dozens

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/delphinus/godo/lib/dozens/endpoint"
	"github.com/pkg/errors"
)

// AuthorizeResponse means response of authorize
type AuthorizeResponse struct {
	AuthToken string `json:"auth_token"`
}

// GetAuthorize returns authorization info
func GetAuthorize() (AuthorizeResponse, error) {
	authorizeResp := AuthorizeResponse{}

	req, err := http.NewRequest("GET", endpoint.Authorize().String(), nil)
	if err != nil {
		return authorizeResp, errors.Wrap(err, "error in NewRequest")
	}
	req.Header.Set("X-Auth-Key", Config.Key)
	req.Header.Set("X-Auth-User", Config.User)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return authorizeResp, errors.Wrap(err, "error in Do")
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return authorizeResp, errors.Wrap(err, "error in ReadAll")
	}

	if err := json.Unmarshal(result, &authorizeResp); err != nil {
		return authorizeResp, errors.Wrap(err, "error in Unmarshal")
	}

	return authorizeResp, nil
}
