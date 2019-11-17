package model

// Defines the type of the billing balance.
type BillingBalanceType1Choice struct {

	// Defines the type of billing balance, as published in an external billing balance type code list.
	Code *ExternalBillingBalanceType1Code `xml:"Cd"`

	// Defines the type of billing balance, in a proprietary format.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (b *BillingBalanceType1Choice) SetCode(value string) {
	b.Code = (*ExternalBillingBalanceType1Code)(&value)
}

func (b *BillingBalanceType1Choice) SetProprietary(value string) {
	b.Proprietary = (*Max35Text)(&value)
}
