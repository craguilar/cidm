// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// TokenValidatPayloadURL generates an URL for the token validat payload operation
type TokenValidatPayloadURL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *TokenValidatPayloadURL) WithBasePath(bp string) *TokenValidatPayloadURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *TokenValidatPayloadURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *TokenValidatPayloadURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/token/authentication"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/idm"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *TokenValidatPayloadURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *TokenValidatPayloadURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *TokenValidatPayloadURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on TokenValidatPayloadURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on TokenValidatPayloadURL")
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
func (o *TokenValidatPayloadURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
