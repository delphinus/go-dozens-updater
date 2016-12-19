package dozens

import (
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
