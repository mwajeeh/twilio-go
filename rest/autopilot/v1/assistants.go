/*
 * Twilio - Autopilot
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

// Optional parameters for the method 'CreateAssistant'
type CreateAssistantParams struct {
	// Reserved.
	CallbackEvents *string `json:"CallbackEvents,omitempty"`
	// Reserved.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// A JSON object that defines the Assistant's [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks.
	Defaults *map[string]interface{} `json:"Defaults,omitempty"`
	// A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether queries should be logged and kept after training. Can be: `true` or `false` and defaults to `true`. If `true`, queries are stored for 30 days, and then deleted. If `false`, no queries are stored.
	LogQueries *bool `json:"LogQueries,omitempty"`
	// The JSON string that defines the Assistant's [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet)
	StyleSheet *map[string]interface{} `json:"StyleSheet,omitempty"`
	// An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. The first 64 characters must be unique.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateAssistantParams) SetCallbackEvents(CallbackEvents string) *CreateAssistantParams {
	params.CallbackEvents = &CallbackEvents
	return params
}
func (params *CreateAssistantParams) SetCallbackUrl(CallbackUrl string) *CreateAssistantParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *CreateAssistantParams) SetDefaults(Defaults map[string]interface{}) *CreateAssistantParams {
	params.Defaults = &Defaults
	return params
}
func (params *CreateAssistantParams) SetFriendlyName(FriendlyName string) *CreateAssistantParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateAssistantParams) SetLogQueries(LogQueries bool) *CreateAssistantParams {
	params.LogQueries = &LogQueries
	return params
}
func (params *CreateAssistantParams) SetStyleSheet(StyleSheet map[string]interface{}) *CreateAssistantParams {
	params.StyleSheet = &StyleSheet
	return params
}
func (params *CreateAssistantParams) SetUniqueName(UniqueName string) *CreateAssistantParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) CreateAssistant(params *CreateAssistantParams) (*AutopilotV1Assistant, error) {
	path := "/v1/Assistants"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CallbackEvents != nil {
		data.Set("CallbackEvents", *params.CallbackEvents)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.Defaults != nil {
		v, err := json.Marshal(params.Defaults)

		if err != nil {
			return nil, err
		}

		data.Set("Defaults", string(v))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.LogQueries != nil {
		data.Set("LogQueries", fmt.Sprint(*params.LogQueries))
	}
	if params != nil && params.StyleSheet != nil {
		v, err := json.Marshal(params.StyleSheet)

		if err != nil {
			return nil, err
		}

		data.Set("StyleSheet", string(v))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Assistant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteAssistant(Sid string) error {
	path := "/v1/Assistants/{Sid}"
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

func (c *ApiService) FetchAssistant(Sid string) (*AutopilotV1Assistant, error) {
	path := "/v1/Assistants/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Assistant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAssistant'
type ListAssistantParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAssistantParams) SetPageSize(PageSize int) *ListAssistantParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAssistantParams) SetLimit(Limit int) *ListAssistantParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Assistant records from the API. Request is executed immediately.
func (c *ApiService) PageAssistant(params *ListAssistantParams, pageToken, pageNumber string) (*ListAssistantResponse, error) {
	path := "/v1/Assistants"

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

	ps := &ListAssistantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Assistant records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAssistant(params *ListAssistantParams) ([]AutopilotV1Assistant, error) {
	if params == nil {
		params = &ListAssistantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAssistant(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []AutopilotV1Assistant

	for response != nil {
		records = append(records, response.Assistants...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAssistantResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListAssistantResponse)
	}

	return records, err
}

// Streams Assistant records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAssistant(params *ListAssistantParams) (chan AutopilotV1Assistant, error) {
	if params == nil {
		params = &ListAssistantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAssistant(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan AutopilotV1Assistant, 1)

	go func() {
		for response != nil {
			for item := range response.Assistants {
				channel <- response.Assistants[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAssistantResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListAssistantResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListAssistantResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAssistantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateAssistant'
type UpdateAssistantParams struct {
	// Reserved.
	CallbackEvents *string `json:"CallbackEvents,omitempty"`
	// Reserved.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// A JSON object that defines the Assistant's [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks.
	Defaults *map[string]interface{} `json:"Defaults,omitempty"`
	// A string describing the state of the assistant.
	DevelopmentStage *string `json:"DevelopmentStage,omitempty"`
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether queries should be logged and kept after training. Can be: `true` or `false` and defaults to `true`. If `true`, queries are stored for 30 days, and then deleted. If `false`, no queries are stored.
	LogQueries *bool `json:"LogQueries,omitempty"`
	// The JSON string that defines the Assistant's [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet)
	StyleSheet *map[string]interface{} `json:"StyleSheet,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used as an alternative to the `sid` in the URL path to address the resource. The first 64 characters must be unique.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateAssistantParams) SetCallbackEvents(CallbackEvents string) *UpdateAssistantParams {
	params.CallbackEvents = &CallbackEvents
	return params
}
func (params *UpdateAssistantParams) SetCallbackUrl(CallbackUrl string) *UpdateAssistantParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *UpdateAssistantParams) SetDefaults(Defaults map[string]interface{}) *UpdateAssistantParams {
	params.Defaults = &Defaults
	return params
}
func (params *UpdateAssistantParams) SetDevelopmentStage(DevelopmentStage string) *UpdateAssistantParams {
	params.DevelopmentStage = &DevelopmentStage
	return params
}
func (params *UpdateAssistantParams) SetFriendlyName(FriendlyName string) *UpdateAssistantParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateAssistantParams) SetLogQueries(LogQueries bool) *UpdateAssistantParams {
	params.LogQueries = &LogQueries
	return params
}
func (params *UpdateAssistantParams) SetStyleSheet(StyleSheet map[string]interface{}) *UpdateAssistantParams {
	params.StyleSheet = &StyleSheet
	return params
}
func (params *UpdateAssistantParams) SetUniqueName(UniqueName string) *UpdateAssistantParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) UpdateAssistant(Sid string, params *UpdateAssistantParams) (*AutopilotV1Assistant, error) {
	path := "/v1/Assistants/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CallbackEvents != nil {
		data.Set("CallbackEvents", *params.CallbackEvents)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.Defaults != nil {
		v, err := json.Marshal(params.Defaults)

		if err != nil {
			return nil, err
		}

		data.Set("Defaults", string(v))
	}
	if params != nil && params.DevelopmentStage != nil {
		data.Set("DevelopmentStage", *params.DevelopmentStage)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.LogQueries != nil {
		data.Set("LogQueries", fmt.Sprint(*params.LogQueries))
	}
	if params != nil && params.StyleSheet != nil {
		v, err := json.Marshal(params.StyleSheet)

		if err != nil {
			return nil, err
		}

		data.Set("StyleSheet", string(v))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Assistant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
