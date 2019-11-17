package model

// Provides details on either the delivering or receiving parties.
type SettlementParties3Choice struct {

	// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
	DeliveringSettlementParties *DeliveringPartiesAndAccount10 `xml:"DlvrgSttlmPties,omitempty"`

	// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
	ReceivingSettlementParties *ReceivingPartiesAndAccount10 `xml:"RcvgSttlmPties,omitempty"`
}

func (s *SettlementParties3Choice) AddDeliveringSettlementParties() *DeliveringPartiesAndAccount10 {
	s.DeliveringSettlementParties = new(DeliveringPartiesAndAccount10)
	return s.DeliveringSettlementParties
}

func (s *SettlementParties3Choice) AddReceivingSettlementParties() *ReceivingPartiesAndAccount10 {
	s.ReceivingSettlementParties = new(ReceivingPartiesAndAccount10)
	return s.ReceivingSettlementParties
}
