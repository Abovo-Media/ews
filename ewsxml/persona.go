package ewsxml

import (
	"encoding/xml"
)

type GetPersona struct {
	XMLName   xml.Name  `xml:"m:GetPersona"`
	PersonaId PersonaId `xml:"m:PersonaId"`
}

type Persona struct {
	PersonaId            PersonaId
	DisplayName          string
	GivenName            string
	Surname              string
	Title                string
	Department           string
	Departments          Departments
	EmailAddress         Mailbox
	EmailAddresses       []Mailbox
	RelevanceScore       int
	BusinessPhoneNumbers BusinessPhoneNumbers
	MobilePhones         MobilePhones
	OfficeLocations      OfficeLocations
}

type PersonaId struct {
	Id        string `xml:",attr"`
	ChangeKey string `xml:",attr,omitempty"`
}

type BusinessPhoneNumbers struct {
	PhoneNumberAttributedValue PhoneNumberAttributedValue
}

type MobilePhones struct {
	PhoneNumberAttributedValue PhoneNumberAttributedValue
}

type Value struct {
	Number string
	Type   string
}

type PhoneNumberAttributedValue struct {
	Value Value
}

type OfficeLocations struct {
	StringAttributedValue StringAttributedValue
}

type Departments struct {
	StringAttributedValue StringAttributedValue
}

type StringAttributedValue struct {
	Value string
}
