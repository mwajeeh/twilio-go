# Go API client for openapi

This is the public Twilio REST API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/twilio-oai](https://github.com/twilio/twilio-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.8.0
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
*DefaultApi* | [**CreateCommand**](docs/DefaultApi.md#createcommand) | **Post** /v1/Commands | 
*DefaultApi* | [**CreateRatePlan**](docs/DefaultApi.md#createrateplan) | **Post** /v1/RatePlans | 
*DefaultApi* | [**DeleteCommand**](docs/DefaultApi.md#deletecommand) | **Delete** /v1/Commands/{Sid} | 
*DefaultApi* | [**DeleteRatePlan**](docs/DefaultApi.md#deleterateplan) | **Delete** /v1/RatePlans/{Sid} | 
*DefaultApi* | [**DeleteSim**](docs/DefaultApi.md#deletesim) | **Delete** /v1/Sims/{Sid} | 
*DefaultApi* | [**FetchCommand**](docs/DefaultApi.md#fetchcommand) | **Get** /v1/Commands/{Sid} | 
*DefaultApi* | [**FetchRatePlan**](docs/DefaultApi.md#fetchrateplan) | **Get** /v1/RatePlans/{Sid} | 
*DefaultApi* | [**FetchSim**](docs/DefaultApi.md#fetchsim) | **Get** /v1/Sims/{Sid} | 
*DefaultApi* | [**ListAccountUsageRecord**](docs/DefaultApi.md#listaccountusagerecord) | **Get** /v1/UsageRecords | 
*DefaultApi* | [**ListCommand**](docs/DefaultApi.md#listcommand) | **Get** /v1/Commands | 
*DefaultApi* | [**ListDataSession**](docs/DefaultApi.md#listdatasession) | **Get** /v1/Sims/{SimSid}/DataSessions | 
*DefaultApi* | [**ListRatePlan**](docs/DefaultApi.md#listrateplan) | **Get** /v1/RatePlans | 
*DefaultApi* | [**ListSim**](docs/DefaultApi.md#listsim) | **Get** /v1/Sims | 
*DefaultApi* | [**ListUsageRecord**](docs/DefaultApi.md#listusagerecord) | **Get** /v1/Sims/{SimSid}/UsageRecords | 
*DefaultApi* | [**UpdateRatePlan**](docs/DefaultApi.md#updaterateplan) | **Post** /v1/RatePlans/{Sid} | 
*DefaultApi* | [**UpdateSim**](docs/DefaultApi.md#updatesim) | **Post** /v1/Sims/{Sid} | 


## Documentation For Models

 - [CommandMode](docs/CommandMode.md)
 - [CreateCommandRequest](docs/CreateCommandRequest.md)
 - [CreateRatePlanRequest](docs/CreateRatePlanRequest.md)
 - [Direction](docs/Direction.md)
 - [HttpMethod](docs/HttpMethod.md)
 - [ListAccountUsageRecordResponse](docs/ListAccountUsageRecordResponse.md)
 - [ListCommandResponse](docs/ListCommandResponse.md)
 - [ListCommandResponseMeta](docs/ListCommandResponseMeta.md)
 - [ListDataSessionResponse](docs/ListDataSessionResponse.md)
 - [ListRatePlanResponse](docs/ListRatePlanResponse.md)
 - [ListSimResponse](docs/ListSimResponse.md)
 - [ListUsageRecordResponse](docs/ListUsageRecordResponse.md)
 - [ResetStatus](docs/ResetStatus.md)
 - [Status](docs/Status.md)
 - [Transport](docs/Transport.md)
 - [UpdateRatePlanRequest](docs/UpdateRatePlanRequest.md)
 - [UpdateSimRequest](docs/UpdateSimRequest.md)
 - [WirelessV1AccountUsageRecord](docs/WirelessV1AccountUsageRecord.md)
 - [WirelessV1Command](docs/WirelessV1Command.md)
 - [WirelessV1RatePlan](docs/WirelessV1RatePlan.md)
 - [WirelessV1Sim](docs/WirelessV1Sim.md)
 - [WirelessV1SimDataSession](docs/WirelessV1SimDataSession.md)
 - [WirelessV1SimUsageRecord](docs/WirelessV1SimUsageRecord.md)


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
