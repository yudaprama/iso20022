package model

// Special conditions for the loan.
type SpecialCondition1 struct {

	// Incoming amount on special conditions.
	IncomingAmount *ActiveCurrencyAndAmount `xml:"IncmgAmt"`

	// Outgoing amount on special conditions.
	OutgoingAmount *ActiveCurrencyAndAmount `xml:"OutgngAmt"`

	// Incoming amount to other account on special conditions.
	IncomingAmountToOtherAccount *ActiveCurrencyAndAmount `xml:"IncmgAmtToOthrAcct,omitempty"`

	// Outgoing payment amount from other account on special conditions.
	PaymentFromOtherAccount *ActiveCurrencyAndAmount `xml:"PmtFrOthrAcct,omitempty"`
}

func (s *SpecialCondition1) SetIncomingAmount(value, currency string) {
	s.IncomingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SpecialCondition1) SetOutgoingAmount(value, currency string) {
	s.OutgoingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SpecialCondition1) SetIncomingAmountToOtherAccount(value, currency string) {
	s.IncomingAmountToOtherAccount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SpecialCondition1) SetPaymentFromOtherAccount(value, currency string) {
	s.PaymentFromOtherAccount = NewActiveCurrencyAndAmount(value, currency)
}
