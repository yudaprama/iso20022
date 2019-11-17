package model

// Provides information about pending balance and pending transactions.
type PendingBalance3 struct {

	// Signed quantity of balance.
	Balance *SignedQuantityFormat6 `xml:"Bal"`

	// Overall process covering the trade and settlement transactions of financial instruments.
	PendingTransactions []*SettlementTypeAndIdentification20 `xml:"PdgTxs,omitempty"`
}

func (p *PendingBalance3) AddBalance() *SignedQuantityFormat6 {
	p.Balance = new(SignedQuantityFormat6)
	return p.Balance
}

func (p *PendingBalance3) AddPendingTransactions() *SettlementTypeAndIdentification20 {
	newValue := new(SettlementTypeAndIdentification20)
	p.PendingTransactions = append(p.PendingTransactions, newValue)
	return newValue
}
