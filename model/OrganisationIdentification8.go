package model

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification8 struct {

	// Code allocated to a financial institution or non financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC,omitempty"`

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Other []*GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

func (o *OrganisationIdentification8) SetAnyBIC(value string) {
	o.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification8) AddOther() *GenericOrganisationIdentification1 {
	newValue := new(GenericOrganisationIdentification1)
	o.Other = append(o.Other, newValue)
	return newValue
}
