/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListUserBindingResponse struct for ListUserBindingResponse
type ListUserBindingResponse struct {
	Bindings []IpMessagingV2UserBinding `json:"bindings,omitempty"`
	Meta     ListCredentialResponseMeta `json:"meta,omitempty"`
}
