package model

// Pre-agreed offering between clearing agents, or the channel through which the payment instruction is to be processed. This payment scheme can point to a specific rulebook governing the rules of clearing and settlement between two parties.
type PaymentSchemeChoice struct {

	// Specifies the channel through which the payment instruction is to be processed in coded form.
	Code *CashClearingSystem2Code `xml:"Cd"`

	// Channel that is specific to a user community and is required for use within that user community.
	//
	// Usage : if the channel is included in the list of codes provided for the payment scheme, the code element should be used instead of the proprietary element.
	ProprietaryInformation *Max35Text `xml:"PrtryInf"`
}

func (p *PaymentSchemeChoice) SetCode(value string) {
	p.Code = (*CashClearingSystem2Code)(&value)
}

func (p *PaymentSchemeChoice) SetProprietaryInformation(value string) {
	p.ProprietaryInformation = (*Max35Text)(&value)
}
