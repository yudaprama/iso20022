package model

// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
type PartyIdentification15 struct {

	// Country in which the organisation is registered.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Description of an organisation.
	Identification *Max35Text `xml:"Id"`
}

func (p *PartyIdentification15) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

func (p *PartyIdentification15) SetIdentification(value string) {
	p.Identification = (*Max35Text)(&value)
}
