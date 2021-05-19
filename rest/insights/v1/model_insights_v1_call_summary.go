/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// InsightsV1CallSummary struct for InsightsV1CallSummary
type InsightsV1CallSummary struct {
	AccountSid      *string                 `json:"account_sid,omitempty"`
	Attributes      *map[string]interface{} `json:"attributes,omitempty"`
	CallSid         *string                 `json:"call_sid,omitempty"`
	CallState       *string                 `json:"call_state,omitempty"`
	CallType        *string                 `json:"call_type,omitempty"`
	CarrierEdge     *map[string]interface{} `json:"carrier_edge,omitempty"`
	ClientEdge      *map[string]interface{} `json:"client_edge,omitempty"`
	ConnectDuration *int32                  `json:"connect_duration,omitempty"`
	CreatedTime     *time.Time              `json:"created_time,omitempty"`
	Duration        *int32                  `json:"duration,omitempty"`
	EndTime         *time.Time              `json:"end_time,omitempty"`
	From            *map[string]interface{} `json:"from,omitempty"`
	ProcessingState *string                 `json:"processing_state,omitempty"`
	Properties      *map[string]interface{} `json:"properties,omitempty"`
	SdkEdge         *map[string]interface{} `json:"sdk_edge,omitempty"`
	SipEdge         *map[string]interface{} `json:"sip_edge,omitempty"`
	StartTime       *time.Time              `json:"start_time,omitempty"`
	Tags            *[]string               `json:"tags,omitempty"`
	To              *map[string]interface{} `json:"to,omitempty"`
	Trust           *map[string]interface{} `json:"trust,omitempty"`
	Url             *string                 `json:"url,omitempty"`
}
