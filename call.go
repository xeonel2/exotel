package exotel

import (
	"encoding/json"
	"strconv"
	"time"
)

// callResponse : Defines response struct to parse JSON response.
type callResponse struct {
	Call struct {
		SID            string `json:"Sid"`
		From           string `json:"From"`
		ParentCallSID  string `json:"ParentCallSid"`
		DateCreated    string `json:"DateCreated"`
		DateUpdated    string `json:"DateUpdated"`
		AccountSid     string `json:"AccountSid"`
		To             string `json:"To"`
		PhoneNumberSid string `json:"PhoneNumberSid"`
		Status         string `json:"Status"`
		StartTime      string `json:"StartTime"`
		EndTime        string `json:"EndTime"`
		Duration       string `json:"Duration"`
		Price          string `json:"Price"`
		Direction      string `json:"Direction"`
		AnsweredBy     string `json:"AnsweredBy"`
		ForwardedFrom  string `json:"ForwardedFrom"`
		CallerName     string `json:"CallerName"`
		URI            string `json:"Uri"`
		RecordingURL   string `json:"RecordingUrl"`
	} `json:"Call"`
}

// CallResponse : Defines response struct.
type CallResponse struct {
	Call struct {
		SID            string
		From           string
		ParentCallSID  string
		DateCreated    time.Time
		DateUpdated    time.Time
		AccountSid     string
		To             string
		PhoneNumberSid string
		Status         string
		StartTime      time.Time
		EndTime        time.Time
		Duration       time.Duration
		Price          float64
		Direction      string
		AnsweredBy     string
		ForwardedFrom  string
		CallerName     string
		URI            string
		RecordingURL   string
	} `json:"Call"`
}

// makeResponse : Creates CallResponse struct from callResponse
func (cRes *CallResponse) makeResponse(c callResponse) (err error) {
	cRes.Call.SID = c.Call.SID
	cRes.Call.From = c.Call.From
	cRes.Call.ParentCallSID = c.Call.ParentCallSID
	if c.Call.DateCreated != "" {
		cRes.Call.DateCreated, err = time.Parse(timeFormat, c.Call.DateCreated+" IST")
	}
	if c.Call.DateUpdated != "" {
		cRes.Call.DateUpdated, err = time.Parse(timeFormat, c.Call.DateUpdated+" IST")
	}
	cRes.Call.AccountSid = c.Call.AccountSid
	cRes.Call.To = c.Call.To
	cRes.Call.PhoneNumberSid = c.Call.PhoneNumberSid
	cRes.Call.Status = c.Call.Status
	if c.Call.StartTime != "" {
		cRes.Call.StartTime, err = time.Parse(timeFormat, c.Call.StartTime+" IST")
	}
	if c.Call.EndTime != "" {
		cRes.Call.EndTime, err = time.Parse(timeFormat, c.Call.EndTime+" IST")
	}

	var duration int
	if c.Call.Duration != "" {
		duration, err = strconv.Atoi(c.Call.Duration)
		cRes.Call.Duration = time.Duration(duration) * time.Second
	}

	if c.Call.Price != "" {
		cRes.Call.Price, err = strconv.ParseFloat(c.Call.Price, 64)
	}

	cRes.Call.Direction = c.Call.Direction
	cRes.Call.AnsweredBy = c.Call.AnsweredBy
	cRes.Call.ForwardedFrom = c.Call.ForwardedFrom
	cRes.Call.CallerName = c.Call.CallerName
	cRes.Call.URI = c.Call.URI
	cRes.Call.RecordingURL = c.Call.RecordingURL
	return
}

func parseResponseBody(b []byte) (call map[string]interface{}, err error) {
	var v struct {
		Call map[string]interface{} `json:"Call"`
	}
	json.Unmarshal(b, &v)
	call = v.Call
	return
}
