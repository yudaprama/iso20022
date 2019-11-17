package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type Account15 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Institution that maintains the records where the account is held.
	Servicer *PartyIdentification2Choice `xml:"Svcr"`
}

func (a *Account15) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Account15) SetDesignation(value string) {
	a.Designation = (*Max35Text)(&value)
}

func (a *Account15) AddServicer() *PartyIdentification2Choice {
	a.Servicer = new(PartyIdentification2Choice)
	return a.Servicer
}
