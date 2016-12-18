package dozens

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/delphinus/godo/lib/dozens/endpoint"
	"github.com/pkg/errors"
)

// ZoneDelete creates zone and returns zones list
func ZoneDelete(token, zoneID string) (ZoneResponse, error) {
	zoneResp := ZoneResponse{}

	req, err := MakeDelete(token, endpoint.ZoneDelete(zoneID))
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
