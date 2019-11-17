package model

// Entity requiring the regulatory reporting information.
type RegulatoryAuthority2 struct {

	// Name of the entity requiring the regulatory reporting information.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Country of the entity that requires the regulatory reporting information.
	Country *CountryCode `xml:"Ctry,omitempty"`
}

func (r *RegulatoryAuthority2) SetName(value string) {
	r.Name = (*Max140Text)(&value)
}

func (r *RegulatoryAuthority2) SetCountry(value string) {
	r.Country = (*CountryCode)(&value)
}
