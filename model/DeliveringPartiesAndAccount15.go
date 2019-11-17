package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount15 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification102Choice `xml:"Dpstry"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount126 `xml:"Pty1"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount127 `xml:"Pty2,omitempty"`
}

func (d *DeliveringPartiesAndAccount15) AddDepository() *PartyIdentification102Choice {
	d.Depository = new(PartyIdentification102Choice)
	return d.Depository
}

func (d *DeliveringPartiesAndAccount15) AddParty1() *PartyIdentificationAndAccount126 {
	d.Party1 = new(PartyIdentificationAndAccount126)
	return d.Party1
}

func (d *DeliveringPartiesAndAccount15) AddParty2() *PartyIdentificationAndAccount127 {
	d.Party2 = new(PartyIdentificationAndAccount127)
	return d.Party2
}
