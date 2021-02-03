/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateSyncListPermissionRequest struct for UpdateSyncListPermissionRequest
type UpdateSyncListPermissionRequest struct {
	// Whether the identity can delete the Sync List. Default value is `false`.
	Manage bool `json:"Manage"`
	// Whether the identity can read the Sync List and its Items. Default value is `false`.
	Read bool `json:"Read"`
	// Whether the identity can create, update, and delete Items in the Sync List. Default value is `false`.
	Write bool `json:"Write"`
}