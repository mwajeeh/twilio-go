/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListRatePlanResponse struct for ListRatePlanResponse
type ListRatePlanResponse struct {
	Meta ListDayResponseMeta `json:"Meta,omitempty"`
	RatePlans []PreviewWirelessRatePlan `json:"RatePlans,omitempty"`
}