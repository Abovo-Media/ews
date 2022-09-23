package ewsxml

import (
	"encoding/xml"
)

type GetPersona struct {
	XMLName   xml.Name  `xml:"m:GetPersona"`
	PersonaId PersonaId `xml:"m:PersonaId"`
}

type GetRooms struct {
	XMLName  xml.Name `xml:"m:GetRooms"`
	RoomList struct {
		EmailAddress string
	} `xml:"m:RoomList"`
}

type GetServerTimeZones struct {
	XMLName                xml.Name `xml:"m:GetServerTimeZones"`
	ReturnFullTimeZoneData bool     `xml:",attr"`
	Ids                    []string `xml:"m:Ids>Id"`
}
