/*
 * Twilio - Studio
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

// Retrieve a Step.
func (c *ApiService) FetchStep(FlowSid string, EngagementSid string, Sid string) (*StudioV1FlowEngagementStep, error) {
	path := "/v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"EngagementSid"+"}", EngagementSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1FlowEngagementStep{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListStep'
type ListStepParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListStepParams) SetPageSize(PageSize int) *ListStepParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of Step records from the API. Request is executed immediately.
func (c *ApiService) PageStep(FlowSid string, EngagementSid string, params *ListStepParams, pageToken string, pageNumber string) (*ListStepResponse, error) {
	path := "/v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps"

	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"EngagementSid"+"}", EngagementSid, -1)

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

	ps := &ListStepResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Step records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListStep(FlowSid string, EngagementSid string, params *ListStepParams, limit int) ([]StudioV1FlowEngagementStep, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageStep(FlowSid, EngagementSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []StudioV1FlowEngagementStep

	for response != nil {
		records = append(records, response.Steps...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListStepResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListStepResponse)
	}

	return records, err
}

// Streams Step records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamStep(FlowSid string, EngagementSid string, params *ListStepParams, limit int) (chan StudioV1FlowEngagementStep, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageStep(FlowSid, EngagementSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan StudioV1FlowEngagementStep, 1)

	go func() {
		for response != nil {
			for item := range response.Steps {
				channel <- response.Steps[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListStepResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListStepResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListStepResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListStepResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}