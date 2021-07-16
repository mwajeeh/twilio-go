/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
	"time"
)

// Delete a specific User Conversation.
func (c *ApiService) DeleteServiceUserConversation(ChatServiceSid string, UserSid string, ConversationSid string) error {
	path := "/v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific User Conversation.
func (c *ApiService) FetchServiceUserConversation(ChatServiceSid string, UserSid string, ConversationSid string) (*ConversationsV1ServiceServiceUserServiceUserConversation, error) {
	path := "/v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceServiceUserServiceUserConversation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListServiceUserConversation'
type ListServiceUserConversationParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListServiceUserConversationParams) SetPageSize(PageSize int) *ListServiceUserConversationParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all User Conversations for the User.
func (c *ApiService) ListServiceUserConversation(ChatServiceSid string, UserSid string, params *ListServiceUserConversationParams) (*ListServiceUserConversationResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceUserConversationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateServiceUserConversation'
type UpdateServiceUserConversationParams struct {
	// The index of the last Message in the Conversation that the Participant has read.
	LastReadMessageIndex *int `json:"LastReadMessageIndex,omitempty"`
	// The date of the last message read in conversation by the user, given in ISO 8601 format.
	LastReadTimestamp *time.Time `json:"LastReadTimestamp,omitempty"`
	// The Notification Level of this User Conversation. One of `default` or `muted`.
	NotificationLevel *string `json:"NotificationLevel,omitempty"`
}

func (params *UpdateServiceUserConversationParams) SetLastReadMessageIndex(LastReadMessageIndex int) *UpdateServiceUserConversationParams {
	params.LastReadMessageIndex = &LastReadMessageIndex
	return params
}
func (params *UpdateServiceUserConversationParams) SetLastReadTimestamp(LastReadTimestamp time.Time) *UpdateServiceUserConversationParams {
	params.LastReadTimestamp = &LastReadTimestamp
	return params
}
func (params *UpdateServiceUserConversationParams) SetNotificationLevel(NotificationLevel string) *UpdateServiceUserConversationParams {
	params.NotificationLevel = &NotificationLevel
	return params
}

// Update a specific User Conversation.
func (c *ApiService) UpdateServiceUserConversation(ChatServiceSid string, UserSid string, ConversationSid string, params *UpdateServiceUserConversationParams) (*ConversationsV1ServiceServiceUserServiceUserConversation, error) {
	path := "/v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.LastReadMessageIndex != nil {
		data.Set("LastReadMessageIndex", fmt.Sprint(*params.LastReadMessageIndex))
	}
	if params != nil && params.LastReadTimestamp != nil {
		data.Set("LastReadTimestamp", fmt.Sprint((*params.LastReadTimestamp).Format(time.RFC3339)))
	}
	if params != nil && params.NotificationLevel != nil {
		data.Set("NotificationLevel", *params.NotificationLevel)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceServiceUserServiceUserConversation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}