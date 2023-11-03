package types

type Json map[string]interface{}

type StatusCode int

func (code *StatusCode) Parse() int {
	return int(*code)
}

// Success
const (
	// (HTTP 200) status code is one of the most commonly used status codes and indicates a successful response.
	SC_S_OK StatusCode = 200
	// (HTTP 201) status code is often used when a new resource, like a new user account, has been successfully created.
	SC_S_Created StatusCode = 201
	// (HTTP 202) status code might be used to indicate that the server has accepted a request for processing.
	SC_S_Accepted StatusCode = 202
	// (HTTP 204) status code indicates a successful request where there is no content to return in the response body, commonly used in delete or update operations.
	SC_S_NoContent StatusCode = 204
)

// Client Errors
const (
	// (HTTP 400) status code indicates that the server could not understand the request due to invalid syntax.
	SC_CE_BadRequest StatusCode = 400
	// (HTTP 401) status code indicates that the request has not been applied because it lacks valid authentication credentials for the target resource.
	SC_CE_Unauthorized StatusCode = 401
	// (HTTP 403) status code indicates that the server understood the request but refuses to authorize it.
	SC_CE_Forbidden StatusCode = 403
	// (HTTP 404) status code indicates that the resource requested could not be found on the server.
	SC_CE_NotFound StatusCode = 404
	// (HTTP 405) status code indicates that the method specified in the Request-Line is not allowed for the resource identified by the Request-URI.
	SC_CE_MethodNotAllowed StatusCode = 405
	// (HTTP 406) status code indicates that the server can not find a content type that it supports.
	SC_CE_NotAcceptable StatusCode = 406
	// (HTTP 407) status code indicates that the client must first authenticate itself with the proxy.
	SC_CE_ProxyAuthRequired StatusCode = 407
	// (HTTP 408) status code indicates that the client did not produce a request within the time that the server was prepared to wait.
	SC_CE_RequestTimeout StatusCode = 408
	// (HTTP 409) status code indicates that the request could not be processed because of conflict in the request data.
	SC_CE_Conflict StatusCode = 409
	// (HTTP 410) status code indicates that the requested resource is no longer available and will not be available again.
	SC_CE_Gone StatusCode = 410
	// (HTTP 411) status code indicates that the server refuses to accept the request without a defined Content-Length.
	SC_CE_LengthRequired StatusCode = 411
	// (HTTP 412) status code indicates that the server does not meet one of the preconditions that the requester put on the request.
	SC_CE_PreconditionFailed StatusCode = 412
	// (HTTP 413) status code indicates that the request entity is larger than the server is willing or able to process.
	SC_CE_RequestEntityTooLarge StatusCode = 413
	// (HTTP 414) status code indicates that the URI provided was too long for the server to process.
	SC_CE_RequestURITooLong StatusCode = 414
	// (HTTP 415) status code indicates that the server does not support the media type of the request entity.
	SC_CE_UnsupportedMediaType StatusCode = 415
	// (HTTP 416) status code indicates that the range specified by the Range header was not satisfiable.
	SC_CE_RequestedRangeNotSatisfiable StatusCode = 416
	// (HTTP 417) status code indicates that the expectation given in an Expect request header could not be met by the server.
	SC_CE_ExpectationFailed StatusCode = 417
	// (HTTP 422) status code indicates that the request was well-formed but was unable to be followed due to semantic errors.
	SC_CE_UnprocessableEntity StatusCode = 422
)

// Server Errors
const (
	// (HTTP 500) status code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request as specified in the request.
	SC_SE_InternalServerError StatusCode = 500
	// (HTTP 501) status code indicates that the server does not support the functionality required to fulfill the request.
	SC_SE_NotImplemented StatusCode = 501
	// (HTTP 502) status code indicates that the server was acting as a gateway or proxy and received an invalid response from the upstream server.
	SC_SE_BadGateway StatusCode = 502
	// (HTTP 503) status code indicates that the server is currently unavailable (because it is overloaded or down for maintenance).
	SC_SE_ServiceUnavailable StatusCode = 503
	// (HTTP 504) status code indicates that the server was acting as a gateway or proxy and did not receive a timely response from the upstream server.
	SC_SE_GatewayTimeout StatusCode = 504
	// (HTTP 505) status code indicates that the server does not support the HTTP protocol version that was used in the request.
	SC_SE_HTTPVersionNotSupported StatusCode = 505
	// (HTTP 506) status code indicates that the server did not meet one of the preconditions that the requester put on the request.
	SC_SE_VariantAlsoNegotiates StatusCode = 506
	// (HTTP 507) status code indicates that the server is unable to store the representation needed to complete the request.
	SC_SE_InsufficientStorage StatusCode = 507
	// (HTTP 508) status code indicates that the server detected an infinite loop while processing the request.
	SC_SE_LoopDetected StatusCode = 508
	// (HTTP 510) status code indicates that the server does not meet one of the preconditions that the requester put on the request.
	SC_SE_NotExtended StatusCode = 510
	// (HTTP 511) status code indicates that the client needs to authenticate to gain network access.
	SC_SE_NetworkAuthenticationRequired StatusCode = 511
)
