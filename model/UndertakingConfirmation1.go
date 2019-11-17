package model

// Confirmation information for the issued undertaking.
type UndertakingConfirmation1 struct {

	// Party that adds its undertaking to honour the undertaking or amendment of the undertaking.
	Confirmer *PartyIdentification43 `xml:"Cnfrmr"`

	// Unique and unambiguous identifier assigned by the confirmer to the undertaking.
	ReferenceNumber *Max35Text `xml:"RefNb"`

	// Date and time when the undertaking or amendment of the undertaking was confirmed.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Confirmation of the undertaking or amendment of the confirmed undertaking.
	Confirmation []*Max2000Text `xml:"Conf,omitempty"`
}

func (u *UndertakingConfirmation1) AddConfirmer() *PartyIdentification43 {
	u.Confirmer = new(PartyIdentification43)
	return u.Confirmer
}

func (u *UndertakingConfirmation1) SetReferenceNumber(value string) {
	u.ReferenceNumber = (*Max35Text)(&value)
}

func (u *UndertakingConfirmation1) AddDate() *DateAndDateTimeChoice {
	u.Date = new(DateAndDateTimeChoice)
	return u.Date
}

func (u *UndertakingConfirmation1) AddConfirmation(value string) {
	u.Confirmation = append(u.Confirmation, (*Max2000Text)(&value))
}
