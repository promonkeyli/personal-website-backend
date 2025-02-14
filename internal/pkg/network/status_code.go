package network

type StatusCode int

const (
	// Successful responses
	// swagger:enum StatusCode
	StatusOK StatusCode = 200

	// Client error responses
	StatusBadRequest   StatusCode = 400
	StatusUnauthorized StatusCode = 401
	StatusForbidden    StatusCode = 403
	StatusNotFound     StatusCode = 404

	// Server error responses
	StatusInternalServerError StatusCode = 500
	StatusNotImplemented      StatusCode = 501
	StatusBadGateway          StatusCode = 502
	StatusServiceUnavailable  StatusCode = 503
)

// 实现 StatusCode 类型的string方法，会返回状态码对应的描述
func (s StatusCode) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusBadRequest:
		return "Bad Request"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusForbidden:
		return "Forbidden"
	case StatusNotFound:
		return "Not Found"
	case StatusInternalServerError:
		return "Internal Server Error"
	case StatusNotImplemented:
		return "Not Implemented"
	case StatusBadGateway:
		return "Bad Gateway"
	case StatusServiceUnavailable:
		return "Service Unavailable"
	default:
		return "Unknown Status"
	}
}
