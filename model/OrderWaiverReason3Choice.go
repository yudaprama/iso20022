package model

// Choice of formats for an order waiver reason.
type OrderWaiverReason3Choice struct {

	// Reason for the waiver expressed as a code.
	Code *OrderWaiverReason1Code `xml:"Cd"`

	// Reason for the waiver expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OrderWaiverReason3Choice) SetCode(value string) {
	o.Code = (*OrderWaiverReason1Code)(&value)
}

func (o *OrderWaiverReason3Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
