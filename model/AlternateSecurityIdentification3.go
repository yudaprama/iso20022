package model

// Proprietary or domestic identification scheme that uniquely identifies a security.
type AlternateSecurityIdentification3 struct {

	// Unique and unambiguous identifier of a security.
	Identification *Max70Text `xml:"Id"`

	// Country of the proprietary identification scheme.
	DomesticIdentificationSource *CountryCode `xml:"DmstIdSrc"`

	// Entity that issues the proprietary identification.
	ProprietaryIdentificationSource *Max35Text `xml:"PrtryIdSrc"`
}

func (a *AlternateSecurityIdentification3) SetIdentification(value string) {
	a.Identification = (*Max70Text)(&value)
}

func (a *AlternateSecurityIdentification3) SetDomesticIdentificationSource(value string) {
	a.DomesticIdentificationSource = (*CountryCode)(&value)
}

func (a *AlternateSecurityIdentification3) SetProprietaryIdentificationSource(value string) {
	a.ProprietaryIdentificationSource = (*Max35Text)(&value)
}
