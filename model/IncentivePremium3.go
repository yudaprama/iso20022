package model

// Cash premium made available if the securities holder consents or participates to an event.
type IncentivePremium3 struct {

	// Description of the premium.
	Description *Max350Text `xml:"Desc,omitempty"`

	// Cash premium paid per security, per vote or per attendee.
	Amount *PriceRateOrAmountChoice `xml:"Amt"`

	// Type of incentive premium.
	Type *IncentivePremiumType1Choice `xml:"Tp"`

	// Date/time for the payment of the premium.
	PaymentDate *DateFormat3Choice `xml:"PmtDt,omitempty"`
}

func (i *IncentivePremium3) SetDescription(value string) {
	i.Description = (*Max350Text)(&value)
}

func (i *IncentivePremium3) AddAmount() *PriceRateOrAmountChoice {
	i.Amount = new(PriceRateOrAmountChoice)
	return i.Amount
}

func (i *IncentivePremium3) AddType() *IncentivePremiumType1Choice {
	i.Type = new(IncentivePremiumType1Choice)
	return i.Type
}

func (i *IncentivePremium3) AddPaymentDate() *DateFormat3Choice {
	i.PaymentDate = new(DateFormat3Choice)
	return i.PaymentDate
}
