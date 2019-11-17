package model

// Choice of format for the block trade information.
type BlockTrade1Choice struct {

	// Block parent or child information expressed as an ISO 20022 code.
	Code *BlockTrade1Code `xml:"Cd"`

	// Block parent or child information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (b *BlockTrade1Choice) SetCode(value string) {
	b.Code = (*BlockTrade1Code)(&value)
}

func (b *BlockTrade1Choice) AddProprietary() *GenericIdentification20 {
	b.Proprietary = new(GenericIdentification20)
	return b.Proprietary
}
