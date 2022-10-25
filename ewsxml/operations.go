package ewsxml

import (
	"encoding/xml"
)

type GetRooms struct {
	XMLName  xml.Name `xml:"m:GetRooms"`
	RoomList struct {
		EmailAddress string
	} `xml:"m:RoomList"`
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getservertimezones
type GetServerTimeZones struct {
	XMLName                xml.Name  `xml:"m:GetServerTimeZones"`
	ReturnFullTimeZoneData bool      `xml:",attr,omitempty"`
	Ids                    *[]string `xml:"m:Ids>Id,omitempty"`
}
