/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'DeleteConferenceRecording'
type DeleteConferenceRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteConferenceRecordingParams) SetPathAccountSid(PathAccountSid string) *DeleteConferenceRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a recording from your account
func (c *ApiService) DeleteConferenceRecording(ConferenceSid string, Sid string, params *DeleteConferenceRecordingParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchConferenceRecording'
type FetchConferenceRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchConferenceRecordingParams) SetPathAccountSid(PathAccountSid string) *FetchConferenceRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a recording for a call
func (c *ApiService) FetchConferenceRecording(ConferenceSid string, Sid string, params *FetchConferenceRecordingParams) (*ApiV2010ConferenceRecording, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010ConferenceRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConferenceRecording'
type ListConferenceRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date.
	DateCreated *string `json:"DateCreated,omitempty"`
	// The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date.
	DateCreatedBefore *string `json:"DateCreated&lt;,omitempty"`
	// The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date.
	DateCreatedAfter *string `json:"DateCreated&gt;,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConferenceRecordingParams) SetPathAccountSid(PathAccountSid string) *ListConferenceRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListConferenceRecordingParams) SetDateCreated(DateCreated string) *ListConferenceRecordingParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *ListConferenceRecordingParams) SetDateCreatedBefore(DateCreatedBefore string) *ListConferenceRecordingParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListConferenceRecordingParams) SetDateCreatedAfter(DateCreatedAfter string) *ListConferenceRecordingParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListConferenceRecordingParams) SetPageSize(PageSize int) *ListConferenceRecordingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConferenceRecordingParams) SetLimit(Limit int) *ListConferenceRecordingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConferenceRecording records from the API. Request is executed immediately.
func (c *ApiService) PageConferenceRecording(ConferenceSid string, params *ListConferenceRecordingParams, pageToken, pageNumber string) (*ListConferenceRecordingResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint(*params.DateCreated))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreated<", fmt.Sprint(*params.DateCreatedBefore))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreated>", fmt.Sprint(*params.DateCreatedAfter))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConferenceRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConferenceRecording records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConferenceRecording(ConferenceSid string, params *ListConferenceRecordingParams) ([]ApiV2010ConferenceRecording, error) {
	if params == nil {
		params = &ListConferenceRecordingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConferenceRecording(ConferenceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010ConferenceRecording

	for response != nil {
		records = append(records, response.Recordings...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConferenceRecordingResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListConferenceRecordingResponse)
	}

	return records, err
}

// Streams ConferenceRecording records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConferenceRecording(ConferenceSid string, params *ListConferenceRecordingParams) (chan ApiV2010ConferenceRecording, error) {
	if params == nil {
		params = &ListConferenceRecordingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConferenceRecording(ConferenceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010ConferenceRecording, 1)

	go func() {
		for response != nil {
			for item := range response.Recordings {
				channel <- response.Recordings[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConferenceRecordingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListConferenceRecordingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListConferenceRecordingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConferenceRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateConferenceRecording'
type UpdateConferenceRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Whether to record during a pause. Can be: `skip` or `silence` and the default is `silence`. `skip` does not record during the pause period, while `silence` will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting `status` is set to `paused`.
	PauseBehavior *string `json:"PauseBehavior,omitempty"`
	// The new status of the recording. Can be: `stopped`, `paused`, `in-progress`.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateConferenceRecordingParams) SetPathAccountSid(PathAccountSid string) *UpdateConferenceRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateConferenceRecordingParams) SetPauseBehavior(PauseBehavior string) *UpdateConferenceRecordingParams {
	params.PauseBehavior = &PauseBehavior
	return params
}
func (params *UpdateConferenceRecordingParams) SetStatus(Status string) *UpdateConferenceRecordingParams {
	params.Status = &Status
	return params
}

// Changes the status of the recording to paused, stopped, or in-progress. Note: To use &#x60;Twilio.CURRENT&#x60;, pass it as recording sid.
func (c *ApiService) UpdateConferenceRecording(ConferenceSid string, Sid string, params *UpdateConferenceRecordingParams) (*ApiV2010ConferenceRecording, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConferenceSid"+"}", ConferenceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PauseBehavior != nil {
		data.Set("PauseBehavior", *params.PauseBehavior)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010ConferenceRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
