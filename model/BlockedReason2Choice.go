package model

// Choice of formats for the specification of the reason.
type BlockedReason2Choice struct {

	// Reason expressed as a code.
	Code *BlockedReason2Code `xml:"Cd"`

	// Reason expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (b *BlockedReason2Choice) SetCode(value string) {
	b.Code = (*BlockedReason2Code)(&value)
}

func (b *BlockedReason2Choice) AddProprietary() *GenericIdentification47 {
	b.Proprietary = new(GenericIdentification47)
	return b.Proprietary
}
