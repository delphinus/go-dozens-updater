package dozens

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/delphinus/godo/lib/dozens/endpoint"
	"github.com/pkg/errors"
)

// ZoneCreateBody means post data for `create` request
type ZoneCreateBody struct {
	Name            string `json:"name"`
	AddGoogleApps   bool   `json:"add_google_apps"`
	GoogleAuthorize string `json:"google_authorize,omitempty"`
	MailAddress     string `json:"mailaddress,omitempty"`
}

// ZoneCreate creates zone and returns zones list
func ZoneCreate(body ZoneCreateBody) (ZoneResponse, error) {
	zoneResp := ZoneResponse{}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return zoneResp, errors.Wrap(err, "error in Marshal")
	}

	req, err := MakePost(endpoint.Create(), bytes.NewBuffer(bodyJSON))
	if err != nil {
		return zoneResp, errors.Wrap(err, "error in MakeGet")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return zoneResp, errors.Wrap(err, "error in Do")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return zoneResp, errors.Wrap(err, "error in ReadAll")
		}
		return zoneResp, errors.Errorf("error body: %s", body)
	}

	if err := json.NewDecoder(resp.Body).Decode(&zoneResp); err != nil {
		return zoneResp, errors.Wrap(err, "error in Decode")
	}

	return zoneResp, nil
}
