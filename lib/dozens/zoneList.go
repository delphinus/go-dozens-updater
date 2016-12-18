package dozens

import (
	"encoding/json"
	"net/http"

	"github.com/delphinus/godo/lib/dozens/endpoint"
	"github.com/pkg/errors"
)

// ZoneResponse means response of zones
type ZoneResponse struct {
	Domain []domain `json:"domain"`
}

type domain struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ZoneList returns zones list
func ZoneList() (ZoneResponse, error) {
	zoneResp := ZoneResponse{}

	req, err := MakeGet(endpoint.ZoneList())
	if err != nil {
		return zoneResp, errors.Wrap(err, "error in MakeGet")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return zoneResp, errors.Wrap(err, "error in Do")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return zoneResp, errors.Errorf("error status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&zoneResp); err != nil {
		return zoneResp, errors.Wrap(err, "error in Decode")
	}

	return zoneResp, nil
}
