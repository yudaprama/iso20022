package model

// A proprietary or domestic identification scheme that uniquely identifies a financial instrument.
type AlternateFinancialInstrumentIdentification1 struct {

	// Country of the proprietary identification scheme.
	DomesticIdentificationSource *CountryCode `xml:"DmstIdSrc"`

	// Entity that issues the proprietary identification.
	ProprietaryIdentificationSource *Max16Text `xml:"PrtryIdSrc"`

	// Unique and unambiguous identifier of a security.
	Identification *Max35Text `xml:"Id"`
}

func (a *AlternateFinancialInstrumentIdentification1) SetDomesticIdentificationSource(value string) {
	a.DomesticIdentificationSource = (*CountryCode)(&value)
}

func (a *AlternateFinancialInstrumentIdentification1) SetProprietaryIdentificationSource(value string) {
	a.ProprietaryIdentificationSource = (*Max16Text)(&value)
}

func (a *AlternateFinancialInstrumentIdentification1) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}
