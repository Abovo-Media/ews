package ewsxml

import (
	"encoding/xml"
)

type Version string

func (v Version) String() string { return string(v) }

type Header struct {
	XMLName               xml.Name `xml:"soap:Header"`
	RequestServerVersion  RequestServerVersion
	ExchangeImpersonation *ExchangeImpersonation `xml:",omitempty"`
	TimeZoneContext       *TimeZoneContext       `xml:",omitempty"`
	DateTimePrecision     DateTimePrecision      `xml:",omitempty"`
}

func (h *Header) ServerVersion() Version { return h.RequestServerVersion.Version }

func (h *Header) WithServerVersion(ver Version) *Header {
	h.RequestServerVersion.Version = ver
	return h
}

func (h *Header) DiscardImpersonation() *Header {
	h.ExchangeImpersonation = nil
	return h
}

func (h *Header) WithImpersonateSmtpAddress(v string) *Header {
	if h.ExchangeImpersonation == nil {
		h.ExchangeImpersonation = new(ExchangeImpersonation)
	}
	h.ExchangeImpersonation.ConnectingSID.SmtpAddress = v
	return h
}

func (h *Header) WithImpersonatePrimarySmtpAddress(v string) *Header {
	if h.ExchangeImpersonation == nil {
		h.ExchangeImpersonation = new(ExchangeImpersonation)
	}
	h.ExchangeImpersonation.ConnectingSID.PrimarySmtpAddress = v
	return h
}

func (h *Header) DiscardTimeZone() *Header {
	h.TimeZoneContext = nil
	return h
}

func (h *Header) WithTimeZoneId(id string) *Header {
	if id == "" {
		return h
	}
	if h.TimeZoneContext == nil {
		h.TimeZoneContext = new(TimeZoneContext)
	}
	h.TimeZoneContext.TimeZoneDefinition.Id = id
	return h
}

type RequestServerVersion struct {
	Version Version `xml:",attr"`
}

type ExchangeImpersonation struct {
	ConnectingSID ConnectingSID `xml:",omitempty"`
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/connectingsid
type ConnectingSID struct {
	PrincipalName      string `xml:",omitempty"`
	SID                string `xml:",omitempty"`
	SmtpAddress        string `xml:",omitempty"`
	PrimarySmtpAddress string `xml:",omitempty"`
}
