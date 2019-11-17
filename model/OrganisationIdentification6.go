package model

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification6 struct {

	// Code allocated to organisations by the ISO 9362 Registration Authority, under an international identification scheme, as described in the latest version of the standard ISO 9362 Banking (Banking telecommunication messages, Business Identifier Codes).
	BIC *AnyBICIdentifier `xml:"BIC,omitempty"`

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Other []*GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

func (o *OrganisationIdentification6) SetBIC(value string) {
	o.BIC = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification6) AddOther() *GenericOrganisationIdentification1 {
	newValue := new(GenericOrganisationIdentification1)
	o.Other = append(o.Other, newValue)
	return newValue
}
