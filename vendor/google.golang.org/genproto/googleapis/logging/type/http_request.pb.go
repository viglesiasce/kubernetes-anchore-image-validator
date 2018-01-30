// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/logging/type/http_request.proto

/*
Package ltype is a generated protocol buffer package.

It is generated from these files:
	google/logging/type/http_request.proto
	google/logging/type/log_severity.proto

It has these top-level messages:
	HttpRequest
*/
package ltype

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A common proto for logging HTTP requests. Only contains semantics
// defined by the HTTP specification. Product-specific logging
// information MUST be defined in a separate message.
type HttpRequest struct {
	// The request method. Examples: `"GET"`, `"HEAD"`, `"PUT"`, `"POST"`.
	RequestMethod string `protobuf:"bytes,1,opt,name=request_method,json=requestMethod" json:"request_method,omitempty"`
	// The scheme (http, https), the host name, the path and the query
	// portion of the URL that was requested.
	// Example: `"http://example.com/some/info?color=red"`.
	RequestUrl string `protobuf:"bytes,2,opt,name=request_url,json=requestUrl" json:"request_url,omitempty"`
	// The size of the HTTP request message in bytes, including the request
	// headers and the request body.
	RequestSize int64 `protobuf:"varint,3,opt,name=request_size,json=requestSize" json:"request_size,omitempty"`
	// The response code indicating the status of response.
	// Examples: 200, 404.
	Status int32 `protobuf:"varint,4,opt,name=status" json:"status,omitempty"`
	// The size of the HTTP response message sent back to the client, in bytes,
	// including the response headers and the response body.
	ResponseSize int64 `protobuf:"varint,5,opt,name=response_size,json=responseSize" json:"response_size,omitempty"`
	// The user agent sent by the client. Example:
	// `"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; Q312461; .NET CLR 1.0.3705)"`.
	UserAgent string `protobuf:"bytes,6,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	// The IP address (IPv4 or IPv6) of the client that issued the HTTP
	// request. Examples: `"192.168.1.1"`, `"FE80::0202:B3FF:FE1E:8329"`.
	RemoteIp string `protobuf:"bytes,7,opt,name=remote_ip,json=remoteIp" json:"remote_ip,omitempty"`
	// The IP address (IPv4 or IPv6) of the origin server that the request was
	// sent to.
	ServerIp string `protobuf:"bytes,13,opt,name=server_ip,json=serverIp" json:"server_ip,omitempty"`
	// The referer URL of the request, as defined in
	// [HTTP/1.1 Header Field Definitions](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).
	Referer string `protobuf:"bytes,8,opt,name=referer" json:"referer,omitempty"`
	// The request processing latency on the server, from the time the request was
	// received until the response was sent.
	Latency *google_protobuf1.Duration `protobuf:"bytes,14,opt,name=latency" json:"latency,omitempty"`
	// Whether or not a cache lookup was attempted.
	CacheLookup bool `protobuf:"varint,11,opt,name=cache_lookup,json=cacheLookup" json:"cache_lookup,omitempty"`
	// Whether or not an entity was served from cache
	// (with or without validation).
	CacheHit bool `protobuf:"varint,9,opt,name=cache_hit,json=cacheHit" json:"cache_hit,omitempty"`
	// Whether or not the response was validated with the origin server before
	// being served from cache. This field is only meaningful if `cache_hit` is
	// True.
	CacheValidatedWithOriginServer bool `protobuf:"varint,10,opt,name=cache_validated_with_origin_server,json=cacheValidatedWithOriginServer" json:"cache_validated_with_origin_server,omitempty"`
	// The number of HTTP response bytes inserted into cache. Set only when a
	// cache fill was attempted.
	CacheFillBytes int64 `protobuf:"varint,12,opt,name=cache_fill_bytes,json=cacheFillBytes" json:"cache_fill_bytes,omitempty"`
}

func (m *HttpRequest) Reset()                    { *m = HttpRequest{} }
func (m *HttpRequest) String() string            { return proto.CompactTextString(m) }
func (*HttpRequest) ProtoMessage()               {}
func (*HttpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HttpRequest) GetRequestMethod() string {
	if m != nil {
		return m.RequestMethod
	}
	return ""
}

func (m *HttpRequest) GetRequestUrl() string {
	if m != nil {
		return m.RequestUrl
	}
	return ""
}

func (m *HttpRequest) GetRequestSize() int64 {
	if m != nil {
		return m.RequestSize
	}
	return 0
}

func (m *HttpRequest) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *HttpRequest) GetResponseSize() int64 {
	if m != nil {
		return m.ResponseSize
	}
	return 0
}

func (m *HttpRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *HttpRequest) GetRemoteIp() string {
	if m != nil {
		return m.RemoteIp
	}
	return ""
}

func (m *HttpRequest) GetServerIp() string {
	if m != nil {
		return m.ServerIp
	}
	return ""
}

func (m *HttpRequest) GetReferer() string {
	if m != nil {
		return m.Referer
	}
	return ""
}

func (m *HttpRequest) GetLatency() *google_protobuf1.Duration {
	if m != nil {
		return m.Latency
	}
	return nil
}

func (m *HttpRequest) GetCacheLookup() bool {
	if m != nil {
		return m.CacheLookup
	}
	return false
}

func (m *HttpRequest) GetCacheHit() bool {
	if m != nil {
		return m.CacheHit
	}
	return false
}

func (m *HttpRequest) GetCacheValidatedWithOriginServer() bool {
	if m != nil {
		return m.CacheValidatedWithOriginServer
	}
	return false
}

func (m *HttpRequest) GetCacheFillBytes() int64 {
	if m != nil {
		return m.CacheFillBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*HttpRequest)(nil), "google.logging.type.HttpRequest")
}

func init() { proto.RegisterFile("google/logging/type/http_request.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x5b, 0x6b, 0x14, 0x31,
	0x14, 0xc7, 0x99, 0x5e, 0xf6, 0x92, 0xbd, 0x50, 0x22, 0x68, 0x5a, 0xb5, 0xae, 0x15, 0x65, 0x9e,
	0x66, 0xc0, 0xbe, 0x08, 0x3e, 0xb9, 0x8a, 0xb6, 0x52, 0xb1, 0x4c, 0xbd, 0x80, 0x2f, 0x43, 0x76,
	0xf7, 0x6c, 0x26, 0x98, 0x9d, 0xc4, 0x24, 0x53, 0xd9, 0xbe, 0xfa, 0x11, 0xfc, 0x16, 0x7e, 0x4a,
	0x99, 0x93, 0x0c, 0x28, 0xf4, 0x65, 0x21, 0xbf, 0xff, 0xef, 0x3f, 0x67, 0xf6, 0x4c, 0xc8, 0x33,
	0xa1, 0xb5, 0x50, 0x90, 0x2b, 0x2d, 0x84, 0xac, 0x45, 0xee, 0xb7, 0x06, 0xf2, 0xca, 0x7b, 0x53,
	0x5a, 0xf8, 0xd1, 0x80, 0xf3, 0x99, 0xb1, 0xda, 0x6b, 0x7a, 0x27, 0x78, 0x59, 0xf4, 0xb2, 0xd6,
	0x3b, 0x7a, 0x10, 0xcb, 0xdc, 0xc8, 0x9c, 0xd7, 0xb5, 0xf6, 0xdc, 0x4b, 0x5d, 0xbb, 0x50, 0x39,
	0x3a, 0x8e, 0x29, 0x9e, 0x16, 0xcd, 0x3a, 0x5f, 0x35, 0x16, 0x85, 0x90, 0x9f, 0xfc, 0xde, 0x23,
	0xa3, 0x33, 0xef, 0x4d, 0x11, 0x06, 0xd1, 0xa7, 0x64, 0x1a, 0x67, 0x96, 0x1b, 0xf0, 0x95, 0x5e,
	0xb1, 0x64, 0x96, 0xa4, 0xc3, 0x62, 0x12, 0xe9, 0x07, 0x84, 0xf4, 0x11, 0x19, 0x75, 0x5a, 0x63,
	0x15, 0xdb, 0x41, 0x87, 0x44, 0xf4, 0xd9, 0x2a, 0xfa, 0x98, 0x8c, 0x3b, 0xc1, 0xc9, 0x1b, 0x60,
	0xbb, 0xb3, 0x24, 0xdd, 0x2d, 0xba, 0xd2, 0x95, 0xbc, 0x01, 0x7a, 0x97, 0xf4, 0x9c, 0xe7, 0xbe,
	0x71, 0x6c, 0x6f, 0x96, 0xa4, 0xfb, 0x45, 0x3c, 0xd1, 0x27, 0x64, 0x62, 0xc1, 0x19, 0x5d, 0x3b,
	0x08, 0xdd, 0x7d, 0xec, 0x8e, 0x3b, 0x88, 0xe5, 0x87, 0x84, 0x34, 0x0e, 0x6c, 0xc9, 0x05, 0xd4,
	0x9e, 0xf5, 0x70, 0xfe, 0xb0, 0x25, 0xaf, 0x5a, 0x40, 0xef, 0x93, 0xa1, 0x85, 0x8d, 0xf6, 0x50,
	0x4a, 0xc3, 0xfa, 0x98, 0x0e, 0x02, 0x38, 0x37, 0x6d, 0xe8, 0xc0, 0x5e, 0x83, 0x6d, 0xc3, 0x49,
	0x08, 0x03, 0x38, 0x37, 0x94, 0x91, 0xbe, 0x85, 0x35, 0x58, 0xb0, 0x6c, 0x80, 0x51, 0x77, 0xa4,
	0xa7, 0xa4, 0xaf, 0xb8, 0x87, 0x7a, 0xb9, 0x65, 0xd3, 0x59, 0x92, 0x8e, 0x9e, 0x1f, 0x66, 0xf1,
	0x7b, 0x74, 0xcb, 0xcd, 0xde, 0xc4, 0xe5, 0x16, 0x9d, 0xd9, 0xee, 0x61, 0xc9, 0x97, 0x15, 0x94,
	0x4a, 0xeb, 0xef, 0x8d, 0x61, 0xa3, 0x59, 0x92, 0x0e, 0x8a, 0x11, 0xb2, 0x0b, 0x44, 0xed, 0xeb,
	0x04, 0xa5, 0x92, 0x9e, 0x0d, 0x31, 0x1f, 0x20, 0x38, 0x93, 0x9e, 0xbe, 0x27, 0x27, 0x21, 0xbc,
	0xe6, 0x4a, 0xae, 0xb8, 0x87, 0x55, 0xf9, 0x53, 0xfa, 0xaa, 0xd4, 0x56, 0x0a, 0x59, 0x97, 0xe1,
	0xb5, 0x19, 0xc1, 0xd6, 0x31, 0x9a, 0x5f, 0x3a, 0xf1, 0xab, 0xf4, 0xd5, 0x47, 0xd4, 0xae, 0xd0,
	0xa2, 0x29, 0x39, 0x08, 0xcf, 0x5a, 0x4b, 0xa5, 0xca, 0xc5, 0xd6, 0x83, 0x63, 0x63, 0xdc, 0xed,
	0x14, 0xf9, 0x5b, 0xa9, 0xd4, 0xbc, 0xa5, 0xf3, 0x5f, 0x09, 0xb9, 0xb7, 0xd4, 0x9b, 0xec, 0x96,
	0xfb, 0x36, 0x3f, 0xf8, 0xe7, 0xba, 0x5c, 0xb6, 0x7f, 0xfc, 0x32, 0xf9, 0xf6, 0x22, 0x8a, 0x42,
	0x2b, 0x5e, 0x8b, 0x4c, 0x5b, 0x91, 0x0b, 0xa8, 0x71, 0x2d, 0x79, 0x88, 0xb8, 0x91, 0xee, 0xbf,
	0xfb, 0xfd, 0x52, 0xb5, 0xbf, 0x7f, 0x76, 0x0e, 0xdf, 0x85, 0xea, 0x6b, 0xa5, 0x9b, 0x55, 0x76,
	0x11, 0x27, 0x7d, 0xda, 0x1a, 0x58, 0xf4, 0xf0, 0x01, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x6f, 0xb5, 0x28, 0xee, 0x1f, 0x03, 0x00, 0x00,
}
