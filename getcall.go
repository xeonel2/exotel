package exotel

import (
	"encoding/json"
	"fmt"

	"github.com/xeonel2/exotel/exoerror"
)

// GetCallRequest : Defines request for getting details of a call.
type GetCallRequest struct {
	SID string
}

// Do : Makes the HTTP request to get the details of a call.
func (cReq *GetCallRequest) Do(e *Exotel) (cRes CallResponse, err *exoerror.Error) {
	url := cReq.getURL(e)
	resp, err := e.doRequest(get, url, nil)
	defer resp.Body.Close()
	var c callResponse
	er := json.NewDecoder(resp.Body).Decode(&c)
	if er != nil {
		err = exoerror.DecodeError
		return
	}
	er = cRes.makeResponse(c)
	if er != nil {
		err = exoerror.ResponseError
	}
	return
}

func (cReq GetCallRequest) getURL(e *Exotel) (url string) {
	tenantID := e.auth.Username
	url = fmt.Sprintf("%s/Accounts/%s/Calls/%s.json", baseURL, tenantID, cReq.SID)
	return
}
