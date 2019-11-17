package model

// Investment fund transactions that result in cash movements into a fund, eg, subscription, switch-in.
type InvestmentFundTransactionInType1 struct {

	// Type of transaction, expressed as a code.
	Structured *InvestmentFundTransactionInType3Code `xml:"Strd"`

	// Additional information about the type of transaction.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (i *InvestmentFundTransactionInType1) SetStructured(value string) {
	i.Structured = (*InvestmentFundTransactionInType3Code)(&value)
}

func (i *InvestmentFundTransactionInType1) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max350Text)(&value)
}
