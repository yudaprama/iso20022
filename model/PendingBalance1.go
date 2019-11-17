package model

// Provides information about pending balance and pending transactions.
type PendingBalance1 struct {

	// Signed quantity of balance.
	Balance *SignedQuantityFormat2 `xml:"Bal"`

	// Overall process covering the trade and settlement transactions of financial instruments.
	PendingTransactions []*SettlementTypeAndIdentification2 `xml:"PdgTxs,omitempty"`
}

func (p *PendingBalance1) AddBalance() *SignedQuantityFormat2 {
	p.Balance = new(SignedQuantityFormat2)
	return p.Balance
}

func (p *PendingBalance1) AddPendingTransactions() *SettlementTypeAndIdentification2 {
	newValue := new(SettlementTypeAndIdentification2)
	p.PendingTransactions = append(p.PendingTransactions, newValue)
	return newValue
}
