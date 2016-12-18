package dozens

import (
	"encoding/json"
	"net/http"

	"github.com/delphinus/godo/lib/dozens/endpoint"
	"github.com/pkg/errors"
)

// ZonesResponse means response of zones
type ZonesResponse struct {
	Domain []domain `json:"domain"`
}

type domain struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetZones returns zones list
func GetZones() (ZonesResponse, error) {
	zonesResp := ZonesResponse{}

	req, err := MakeGet(endpoint.Zone())
	if err != nil {
		return zonesResp, errors.Wrap(err, "error in MakeGet")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return zonesResp, errors.Wrap(err, "error in Do")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return zonesResp, errors.Errorf("error status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&zonesResp); err != nil {
		return zonesResp, errors.Wrap(err, "error in Decode")
	}

	return zonesResp, nil
}
