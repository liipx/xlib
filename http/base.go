package http

import "io"

func Get(url string, headers ...map[string]string) *Response {
	return DoRequest("get", url, nil, headers...)
}

func Post(url string, params io.Reader, headers ...map[string]string) *Response {
	return DoRequest("post", url, params, headers...)
}

// DoRequest returns a Response & an error if something wrong.
func DoRequest(method, url string, params io.Reader, headers ...map[string]string) *Response {
	return NewRequest(method, url, params, headers...).Do()
}