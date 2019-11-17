package model

// Provides information about pending balance and pending transactions.
type PendingBalance4 struct {

	// Signed quantity of balance.
	Balance *SignedQuantityFormat9 `xml:"Bal"`

	// Overall process covering the trade and settlement transactions of financial instruments.
	PendingTransactions []*SettlementTypeAndIdentification21 `xml:"PdgTxs,omitempty"`
}

func (p *PendingBalance4) AddBalance() *SignedQuantityFormat9 {
	p.Balance = new(SignedQuantityFormat9)
	return p.Balance
}

func (p *PendingBalance4) AddPendingTransactions() *SettlementTypeAndIdentification21 {
	newValue := new(SettlementTypeAndIdentification21)
	p.PendingTransactions = append(p.PendingTransactions, newValue)
	return newValue
}
