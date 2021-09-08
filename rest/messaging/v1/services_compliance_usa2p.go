/*
 * Twilio - Messaging
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

// Optional parameters for the method 'CreateUsAppToPerson'
type CreateUsAppToPersonParams struct {
	// A2P Brand Registration SID
	BrandRegistrationSid *string `json:"BrandRegistrationSid,omitempty"`
	// A short description of what this SMS campaign does.
	Description *string `json:"Description,omitempty"`
	// Indicates that this SMS campaign will send messages that contain links.
	HasEmbeddedLinks *bool `json:"HasEmbeddedLinks,omitempty"`
	// Indicates that this SMS campaign will send messages that contain phone numbers.
	HasEmbeddedPhone *bool `json:"HasEmbeddedPhone,omitempty"`
	// Message samples, at least 2 and up to 5 sample messages, <=1024 chars each.
	MessageSamples *[]string `json:"MessageSamples,omitempty"`
	// A2P Campaign Use Case. Examples: [ 2FA, EMERGENCY, MARKETING..]
	UsAppToPersonUsecase *string `json:"UsAppToPersonUsecase,omitempty"`
}

func (params *CreateUsAppToPersonParams) SetBrandRegistrationSid(BrandRegistrationSid string) *CreateUsAppToPersonParams {
	params.BrandRegistrationSid = &BrandRegistrationSid
	return params
}
func (params *CreateUsAppToPersonParams) SetDescription(Description string) *CreateUsAppToPersonParams {
	params.Description = &Description
	return params
}
func (params *CreateUsAppToPersonParams) SetHasEmbeddedLinks(HasEmbeddedLinks bool) *CreateUsAppToPersonParams {
	params.HasEmbeddedLinks = &HasEmbeddedLinks
	return params
}
func (params *CreateUsAppToPersonParams) SetHasEmbeddedPhone(HasEmbeddedPhone bool) *CreateUsAppToPersonParams {
	params.HasEmbeddedPhone = &HasEmbeddedPhone
	return params
}
func (params *CreateUsAppToPersonParams) SetMessageSamples(MessageSamples []string) *CreateUsAppToPersonParams {
	params.MessageSamples = &MessageSamples
	return params
}
func (params *CreateUsAppToPersonParams) SetUsAppToPersonUsecase(UsAppToPersonUsecase string) *CreateUsAppToPersonParams {
	params.UsAppToPersonUsecase = &UsAppToPersonUsecase
	return params
}

func (c *ApiService) CreateUsAppToPerson(MessagingServiceSid string, params *CreateUsAppToPersonParams) (*MessagingV1UsAppToPerson, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.BrandRegistrationSid != nil {
		data.Set("BrandRegistrationSid", *params.BrandRegistrationSid)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.HasEmbeddedLinks != nil {
		data.Set("HasEmbeddedLinks", fmt.Sprint(*params.HasEmbeddedLinks))
	}
	if params != nil && params.HasEmbeddedPhone != nil {
		data.Set("HasEmbeddedPhone", fmt.Sprint(*params.HasEmbeddedPhone))
	}
	if params != nil && params.MessageSamples != nil {
		for _, item := range *params.MessageSamples {
			data.Add("MessageSamples", item)
		}
	}
	if params != nil && params.UsAppToPersonUsecase != nil {
		data.Set("UsAppToPersonUsecase", *params.UsAppToPersonUsecase)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1UsAppToPerson{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteUsAppToPerson(MessagingServiceSid string, Sid string) error {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p/{Sid}"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)
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

func (c *ApiService) FetchUsAppToPerson(MessagingServiceSid string, Sid string) (*MessagingV1UsAppToPerson, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p/{Sid}"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1UsAppToPerson{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUsAppToPerson'
type ListUsAppToPersonParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUsAppToPersonParams) SetPageSize(PageSize int) *ListUsAppToPersonParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUsAppToPersonParams) SetLimit(Limit int) *ListUsAppToPersonParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of UsAppToPerson records from the API. Request is executed immediately.
func (c *ApiService) PageUsAppToPerson(MessagingServiceSid string, params *ListUsAppToPersonParams, pageToken, pageNumber string) (*ListUsAppToPersonResponse, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p"

	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

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

	ps := &ListUsAppToPersonResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UsAppToPerson records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUsAppToPerson(MessagingServiceSid string, params *ListUsAppToPersonParams) ([]MessagingV1UsAppToPerson, error) {
	if params == nil {
		params = &ListUsAppToPersonParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageUsAppToPerson(MessagingServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []MessagingV1UsAppToPerson

	for response != nil {
		records = append(records, response.Compliance...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListUsAppToPersonResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListUsAppToPersonResponse)
	}

	return records, err
}

// Streams UsAppToPerson records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUsAppToPerson(MessagingServiceSid string, params *ListUsAppToPersonParams) (chan MessagingV1UsAppToPerson, error) {
	if params == nil {
		params = &ListUsAppToPersonParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageUsAppToPerson(MessagingServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan MessagingV1UsAppToPerson, 1)

	go func() {
		for response != nil {
			for item := range response.Compliance {
				channel <- response.Compliance[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListUsAppToPersonResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListUsAppToPersonResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListUsAppToPersonResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsAppToPersonResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
