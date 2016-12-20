package dozens

import (
	"bytes"
	"encoding/json"

	"github.com/delphinus/godo/lib/dozens/endpoint"
	"github.com/pkg/errors"
)

// RecordList returns records list
func RecordList(token, zoneName string) (RecordResponse, error) {
	recordResp := RecordResponse{}

	req, err := MakeGet(token, endpoint.RecordList(zoneName))
	if err != nil {
		return recordResp, errors.Wrap(err, "error in MakeGet")
	}

	return doRecordRequest(req)
}

// RecordDelete deletes record and returns records list
func RecordDelete(token, recordID string) (RecordResponse, error) {
	req, err := MakeDelete(token, endpoint.RecordDelete(recordID))
	if err != nil {
		return RecordResponse{}, errors.Wrap(err, "error in MakeGet")
	}

	return doRecordRequest(req)
}

// RecordUpdateBody means post data for `update` request
type RecordUpdateBody struct {
	Prio    uint   `json:"prio,omitempty"`
	Content string `json:"content,omitempty"`
	TTL     string `json:"ttl,omitempty"`
}

// RecordUpdate updates record and returns records list
func RecordUpdate(token string, recordID string, body RecordUpdateBody) (RecordResponse, error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return RecordResponse{}, errors.Wrap(err, "error in Marshal")
	}

	req, err := MakePost(token, endpoint.RecordUpdate(recordID), bytes.NewBuffer(bodyJSON))
	if err != nil {
		return RecordResponse{}, errors.Wrap(err, "error in MakeUpdate")
	}

	return doRecordRequest(req)
}
