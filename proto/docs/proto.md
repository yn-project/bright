# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [bright/endpoint/endpoint.proto](#bright_endpoint_endpoint-proto)
    - [Conds](#bright-endpoint-Conds)
    - [CreateEndpointRequest](#bright-endpoint-CreateEndpointRequest)
    - [CreateEndpointResponse](#bright-endpoint-CreateEndpointResponse)
    - [DeleteEndpointRequest](#bright-endpoint-DeleteEndpointRequest)
    - [DeleteEndpointResponse](#bright-endpoint-DeleteEndpointResponse)
    - [Endpoint](#bright-endpoint-Endpoint)
    - [EndpointReq](#bright-endpoint-EndpointReq)
    - [GetEndpointRequest](#bright-endpoint-GetEndpointRequest)
    - [GetEndpointResponse](#bright-endpoint-GetEndpointResponse)
    - [GetEndpointsRequest](#bright-endpoint-GetEndpointsRequest)
    - [GetEndpointsResponse](#bright-endpoint-GetEndpointsResponse)
  
    - [Manager](#bright-endpoint-Manager)
  
- [Scalar Value Types](#scalar-value-types)



<a name="bright_endpoint_endpoint-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## bright/endpoint/endpoint.proto



<a name="bright-endpoint-Conds"></a>

### Conds



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [bright.StringVal](#bright-StringVal) |  |  |
| Name | [bright.StringVal](#bright-StringVal) |  |  |
| Address | [bright.StringVal](#bright-StringVal) |  |  |
| State | [bright.StringVal](#bright-StringVal) |  |  |
| RPS | [bright.Uint32Val](#bright-Uint32Val) |  |  |
| IDs | [bright.StringSliceVal](#bright-StringSliceVal) |  |  |






<a name="bright-endpoint-CreateEndpointRequest"></a>

### CreateEndpointRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EndpointReq](#bright-endpoint-EndpointReq) |  |  |






<a name="bright-endpoint-CreateEndpointResponse"></a>

### CreateEndpointResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Endpoint](#bright-endpoint-Endpoint) |  |  |






<a name="bright-endpoint-DeleteEndpointRequest"></a>

### DeleteEndpointRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="bright-endpoint-DeleteEndpointResponse"></a>

### DeleteEndpointResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Endpoint](#bright-endpoint-Endpoint) |  |  |






<a name="bright-endpoint-Endpoint"></a>

### Endpoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Address | [string](#string) |  |  |
| State | [basetype.EndpointState](#basetype-EndpointState) |  |  |
| RPS | [uint32](#uint32) |  |  |
| Remark | [string](#string) |  |  |
| CreatedAt | [uint64](#uint64) |  |  |
| UpdatedAt | [uint64](#uint64) |  |  |






<a name="bright-endpoint-EndpointReq"></a>

### EndpointReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) | optional |  |
| Name | [string](#string) | optional |  |
| Address | [string](#string) | optional |  |
| State | [basetype.EndpointState](#basetype-EndpointState) | optional |  |
| RPS | [uint32](#uint32) | optional |  |
| Remark | [string](#string) | optional |  |






<a name="bright-endpoint-GetEndpointRequest"></a>

### GetEndpointRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="bright-endpoint-GetEndpointResponse"></a>

### GetEndpointResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Endpoint](#bright-endpoint-Endpoint) |  |  |






<a name="bright-endpoint-GetEndpointsRequest"></a>

### GetEndpointsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [Conds](#bright-endpoint-Conds) |  |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="bright-endpoint-GetEndpointsResponse"></a>

### GetEndpointsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Endpoint](#bright-endpoint-Endpoint) | repeated |  |
| Total | [uint32](#uint32) |  |  |





 

 

 


<a name="bright-endpoint-Manager"></a>

### Manager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateEndpoint | [CreateEndpointRequest](#bright-endpoint-CreateEndpointRequest) | [CreateEndpointResponse](#bright-endpoint-CreateEndpointResponse) |  |
| GetEndpoint | [GetEndpointRequest](#bright-endpoint-GetEndpointRequest) | [GetEndpointResponse](#bright-endpoint-GetEndpointResponse) |  |
| GetEndpoints | [GetEndpointsRequest](#bright-endpoint-GetEndpointsRequest) | [GetEndpointsResponse](#bright-endpoint-GetEndpointsResponse) |  |
| DeleteEndpoint | [DeleteEndpointRequest](#bright-endpoint-DeleteEndpointRequest) | [DeleteEndpointResponse](#bright-endpoint-DeleteEndpointResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

