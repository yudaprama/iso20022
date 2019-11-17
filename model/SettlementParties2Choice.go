package model

// Provides details on either the delivering or receiving parties.
type SettlementParties2Choice struct {

	// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
	DeliveringSettlementParties *DeliveringPartiesAndAccount7 `xml:"DlvrgSttlmPties,omitempty"`

	// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
	ReceivingSettlementParties *ReceivingPartiesAndAccount7 `xml:"RcvgSttlmPties,omitempty"`
}

func (s *SettlementParties2Choice) AddDeliveringSettlementParties() *DeliveringPartiesAndAccount7 {
	s.DeliveringSettlementParties = new(DeliveringPartiesAndAccount7)
	return s.DeliveringSettlementParties
}

func (s *SettlementParties2Choice) AddReceivingSettlementParties() *ReceivingPartiesAndAccount7 {
	s.ReceivingSettlementParties = new(ReceivingPartiesAndAccount7)
	return s.ReceivingSettlementParties
}
