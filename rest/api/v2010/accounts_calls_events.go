/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
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

// Optional parameters for the method 'ListCallEvent'
type ListCallEventParams struct {
	// The unique SID identifier of the Account.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListCallEventParams) SetPathAccountSid(PathAccountSid string) *ListCallEventParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListCallEventParams) SetPageSize(PageSize int) *ListCallEventParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of CallEvent records from the API. Request is executed immediately.
func (c *ApiService) PageCallEvent(CallSid string, params *ListCallEventParams, pageToken string, pageNumber string) (*ListCallEventResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Events.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCallEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CallEvent records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCallEvent(CallSid string, params *ListCallEventParams, limit int) ([]ApiV2010AccountCallCallEvent, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageCallEvent(CallSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010AccountCallCallEvent

	for response != nil {
		records = append(records, response.Events...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListCallEventResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListCallEventResponse)
	}

	return records, err
}

// Streams CallEvent records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCallEvent(CallSid string, params *ListCallEventParams, limit int) (chan ApiV2010AccountCallCallEvent, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageCallEvent(CallSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010AccountCallCallEvent, 1)

	go func() {
		for response != nil {
			for item := range response.Events {
				channel <- response.Events[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListCallEventResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListCallEventResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListCallEventResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCallEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}