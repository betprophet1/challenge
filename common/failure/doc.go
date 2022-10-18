package failure

import "net/http"

var (
	internal_server_error   = "internal_server_error"
	too_many_requests_error = "too_many_requests_error"
)

type ErrorWraper func(error) error

var (
	InternalServerError  ErrorWraper = NewFailure(internal_server_error, http.StatusInternalServerError)
	TooManyRequestsError ErrorWraper = NewFailure(too_many_requests_error, http.StatusTooManyRequests)
)
