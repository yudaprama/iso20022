package model

// Proprietary or domestic identification scheme that uniquely identifies a security.
type AlternateSecurityIdentification1 struct {

	// Unique and unambiguous identifier of a security.
	Identification *Max35Text `xml:"Id"`

	// Country of the proprietary identification scheme.
	DomesticIdentificationSource *CountryCode `xml:"DmstIdSrc"`

	// Entity that issues the proprietary identification.
	ProprietaryIdentificationSource *Max35Text `xml:"PrtryIdSrc"`
}

func (a *AlternateSecurityIdentification1) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AlternateSecurityIdentification1) SetDomesticIdentificationSource(value string) {
	a.DomesticIdentificationSource = (*CountryCode)(&value)
}

func (a *AlternateSecurityIdentification1) SetProprietaryIdentificationSource(value string) {
	a.ProprietaryIdentificationSource = (*Max35Text)(&value)
}
