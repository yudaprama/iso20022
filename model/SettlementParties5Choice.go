package model

// Provides details on either the delivering or receiving parties.
type SettlementParties5Choice struct {

	// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
	DeliveringSettlementParties *DeliveringPartiesAndAccount15 `xml:"DlvrgSttlmPties,omitempty"`

	// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
	ReceivingSettlementParties *ReceivingPartiesAndAccount15 `xml:"RcvgSttlmPties,omitempty"`
}

func (s *SettlementParties5Choice) AddDeliveringSettlementParties() *DeliveringPartiesAndAccount15 {
	s.DeliveringSettlementParties = new(DeliveringPartiesAndAccount15)
	return s.DeliveringSettlementParties
}

func (s *SettlementParties5Choice) AddReceivingSettlementParties() *ReceivingPartiesAndAccount15 {
	s.ReceivingSettlementParties = new(ReceivingPartiesAndAccount15)
	return s.ReceivingSettlementParties
}
