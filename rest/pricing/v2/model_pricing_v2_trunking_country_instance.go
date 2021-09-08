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

// PricingV2TrunkingCountryInstance struct for PricingV2TrunkingCountryInstance
type PricingV2TrunkingCountryInstance struct {
	// The name of the country
	Country *string `json:"country,omitempty"`
	// The ISO country code
	IsoCountry *string `json:"iso_country,omitempty"`
	// The list of OriginatingCallPrice records
	OriginatingCallPrices *[]PricingV2TrunkingCountryInstanceOriginatingCallPrices `json:"originating_call_prices,omitempty"`
	// The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy)
	PriceUnit *string `json:"price_unit,omitempty"`
	// The list of TerminatingPrefixPrice records
	TerminatingPrefixPrices *[]PricingV2TrunkingCountryInstanceTerminatingPrefixPrices `json:"terminating_prefix_prices,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
