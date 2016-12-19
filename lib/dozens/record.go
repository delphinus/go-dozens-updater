package dozens

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// RecordResponse means response of zones
type RecordResponse struct {
	Record []record `json:"record"`
}

type record struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Prio    string `json:"prio"`
	Content string `json:"content"`
	TTL     string `json:"ttl"`
}

func doRecordRequest(req *http.Request) (RecordResponse, error) {
	recordResp := RecordResponse{}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return recordResp, errors.Wrap(err, "error in Do")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return recordResp, errors.Wrap(err, "error in ReadAll")
		}
		return recordResp, errors.Errorf("error body: %s", body)
	}

	if err := json.NewDecoder(resp.Body).Decode(&recordResp); err != nil {
		return recordResp, errors.Wrap(err, "error in Decode")
	}

	return recordResp, nil
}
