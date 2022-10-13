package ewsxml

import (
	"encoding/xml"
)

type nodes []interface{}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/restriction
type SearchExpression struct {
	owner interface{}
	Nodes []interface{}
}

func Expr() *SearchExpression {
	expr := new(SearchExpression)
	initExpr(expr, nil, nil)
	return expr
}

func initExpr(x *SearchExpression, owner interface{}, nodes nodes) {
	x.owner = owner
	if len(nodes) == 0 {
		x.Nodes = make([]interface{}, 0, 2)
	} else {
		x.Nodes = nodes
	}
}

type andExpr struct {
	XMLName xml.Name `xml:"And"`
	SearchExpression
}

func newAndExpr(wrap nodes) *andExpr {
	var expr andExpr
	initExpr(&expr.SearchExpression, &expr, wrap)
	return &expr
}

type orExpr struct {
	XMLName xml.Name `xml:"Or"`
	SearchExpression
}

func newOrExpr(wrap nodes) *orExpr {
	var expr orExpr
	initExpr(&expr.SearchExpression, &expr, wrap)
	return &expr
}

func (expr *SearchExpression) Add(node interface{}) *SearchExpression {
	switch len(expr.Nodes) {
	case 0:
		expr.Nodes = append(expr.Nodes, node)
		return expr

	case 1:
		switch n := expr.Nodes[0].(type) {
		case andExpr:
			// and vervangen bij leeg en wanneer node = Or
			n.Add(node)
		case orExpr:
			// or vervangen bij leeg en wanneer node = And
			n.Add(node)
		default:
			expr.Nodes = []interface{}{
				newAndExpr(append(expr.Nodes, node)),
			}
		}
	}
	return expr
}

type ConstantValue struct {
	Value string `xml:",attr"`
}

type FieldURIOrConstant struct {
	Constant *ConstantValue `xml:",omitempty"`
	// IndexedFieldURI  *IndexedFieldURI  `xml:",omitempty"`
	FieldURI *FieldURI `xml:",omitempty"`
	// ExtendedFieldURI *ExtendedFieldURI `xml:",omitempty"`
}

type ContainsExpression struct {
	XMLName xml.Name `xml:"Contains"`
	// ContainmentMode identifies the boundaries of a search.
	ContainmentMode string `xml:",attr"`
	// ContainmentComparison determines whether the search ignores cases and spaces.
	ContainmentComparison string `xml:",attr"`

	FieldURI FieldURI
	Constant ConstantValue
}

func (expr *SearchExpression) Contains(field FieldUri, val string) *SearchExpression {
	return expr.Add(ContainsExpression{
		ContainmentMode:       "Substring",
		ContainmentComparison: "IgnoreCase",
		FieldURI:              FieldURI{FieldURI: field},
		Constant:              ConstantValue{Value: val},
	})
}

type IsEqualTo struct {
	XMLName xml.Name `xml:"IsEqualTo"`

	FieldURI           FieldURI
	FieldURIOrConstant FieldURIOrConstant
}

func (expr *SearchExpression) Eq(field FieldUri, val interface{}) *SearchExpression {
	eq := IsEqualTo{FieldURI: FieldURI{FieldURI: field}}

	switch v := val.(type) {
	case string:
		eq.FieldURIOrConstant.Constant = &ConstantValue{Value: v}
	case *ConstantValue:
		eq.FieldURIOrConstant.Constant = v
	case ConstantValue:
		eq.FieldURIOrConstant.Constant = &v
	}

	return expr.Add(eq)
}
