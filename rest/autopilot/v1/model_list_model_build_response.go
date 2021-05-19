/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListModelBuildResponse struct for ListModelBuildResponse
type ListModelBuildResponse struct {
	Meta        ListAssistantResponseMeta        `json:"meta,omitempty"`
	ModelBuilds []AutopilotV1AssistantModelBuild `json:"model_builds,omitempty"`
}
