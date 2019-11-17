package model

// Choice of formats for the specification of the reason.
type BlockedReason1Choice struct {

	// Reason expressed as a code.
	Reason *BlockedReason1Code `xml:"Rsn"`

	// Reason expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (b *BlockedReason1Choice) SetReason(value string) {
	b.Reason = (*BlockedReason1Code)(&value)
}

func (b *BlockedReason1Choice) AddProprietary() *GenericIdentification47 {
	b.Proprietary = new(GenericIdentification47)
	return b.Proprietary
}
