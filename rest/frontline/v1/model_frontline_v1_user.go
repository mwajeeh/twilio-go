/*
 * Twilio - Frontline
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// FrontlineV1User struct for FrontlineV1User
type FrontlineV1User struct {
	// The avatar URL which will be shown in Frontline application
	Avatar *string `json:"avatar,omitempty"`
	// The string that you assigned to describe the User
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The string that identifies the resource's User
	Identity *string `json:"identity,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// Current state of this user
	State *string `json:"state,omitempty"`
	// An absolute URL for this user.
	Url *string `json:"url,omitempty"`
}
