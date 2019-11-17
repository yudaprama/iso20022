package model

// Choice of format for the block trade information.
type BlockTrade4Choice struct {

	// Block parent or child information expressed as an ISO 20022 code.
	Code *BlockTrade1Code `xml:"Cd"`

	// Block parent or child information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (b *BlockTrade4Choice) SetCode(value string) {
	b.Code = (*BlockTrade1Code)(&value)
}

func (b *BlockTrade4Choice) AddProprietary() *GenericIdentification30 {
	b.Proprietary = new(GenericIdentification30)
	return b.Proprietary
}
