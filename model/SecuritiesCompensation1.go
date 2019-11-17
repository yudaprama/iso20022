package model

// Provides details about the securities compensation such as the depository and the total settlement amount.
type SecuritiesCompensation1 struct {

	// Place where settlement of the securities takes place.
	Depository *PartyIdentification34Choice `xml:"Dpstry"`

	// Provides the total amount of money to be settled.
	SettlementAmount *AmountAndDirection20 `xml:"SttlmAmt"`

	// Amount of money related to the fees for the securities compensation.
	Fees *AmountAndDirection20 `xml:"Fees,omitempty"`
}

func (s *SecuritiesCompensation1) AddDepository() *PartyIdentification34Choice {
	s.Depository = new(PartyIdentification34Choice)
	return s.Depository
}

func (s *SecuritiesCompensation1) AddSettlementAmount() *AmountAndDirection20 {
	s.SettlementAmount = new(AmountAndDirection20)
	return s.SettlementAmount
}

func (s *SecuritiesCompensation1) AddFees() *AmountAndDirection20 {
	s.Fees = new(AmountAndDirection20)
	return s.Fees
}
