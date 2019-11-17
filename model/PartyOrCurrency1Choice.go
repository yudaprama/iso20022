package model

// Choice of a depostory or settlement currency.
type PartyOrCurrency1Choice struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification63 `xml:"Dpstry"`

	// Currency for the settlement.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`
}

func (p *PartyOrCurrency1Choice) AddDepository() *PartyIdentification63 {
	p.Depository = new(PartyIdentification63)
	return p.Depository
}

func (p *PartyOrCurrency1Choice) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}
