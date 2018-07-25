package postman

import "encoding/json"

// Collection is th root of postman collection file.
// Implementation: https://schema.getpostman.com/json/collection/v2.1.0/docs/index.html
type Collection struct {
	Info      Information `json:"info,omitempty"`
	Items     []Item      `json:"item,omitempty"`
	Events    []Event     `json:"event,omitempty"`
	Variables []Variable  `json:"variable,omitempty"`
	Auth      *Auth       `json:"auth,omitempty"`
}

type Information struct {
	PostmanID   string          `json:"_postman_id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description json.RawMessage `json:"description,omitempty"`
	Version     json.RawMessage `json:"version,omitempty"`
	Schema      string          `json:"schema,omitempty"`
}

type Description struct {
	Content string      `json:"content,omitempty"`
	Type    string      `json:"type,omitempty"`
	Version interface{} `json:"version,omitempty"`
}

type Version struct {
	Major      int         `json:"major,omitempty"`
	Minor      int         `json:"minor,omitempty"`
	Patch      int         `json:"patch,omitempty"`
	Identifier string      `json:"identifier,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
}

type Item struct {
	ID          string          `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description json.RawMessage `json:"description,omitempty"`
	Variables   []Variable      `json:"variable,omitempty"`
	Events      []Event         `json:"event,omitempty"`
	Request     json.RawMessage `json:"request,omitempty"`
	Responses   []Response      `json:"response,omitempty"`
	Items       []Item          `json:"item,omitempty"`
	Auth        *Auth           `json:"auth,omitempty"`
}

type Variable struct {
	ID          string          `json:"id,omitempty"`
	Key         string          `json:"key,omitempty"`
	Value       json.RawMessage `json:"value,omitempty"`
	Type        VariableType    `json:"type,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description json.RawMessage `json:"description,omitempty"`
	System      bool            `json:"system,omitempty"`
	Disabled    bool            `json:"disabled,omitempty"`
}

type VariableType string

const (
	VariableTypeString  VariableType = "string"
	VariableTypeBoolean VariableType = "boolean"
	VariableTypeAny     VariableType = "any"
	VariableTypeNumber  VariableType = "number"
)

type Event struct {
	ID       string      `json:"id,omitempty"`
	Listen   EventListen `json:"listen,omitempty"`
	Script   Script      `json:"script,omitempty"`
	Disabled bool        `json:"disabled,omitempty"`
}

type EventListen string

const (
	EventListenTest       EventListen = "test"
	EventListenPreRequest EventListen = "prerequest"
)

type Script struct {
	ID     string   `json:"id,omitempty"`
	Type   string   `json:"type,omitempty"`
	Exec   []string `json:"exec,omitempty"`
	Source string   `json:"src,omitempty"` // TODO: Add support of URL type.
	Name   string   `json:"name,omitempty"`
}

type Request struct {
	URL         string          `json:"url,omitempty"` // TODO: Add support of URL type.
	Auth        *Auth           `json:"auth,omitempty"`
	Proxy       *ProxyConfig    `json:"proxy,omitempty"`
	Certificate *Certificate    `json:"certificate,omitempty"`
	Method      string          `json:"method,omitempty"`
	Description json.RawMessage `json:"description,omitempty"`
	Headers     json.RawMessage `json:"header,omitempty"`
	Body        *RequestBody    `json:"body,omitempty"`
}

type Auth struct {
	Type   AuthType        `json:"type,omitempty"`
	NoAuth json.RawMessage `json:"noauth,omitempty"`
	AWSV4  []AuthAttribute `json:"awsv4,omitempty"`
	Basic  []AuthAttribute `json:"basic,omitempty"`
	Bearer []AuthAttribute `json:"bearer,omitempty"`
	Digest []AuthAttribute `json:"digest,omitempty"`
	HAWK   []AuthAttribute `json:"hawk,omitempty"`
	NTLM   []AuthAttribute `json:"ntlm,omitempty"`
	OAuth1 []AuthAttribute `json:"oauth1,omitempty"`
	OAuth2 []AuthAttribute `json:"oauth2,omitempty"`
}

type AuthType string

const (
	AuthTypeAWSV4  AuthType = "awsv4"
	AuthTypeBasic  AuthType = "basic"
	AuthTypeBearer AuthType = "bearer"
	AuthTypeDigest AuthType = "digest"
	AuthTypeHAWK   AuthType = "hawk"
	AuthTypeNTLM   AuthType = "ntlm"
	AuthTypeOAuth1 AuthType = "oauth1"
	AuthTypeOAuth2 AuthType = "oauth2"
)

type AuthAttribute struct {
	Key   string          `json:"key,omitempty"`
	Value json.RawMessage `json:"value,omitempty"`
	Type  string          `json:"type,omitempty"`
}

type ProxyConfig struct {
	Match    string `json:"match,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Tunnel   bool   `json:"tunnel,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}

type Certificate struct {
	Name       string            `json:"name,omitempty"`
	Matches    []string          `json:"matches,omitempty"`
	Key        CertificateSource `json:"key,omitempty"`
	Cert       CertificateSource `json:"cert,omitempty"`
	Passphrase string            `json:"passphrase,omitempty"`
}

type CertificateSource struct {
	Source string `json:"src,omitempty"`
}

type RequestHeader struct {
	Key         string      `json:"key,omitempty"`
	Value       string      `json:"value,omitempty"`
	Disabled    string      `json:"disabled,omitempty"`
	Description Description `json:"description,omitempty"`
}

type RequestBody struct {
	Mode       RequestBodyMode       `json:"mode,omitempty"`
	Raw        string                `json:"raw,omitempty"`
	URLEncoded []URLEncodedParameter `json:"urlencoded,omitempty"`
	FormData   []FormParameter       `json:"formdata,omitempty"`
	File       *File                 `json:"file,omitempty"`
}

type RequestBodyMode string

const (
	RequestBodyModeRaw        RequestBodyMode = "raw"
	RequestBodyModeURLEncoded RequestBodyMode = "urlencoded"
	RequestBodyModeFormData   RequestBodyMode = "formdata"
	RequestBodyModeFile       RequestBodyMode = "file"
)

type URLEncodedParameter struct {
	Key         string      `json:"key,omitempty"`
	Value       string      `json:"value,omitempty"`
	Disabled    bool        `json:"disabled,omitempty"`
	Description Description `json:"description,omitempty"`
}

type FormParameter struct {
	Key         string      `json:"key,omitempty"`
	Value       string      `json:"value,omitempty"`
	Source      string      `json:"src,omitempty"`
	Disabled    bool        `json:"disabled,omitempty"`
	Type        string      `json:"type,omitempty"`
	Description Description `json:"description,omitempty"`
}

type File struct {
	Source  string `json:"src,omitempty"`
	Content string `json:"content,omitempty"`
}

type Response struct {
	ID              string          `json:"id,omitempty"`
	OriginalRequest json.RawMessage `json:"originalRequest,omitempty"`
	ResponseTime    json.RawMessage `json:"responseTime,omitempty"`
	Headers         json.RawMessage `json:"header,omitempty"`
	Cookies         []Cookie        `json:"cookie,omitempty"`
	Body            string          `json:"body,omitempty"`
	Status          string          `json:"status,omitempty"`
	Code            int             `json:"code,omitempty"`
}

type Cookie struct {
	Domain     string            `json:"domain,omitempty"`
	Expires    json.RawMessage   `json:"expires,omitempty"`
	MaxAge     string            `json:"maxAge,omitempty"`
	HostOnly   bool              `json:"hostOnly,omitempty"`
	HTTPOnly   bool              `json:"httpOnly,omitempty"`
	Name       string            `json:"name,omitempty"`
	Path       string            `json:"path,omitempty"`
	Secure     bool              `json:"secure,omitempty"`
	Session    bool              `json:"session,omitempty"`
	Value      string            `json:"value,omitempty"`
	Extensions []json.RawMessage `json:"extensions,omitempty"`
}
