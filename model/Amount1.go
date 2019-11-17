package model

// Margin amount payable by one party to the other party.
type Amount1 struct {

	// Undisputed amount of the margin call request.
	AgreedAmount *ActiveCurrencyAndAmount `xml:"AgrdAmt"`

	// Unique identifier for the margin call request.
	MarginCallRequestIdentification *Max35Text `xml:"MrgnCallReqId"`

	// Provides additional information related to the margin call amount that has been agreed.
	AdditionalInformation *Max210Text `xml:"AddtlInf,omitempty"`
}

func (a *Amount1) SetAgreedAmount(value, currency string) {
	a.AgreedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *Amount1) SetMarginCallRequestIdentification(value string) {
	a.MarginCallRequestIdentification = (*Max35Text)(&value)
}

func (a *Amount1) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max210Text)(&value)
}
