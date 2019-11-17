package model

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification4 struct {

	// Code allocated to a financial institution or non financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICOrBEI *AnyBICIdentifier `xml:"BICOrBEI,omitempty"`

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Other []*GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

func (o *OrganisationIdentification4) SetBICOrBEI(value string) {
	o.BICOrBEI = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification4) AddOther() *GenericOrganisationIdentification1 {
	newValue := new(GenericOrganisationIdentification1)
	o.Other = append(o.Other, newValue)
	return newValue
}
