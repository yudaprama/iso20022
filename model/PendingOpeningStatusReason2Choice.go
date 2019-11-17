package model

// Choice of formats for pending account opening reason code.
type PendingOpeningStatusReason2Choice struct {

	// Reason for the pending account opening status expressed as a code.
	Code *PendingOpeningStatusReason1Code `xml:"Cd"`

	// Reason for the pending account opening status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (p *PendingOpeningStatusReason2Choice) SetCode(value string) {
	p.Code = (*PendingOpeningStatusReason1Code)(&value)
}

func (p *PendingOpeningStatusReason2Choice) AddProprietary() *GenericIdentification36 {
	p.Proprietary = new(GenericIdentification36)
	return p.Proprietary
}
