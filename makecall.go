package exotel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/devpyp/exotel/exoerror"
)

// MakeCallRequest : Defines request object for making a call.
type MakeCallRequest struct {
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
func (cReq *MakeCallRequest) validateCallRequest() *exoerror.Error {
	if cReq.From == "" || cReq.CallerID == "" {
		return exoerror.MissingParams
	}
	if cReq.FlowID == "" && cReq.To == "" {
		return exoerror.MissingParams
	}
	return nil
}

// Do : Actually makes the call using the http client.
func (cReq *MakeCallRequest) Do(e *Exotel) (cRes CallResponse, err *exoerror.Error) {
	err = cReq.validateCallRequest()
	if err != nil {
		return
	}

	params := cReq.makeParams()
	url := cReq.getURL(e)
	body := bytes.NewBufferString(params.Encode())
	resp, err := e.doRequest(post, url, body)
	if err != nil {
		fmt.Println("Error making request" + err.Message)
		return
	}
	defer resp.Body.Close()
	var c callResponse
	er := json.NewDecoder(resp.Body).Decode(&c)
	if er != nil {
		err = exoerror.DecodeError
	}
	er = cRes.makeResponse(c)
	if er != nil {
		err = exoerror.ResponseError
	}
	return
}

// makeCallParams : Generates url.Values for make-call API params.
func (cReq *MakeCallRequest) makeParams() (data url.Values) {
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
func (MakeCallRequest) getURL(e *Exotel) (url string) {
	tenantID := e.auth.Username
	url = fmt.Sprintf("%s/Accounts/%s/Calls/connect.json", baseURL, tenantID)
	return
}
