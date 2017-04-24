package exotel

import (
	"encoding/json"
	"fmt"
)

// GetCallRequest : Defines request for getting details of a call.
type GetCallRequest struct {
	SID string
}

// Do : Makes the HTTP request to get the details of a call.
func (cReq *GetCallRequest) Do(e *Exotel) (cRes CallResponse, err error) {
	url := cReq.getURL(e)
	resp, _ := e.doRequest(get, url, nil)
	_ = resp
	defer resp.Body.Close()
	var c callResponse
	err = json.NewDecoder(resp.Body).Decode(&c)
	if err != nil {
		return
	}
	err = cRes.makeResponse(c)
	return
}

func (cReq GetCallRequest) getURL(e *Exotel) (url string) {
	tenantID := e.auth.Username
	url = fmt.Sprintf("%s/Accounts/%s/Calls/%s.json", baseURL, tenantID, cReq.SID)
	return
}
