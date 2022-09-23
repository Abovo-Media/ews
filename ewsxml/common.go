package ewsxml

import (
	"strings"
)

type BasePoint string

func (s BasePoint) String() string { return string(s) }

const (
	BasePointBeginning BasePoint = "Beginning"
	BasePointEnd       BasePoint = "End"
)

// ConcatenatedString represents the concatenated display string that is used
// for the contents of the element. Each part represents its own value.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaycc
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displayto
type ConcatenatedString string

func (c ConcatenatedString) Split(sep string) []string {
	return strings.Split(string(c), sep)
}

func (s ConcatenatedString) String() string { return string(s) }
