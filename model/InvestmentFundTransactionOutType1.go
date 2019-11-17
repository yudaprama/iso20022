package model

// Investment fund transactions that result in cash movements out of a fund, eg, redemption, switch-out.
type InvestmentFundTransactionOutType1 struct {

	// Type of transaction, expressed as a code.
	Structured *InvestmentFundTransactionOutType4Code `xml:"Strd"`

	// Additional information about the type of transaction.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (i *InvestmentFundTransactionOutType1) SetStructured(value string) {
	i.Structured = (*InvestmentFundTransactionOutType4Code)(&value)
}

func (i *InvestmentFundTransactionOutType1) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max350Text)(&value)
}
