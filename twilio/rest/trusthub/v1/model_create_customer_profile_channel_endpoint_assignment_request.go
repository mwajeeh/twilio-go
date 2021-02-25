/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateCustomerProfileChannelEndpointAssignmentRequest struct for CreateCustomerProfileChannelEndpointAssignmentRequest
type CreateCustomerProfileChannelEndpointAssignmentRequest struct {
	// The SID of an channel endpoint
	ChannelEndpointSid string `json:"ChannelEndpointSid"`
	// The type of channel endpoint. eg: phone-number
	ChannelEndpointType string `json:"ChannelEndpointType"`
}