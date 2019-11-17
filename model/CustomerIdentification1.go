package model

// Identifies a customer identification as the search criteria for the financial institution to do the investigation.
type CustomerIdentification1 struct {

	// Identifies the customer for the investigation.
	Party *PartyIdentification43 `xml:"Pty"`

	// Specifies the authority request related to the identified investigation party.
	AuthorityRequest []*AuthorityInvestigation2 `xml:"AuthrtyReq"`
}

func (c *CustomerIdentification1) AddParty() *PartyIdentification43 {
	c.Party = new(PartyIdentification43)
	return c.Party
}

func (c *CustomerIdentification1) AddAuthorityRequest() *AuthorityInvestigation2 {
	newValue := new(AuthorityInvestigation2)
	c.AuthorityRequest = append(c.AuthorityRequest, newValue)
	return newValue
}
