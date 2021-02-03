# Go API client for openapi

This is the public Twilio REST API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/twilio-oai](https://github.com/twilio/twilio-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: com.twilio.oai.TwilioGoGenerator
For more information, please visit [https://support.twilio.com](https://support.twilio.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateSink**](docs/DefaultApi.md#createsink) | **Post** /v1/Sinks | 
*DefaultApi* | [**CreateSinkTest**](docs/DefaultApi.md#createsinktest) | **Post** /v1/Sinks/{Sid}/Test | 
*DefaultApi* | [**CreateSinkValidate**](docs/DefaultApi.md#createsinkvalidate) | **Post** /v1/Sinks/{Sid}/Validate | 
*DefaultApi* | [**CreateSubscription**](docs/DefaultApi.md#createsubscription) | **Post** /v1/Subscriptions | 
*DefaultApi* | [**DeleteSink**](docs/DefaultApi.md#deletesink) | **Delete** /v1/Sinks/{Sid} | 
*DefaultApi* | [**DeleteSubscription**](docs/DefaultApi.md#deletesubscription) | **Delete** /v1/Subscriptions/{Sid} | 
*DefaultApi* | [**FetchEventType**](docs/DefaultApi.md#fetcheventtype) | **Get** /v1/Types/{Type} | 
*DefaultApi* | [**FetchSchema**](docs/DefaultApi.md#fetchschema) | **Get** /v1/Schemas/{Id} | 
*DefaultApi* | [**FetchSink**](docs/DefaultApi.md#fetchsink) | **Get** /v1/Sinks/{Sid} | 
*DefaultApi* | [**FetchSubscription**](docs/DefaultApi.md#fetchsubscription) | **Get** /v1/Subscriptions/{Sid} | 
*DefaultApi* | [**FetchVersion**](docs/DefaultApi.md#fetchversion) | **Get** /v1/Schemas/{Id}/Versions/{SchemaVersion} | 
*DefaultApi* | [**ListEventType**](docs/DefaultApi.md#listeventtype) | **Get** /v1/Types | 
*DefaultApi* | [**ListSink**](docs/DefaultApi.md#listsink) | **Get** /v1/Sinks | 
*DefaultApi* | [**ListSubscribedEvent**](docs/DefaultApi.md#listsubscribedevent) | **Get** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents | 
*DefaultApi* | [**ListSubscription**](docs/DefaultApi.md#listsubscription) | **Get** /v1/Subscriptions | 
*DefaultApi* | [**ListVersion**](docs/DefaultApi.md#listversion) | **Get** /v1/Schemas/{Id}/Versions | 
*DefaultApi* | [**UpdateSubscription**](docs/DefaultApi.md#updatesubscription) | **Post** /v1/Subscriptions/{Sid} | 


## Documentation For Models

 - [CreateSinkRequest](docs/CreateSinkRequest.md)
 - [CreateSinkValidateRequest](docs/CreateSinkValidateRequest.md)
 - [CreateSubscriptionRequest](docs/CreateSubscriptionRequest.md)
 - [EventsV1EventType](docs/EventsV1EventType.md)
 - [EventsV1Schema](docs/EventsV1Schema.md)
 - [EventsV1SchemaVersion](docs/EventsV1SchemaVersion.md)
 - [EventsV1Sink](docs/EventsV1Sink.md)
 - [EventsV1SinkSinkTest](docs/EventsV1SinkSinkTest.md)
 - [EventsV1SinkSinkValidate](docs/EventsV1SinkSinkValidate.md)
 - [EventsV1Subscription](docs/EventsV1Subscription.md)
 - [EventsV1SubscriptionSubscribedEvent](docs/EventsV1SubscriptionSubscribedEvent.md)
 - [ListEventTypeResponse](docs/ListEventTypeResponse.md)
 - [ListSinkResponse](docs/ListSinkResponse.md)
 - [ListSubscribedEventResponse](docs/ListSubscribedEventResponse.md)
 - [ListSubscriptionResponse](docs/ListSubscriptionResponse.md)
 - [ListVersionResponse](docs/ListVersionResponse.md)
 - [ListVersionResponseMeta](docs/ListVersionResponseMeta.md)
 - [UpdateSubscriptionRequest](docs/UpdateSubscriptionRequest.md)


## Documentation For Authorization



## accountSid_authToken

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Author

support@twilio.com
