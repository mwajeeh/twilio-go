/*
 * Twilio - Studio
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

// Delete a specific Flow.
func (c *ApiService) DeleteFlow(Sid string) error {
	path := "/v1/Flows/{Sid}"
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

// Retrieve a specific Flow.
func (c *ApiService) FetchFlow(Sid string) (*StudioV1Flow, error) {
	path := "/v1/Flows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1Flow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFlow'
type ListFlowParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFlowParams) SetPageSize(PageSize int) *ListFlowParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFlowParams) SetLimit(Limit int) *ListFlowParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Flow records from the API. Request is executed immediately.
func (c *ApiService) PageFlow(params *ListFlowParams, pageToken, pageNumber string) (*ListFlowResponse, error) {
	path := "/v1/Flows"

	data := url.Values{}
	headers := make(map[string]interface{})

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

	ps := &ListFlowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Flow records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFlow(params *ListFlowParams) ([]StudioV1Flow, error) {
	if params == nil {
		params = &ListFlowParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageFlow(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []StudioV1Flow

	for response != nil {
		records = append(records, response.Flows...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListFlowResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListFlowResponse)
	}

	return records, err
}

// Streams Flow records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFlow(params *ListFlowParams) (chan StudioV1Flow, error) {
	if params == nil {
		params = &ListFlowParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageFlow(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan StudioV1Flow, 1)

	go func() {
		for response != nil {
			for item := range response.Flows {
				channel <- response.Flows[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListFlowResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListFlowResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListFlowResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFlowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
