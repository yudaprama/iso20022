package model

// Choice of format for the block trade information.
type BlockTrade3Choice struct {

	// Block parent or child information expressed as an ISO 20022 code.
	Code *BlockTrade1Code `xml:"Cd"`

	// Block parent or child information expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (b *BlockTrade3Choice) SetCode(value string) {
	b.Code = (*BlockTrade1Code)(&value)
}

func (b *BlockTrade3Choice) AddProprietary() *GenericIdentification38 {
	b.Proprietary = new(GenericIdentification38)
	return b.Proprietary
}
