package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount7 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification34Choice `xml:"Dpstry"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount32 `xml:"Pty1"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount32 `xml:"Pty2,omitempty"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`
}

func (d *DeliveringPartiesAndAccount7) AddDepository() *PartyIdentification34Choice {
	d.Depository = new(PartyIdentification34Choice)
	return d.Depository
}

func (d *DeliveringPartiesAndAccount7) AddParty1() *PartyIdentificationAndAccount32 {
	d.Party1 = new(PartyIdentificationAndAccount32)
	return d.Party1
}

func (d *DeliveringPartiesAndAccount7) AddParty2() *PartyIdentificationAndAccount32 {
	d.Party2 = new(PartyIdentificationAndAccount32)
	return d.Party2
}

func (d *DeliveringPartiesAndAccount7) SetSecuritiesSettlementSystem(value string) {
	d.SecuritiesSettlementSystem = (*Max35Text)(&value)
}
