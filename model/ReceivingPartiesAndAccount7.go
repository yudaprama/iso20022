package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount7 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification34Choice `xml:"Dpstry"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount32 `xml:"Pty1"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount32 `xml:"Pty2,omitempty"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`
}

func (r *ReceivingPartiesAndAccount7) AddDepository() *PartyIdentification34Choice {
	r.Depository = new(PartyIdentification34Choice)
	return r.Depository
}

func (r *ReceivingPartiesAndAccount7) AddParty1() *PartyIdentificationAndAccount32 {
	r.Party1 = new(PartyIdentificationAndAccount32)
	return r.Party1
}

func (r *ReceivingPartiesAndAccount7) AddParty2() *PartyIdentificationAndAccount32 {
	r.Party2 = new(PartyIdentificationAndAccount32)
	return r.Party2
}

func (r *ReceivingPartiesAndAccount7) SetSecuritiesSettlementSystem(value string) {
	r.SecuritiesSettlementSystem = (*Max35Text)(&value)
}
