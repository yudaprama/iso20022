package model

// Choice of format for the block trade information.
type BlockTrade5Choice struct {

	// Block parent or child information expressed as an ISO 20022 code.
	Code *BlockTrade1Code `xml:"Cd"`

	// Block parent or child information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (b *BlockTrade5Choice) SetCode(value string) {
	b.Code = (*BlockTrade1Code)(&value)
}

func (b *BlockTrade5Choice) AddProprietary() *GenericIdentification47 {
	b.Proprietary = new(GenericIdentification47)
	return b.Proprietary
}
