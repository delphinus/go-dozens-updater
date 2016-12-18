package dozens

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/delphinus/godo/lib/dozens/endpoint"
	"github.com/pkg/errors"
)

// ZoneUpdateBody means post data for `update` request
type ZoneUpdateBody struct {
	MailAddress string `json:"mailaddress"`
}

// ZoneUpdate creates zone and returns zones list
func ZoneUpdate(token, zoneID, mailAddress string) (ZoneResponse, error) {
	zoneResp := ZoneResponse{}
	body := ZoneUpdateBody{mailAddress}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return zoneResp, errors.Wrap(err, "error in Marshal")
	}

	req, err := MakePost(token, endpoint.ZoneUpdate(zoneID), bytes.NewBuffer(bodyJSON))
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
