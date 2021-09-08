/*
 * Twilio - Numbers
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

// Fetch specific Regulation Instance.
func (c *ApiService) FetchRegulation(Sid string) (*NumbersV2Regulation, error) {
	path := "/v2/RegulatoryCompliance/Regulations/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2Regulation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRegulation'
type ListRegulationParams struct {
	// The type of End User the regulation requires - can be `individual` or `business`.
	EndUserType *string `json:"EndUserType,omitempty"`
	// The ISO country code of the phone number's country.
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The type of phone number that the regulatory requiremnt is restricting.
	NumberType *string `json:"NumberType,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRegulationParams) SetEndUserType(EndUserType string) *ListRegulationParams {
	params.EndUserType = &EndUserType
	return params
}
func (params *ListRegulationParams) SetIsoCountry(IsoCountry string) *ListRegulationParams {
	params.IsoCountry = &IsoCountry
	return params
}
func (params *ListRegulationParams) SetNumberType(NumberType string) *ListRegulationParams {
	params.NumberType = &NumberType
	return params
}
func (params *ListRegulationParams) SetPageSize(PageSize int) *ListRegulationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRegulationParams) SetLimit(Limit int) *ListRegulationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Regulation records from the API. Request is executed immediately.
func (c *ApiService) PageRegulation(params *ListRegulationParams, pageToken, pageNumber string) (*ListRegulationResponse, error) {
	path := "/v2/RegulatoryCompliance/Regulations"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.EndUserType != nil {
		data.Set("EndUserType", *params.EndUserType)
	}
	if params != nil && params.IsoCountry != nil {
		data.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.NumberType != nil {
		data.Set("NumberType", *params.NumberType)
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

	ps := &ListRegulationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Regulation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRegulation(params *ListRegulationParams) ([]NumbersV2Regulation, error) {
	if params == nil {
		params = &ListRegulationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRegulation(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []NumbersV2Regulation

	for response != nil {
		records = append(records, response.Results...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRegulationResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListRegulationResponse)
	}

	return records, err
}

// Streams Regulation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRegulation(params *ListRegulationParams) (chan NumbersV2Regulation, error) {
	if params == nil {
		params = &ListRegulationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRegulation(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan NumbersV2Regulation, 1)

	go func() {
		for response != nil {
			for item := range response.Results {
				channel <- response.Results[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRegulationResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListRegulationResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListRegulationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRegulationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
