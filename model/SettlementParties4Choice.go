package model

// Provides details on either the delivering or receiving parties.
type SettlementParties4Choice struct {

	// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
	DeliveringSettlementParties *DeliveringPartiesAndAccount11 `xml:"DlvrgSttlmPties,omitempty"`

	// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
	ReceivingSettlementParties *ReceivingPartiesAndAccount11 `xml:"RcvgSttlmPties,omitempty"`
}

func (s *SettlementParties4Choice) AddDeliveringSettlementParties() *DeliveringPartiesAndAccount11 {
	s.DeliveringSettlementParties = new(DeliveringPartiesAndAccount11)
	return s.DeliveringSettlementParties
}

func (s *SettlementParties4Choice) AddReceivingSettlementParties() *ReceivingPartiesAndAccount11 {
	s.ReceivingSettlementParties = new(ReceivingPartiesAndAccount11)
	return s.ReceivingSettlementParties
}
