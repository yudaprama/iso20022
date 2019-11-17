package model

// Choice between source of identification of a financial instrument.
type IdentificationSource5Choice struct {

	// Country of the proprietary identification scheme.
	DomesticIdentificationSource *CountryCode `xml:"DmstIdSrc"`

	// Entity that issues the proprietary identification.
	ProprietaryIdentificationSource *Max35Text `xml:"PrtryIdSrc"`
}

func (i *IdentificationSource5Choice) SetDomesticIdentificationSource(value string) {
	i.DomesticIdentificationSource = (*CountryCode)(&value)
}

func (i *IdentificationSource5Choice) SetProprietaryIdentificationSource(value string) {
	i.ProprietaryIdentificationSource = (*Max35Text)(&value)
}
