package model

// Choice of formats for pending reason code.
type PendingStatusReason2Choice struct {

	// Reason for the pending account status expressed as a code.
	Code *PendingStatusReason1Code `xml:"Cd"`

	// Reason for the pending account status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (p *PendingStatusReason2Choice) SetCode(value string) {
	p.Code = (*PendingStatusReason1Code)(&value)
}

func (p *PendingStatusReason2Choice) AddProprietary() *GenericIdentification36 {
	p.Proprietary = new(GenericIdentification36)
	return p.Proprietary
}
