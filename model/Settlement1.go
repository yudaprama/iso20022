package model

// Provides the settlement details.
type Settlement1 struct {

	// Total amount to be settled.
	SettlementAmount *AmountAndDirection27 `xml:"SttlmAmt"`

	// Place where settlement of the securities takes place.
	Depository *PartyIdentification34Choice `xml:"Dpstry,omitempty"`
}

func (s *Settlement1) AddSettlementAmount() *AmountAndDirection27 {
	s.SettlementAmount = new(AmountAndDirection27)
	return s.SettlementAmount
}

func (s *Settlement1) AddDepository() *PartyIdentification34Choice {
	s.Depository = new(PartyIdentification34Choice)
	return s.Depository
}
