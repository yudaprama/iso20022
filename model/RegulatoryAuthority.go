package model

// Entity requiring the regulatory reporting information.
type RegulatoryAuthority struct {

	// Name of the entity requiring the regulatory reporting information.
	AuthorityName *Max70Text `xml:"AuthrtyNm,omitempty"`

	// Country of the entity requiring the regulatory reporting information.
	AuthorityCountry *CountryCode `xml:"AuthrtyCtry,omitempty"`
}

func (r *RegulatoryAuthority) SetAuthorityName(value string) {
	r.AuthorityName = (*Max70Text)(&value)
}

func (r *RegulatoryAuthority) SetAuthorityCountry(value string) {
	r.AuthorityCountry = (*CountryCode)(&value)
}
