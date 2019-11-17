package model

// Description of the financial instrument.
type FinancialInstrumentDescription3 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification7 `xml:"SctyId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MICIdentifier `xml:"PlcOfListg,omitempty"`

	// Identification of the place of safekeeping.
	SafekeepingPlace *PartyIdentification2Choice `xml:"SfkpgPlc,omitempty"`
}

func (f *FinancialInstrumentDescription3) AddSecurityIdentification() *SecurityIdentification7 {
	f.SecurityIdentification = new(SecurityIdentification7)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentDescription3) SetPlaceOfListing(value string) {
	f.PlaceOfListing = (*MICIdentifier)(&value)
}

func (f *FinancialInstrumentDescription3) AddSafekeepingPlace() *PartyIdentification2Choice {
	f.SafekeepingPlace = new(PartyIdentification2Choice)
	return f.SafekeepingPlace
}
