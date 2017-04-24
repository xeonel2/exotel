package exotel

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// CallRequest : Defines request object for a call.
type CallRequest struct {
	From           string
	To             string
	CallerID       string
	CallType       string
	TimeLimit      int
	TimeOut        int
	StatusCallBack string
	FlowID         string
	CustomField    string
}

// validateCallRequest : Validates parameters for the make-call API.
func (cReq *CallRequest) validateCallRequest() error {
	if cReq.From == "" || cReq.CallerID == "" {
		return errors.New("Mandatory parameters missing")
	}
	if cReq.FlowID == "" && cReq.To == "" {
		return errors.New("Mandaotyr paramters missing")
	}
	return nil
}

// DoCall : Actually makes the call using the http client.
func (cReq *CallRequest) DoCall(e *Exotel) (cRes CallResponse, err error) {
	err = cReq.validateCallRequest()
	if err != nil {
		return
	}

	params := cReq.makeCallParams()
	url := cReq.makeCallURL(e)
	resp, _ := e.doRequest(url, params)
	defer resp.Body.Close()
	var c callResponse
	err = json.NewDecoder(resp.Body).Decode(&c)
	if err != nil {
		return
	}
	err = cRes.makeResponse(c)
	return
}

// makeCallParams : Generates url.Values for make-call API params.
func (cReq *CallRequest) makeCallParams() (data url.Values) {
	data = url.Values{
		"From":           {cReq.From},
		"CallerId":       {cReq.CallerID},
		"CallType":       {"trans"},
		"TimeLimit":      {strconv.Itoa(cReq.TimeLimit)},
		"TimeOut":        {strconv.Itoa(cReq.TimeOut)},
		"StatusCallBack": {cReq.StatusCallBack},
	}
	if cReq.To != "" {
		data["To"] = []string{cReq.To}
	} else {
		data["Url"] = []string{"http://my.exotel.in/exoml/start/" + cReq.FlowID}
		data["CustomField"] = []string{cReq.CustomField}
	}
	return
}

// makeCallURL : Makes a URL for the make-call API.
func (CallRequest) makeCallURL(e *Exotel) (url string) {
	tenantID := e.auth.Username
	url = fmt.Sprintf("%s/Accounts/%s/Calls/connect.json", baseURL, tenantID)
	return
}
