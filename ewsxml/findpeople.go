package ewsxml

import (
	"encoding/xml"
)

type BasePoint string

func (s BasePoint) String() string { return string(s) }

const (
	BasePoint_Beginning BasePoint = "Beginning"
	BasePoint_End       BasePoint = "End"
)

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findpeople
type FindPeople struct {
	XMLName                xml.Name            `xml:"m:FindPeople"`
	PersonaShape           *PersonaShape       `xml:",omitempty"`
	IndexedPageItemView    IndexedPageItemView `xml:",omitempty"`
	Restriction            *SearchExpression   `xml:"m:Restriction,omitempty"`
	AggregationRestriction *SearchExpression   `xml:"m:AggregationRestriction,omitempty"`
	// SortOrder      *SortOrder      `xml:",omitempty"`
	DistinguishedFolderId *DistinguishedFolderId `xml:"m:ParentFolderId>DistinguishedFolderId,omitempty"`
	FolderId              *FolderId              `xml:"m:ParentFolderId>FolderId,omitempty"`
	QueryString           *string                `xml:"m:QueryString,omitempty"`
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/personashape
type PersonaShape struct {
	XMLName              xml.Name `xml:"m:PersonaShape"`
	BaseShape            BaseShape
	AdditionalProperties *AdditionalProperties `xml:",omitempty"`
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/additionalproperties
type AdditionalProperties struct {
	// ExtendedFieldURI []ExtendedFieldURI
	FieldURI []FieldURI
	// IndexedFieldURI  []IndexedFieldURI
}

func (ap *AdditionalProperties) WithFieldURI(fu ...FieldUri) *AdditionalProperties {
	for _, x := range fu {
		ap.FieldURI = append(ap.FieldURI, FieldURI{FieldURI: x})
	}
	return ap
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/indexedpageitemview
type IndexedPageItemView struct {
	XMLName            xml.Name  `xml:"m:IndexedPageItemView"`
	MaxEntriesReturned int       `xml:",attr,omitempty"`
	Offset             int       `xml:",attr"`
	BasePoint          BasePoint `xml:",attr"`
}
