package failure

import "net/http"

var (
	internal_server_error   = "internal_server_error"
	too_many_requests_error = "too_many_requests_error"
	bad_request_error       = "bad_request_error"
)

type ErrorWraper func(error) error

var (
	InternalServerError  ErrorWraper = NewFailure(internal_server_error, http.StatusInternalServerError)
	TooManyRequestsError ErrorWraper = NewFailure(too_many_requests_error, http.StatusTooManyRequests)
	BadRequestError      ErrorWraper = NewFailure(bad_request_error, http.StatusBadRequest)
)
