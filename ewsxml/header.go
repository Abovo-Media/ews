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
}

func (h *Header) ServerVersion(ver Version) {
	h.RequestServerVersion.Version = ver
}

func (h *Header) DiscardImpersonation() { h.ExchangeImpersonation = nil }

func (h *Header) ImpersonateSmtpAddress(v string) {
	h.ExchangeImpersonation = &ExchangeImpersonation{
		ConnectingSID: ConnectingSID{SmtpAddress: v},
	}
}

func (h *Header) ImpersonatePrimarySmtpAddress(v string) {
	h.ExchangeImpersonation = &ExchangeImpersonation{
		ConnectingSID: ConnectingSID{PrimarySmtpAddress: v},
	}
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
