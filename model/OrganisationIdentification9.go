package model

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification9 struct {

	// Code allocated to an institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC,omitempty"`

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Other []*GenericOrganisationIdentification2 `xml:"Othr,omitempty"`
}

func (o *OrganisationIdentification9) SetAnyBIC(value string) {
	o.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification9) AddOther() *GenericOrganisationIdentification2 {
	newValue := new(GenericOrganisationIdentification2)
	o.Other = append(o.Other, newValue)
	return newValue
}
