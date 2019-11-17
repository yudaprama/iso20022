package model

// Information which describes the organisation.
type Organisation14 struct {

	// Name by which a party is known and which is usually used to identify that party.
	FullLegalName *Max350Text `xml:"FullLglNm,omitempty"`

	// Unique and unambiguous way of identifying an organisation.
	OrganisationIdentification *OrganisationIdentification8 `xml:"OrgId"`
}

func (o *Organisation14) SetFullLegalName(value string) {
	o.FullLegalName = (*Max350Text)(&value)
}

func (o *Organisation14) AddOrganisationIdentification() *OrganisationIdentification8 {
	o.OrganisationIdentification = new(OrganisationIdentification8)
	return o.OrganisationIdentification
}
