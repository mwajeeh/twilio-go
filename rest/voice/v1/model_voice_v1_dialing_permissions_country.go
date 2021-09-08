/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// VoiceV1DialingPermissionsCountry struct for VoiceV1DialingPermissionsCountry
type VoiceV1DialingPermissionsCountry struct {
	// The name of the continent in which the country is located
	Continent *string `json:"continent,omitempty"`
	// The E.164 assigned country codes(s)
	CountryCodes *[]string `json:"country_codes,omitempty"`
	// Whether dialing to high-risk special services numbers is enabled
	HighRiskSpecialNumbersEnabled *bool `json:"high_risk_special_numbers_enabled,omitempty"`
	// Whether dialing to high-risk toll fraud numbers is enabled, else `false`
	HighRiskTollfraudNumbersEnabled *bool `json:"high_risk_tollfraud_numbers_enabled,omitempty"`
	// The ISO country code
	IsoCode *string `json:"iso_code,omitempty"`
	// A list of URLs related to this resource
	Links *map[string]interface{} `json:"links,omitempty"`
	// Whether dialing to low-risk numbers is enabled
	LowRiskNumbersEnabled *bool `json:"low_risk_numbers_enabled,omitempty"`
	// The name of the country
	Name *string `json:"name,omitempty"`
	// The absolute URL of this resource
	Url *string `json:"url,omitempty"`
}
