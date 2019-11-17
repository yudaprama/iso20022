package model

// Choice of format for the commercial contract type.
type UnderlyingTradeTransactionType1Choice struct {

	// Type of commercial contract.
	//
	Code *ExternalUnderlyingTradeTransactionType1Code `xml:"Cd"`

	// Type of commercial contract expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (u *UnderlyingTradeTransactionType1Choice) SetCode(value string) {
	u.Code = (*ExternalUnderlyingTradeTransactionType1Code)(&value)
}

func (u *UnderlyingTradeTransactionType1Choice) AddProprietary() *GenericIdentification1 {
	u.Proprietary = new(GenericIdentification1)
	return u.Proprietary
}
