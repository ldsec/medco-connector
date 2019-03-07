// Code generated by go-swagger; DO NOT EDIT.

package picsure2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// QueryStatusURL generates an URL for the query status operation
type QueryStatusURL struct {
	QueryID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *QueryStatusURL) WithBasePath(bp string) *QueryStatusURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *QueryStatusURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *QueryStatusURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/picsure2/{queryId}/status"

	queryID := o.QueryID
	if queryID != "" {
		_path = strings.Replace(_path, "{queryId}", queryID, -1)
	} else {
		return nil, errors.New("queryId is required on QueryStatusURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/medco"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *QueryStatusURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *QueryStatusURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *QueryStatusURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on QueryStatusURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on QueryStatusURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *QueryStatusURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
