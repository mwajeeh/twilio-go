/*
 * Twilio - Pricing
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
	"net/url"

	"strings"
)

// Optional parameters for the method 'FetchVoiceNumber'
type FetchVoiceNumberParams struct {
	// The origination phone number, in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, for which to fetch the origin-based voice pricing information. E.164 format consists of a + followed by the country code and subscriber number.
	OriginationNumber *string `json:"OriginationNumber,omitempty"`
}

func (params *FetchVoiceNumberParams) SetOriginationNumber(OriginationNumber string) *FetchVoiceNumberParams {
	params.OriginationNumber = &OriginationNumber
	return params
}

// Fetch pricing information for a specific destination and, optionally, origination phone number.
func (c *ApiService) FetchVoiceNumber(DestinationNumber string, params *FetchVoiceNumberParams) (*PricingV2VoiceNumber, error) {
	path := "/v2/Voice/Numbers/{DestinationNumber}"
	path = strings.Replace(path, "{"+"DestinationNumber"+"}", DestinationNumber, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.OriginationNumber != nil {
		data.Set("OriginationNumber", *params.OriginationNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV2VoiceNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
