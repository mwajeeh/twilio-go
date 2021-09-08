/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010SipCredential struct for ApiV2010SipCredential
type ApiV2010SipCredential struct {
	// The unique id of the Account that is responsible for this resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique id that identifies the credential list that includes this credential
	CredentialListSid *string `json:"credential_list_sid,omitempty"`
	// The date that this resource was created, given as GMT in RFC 2822 format.
	DateCreated *string `json:"date_created,omitempty"`
	// The date that this resource was last updated, given as GMT in RFC 2822 format.
	DateUpdated *string `json:"date_updated,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// The URI for this resource, relative to https://api.twilio.com
	Uri *string `json:"uri,omitempty"`
	// The username for this credential.
	Username *string `json:"username,omitempty"`
}
