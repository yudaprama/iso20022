package model

// Defines and associates identifications for a party as a list of other global or qualified relative identifiers.
// It is assumed that customers of a party can be referenced by an identifier local to the party. The party together with the local identifier can be used to reference the customer.
// Multiple references can be given to identify the same party.
// A short identification can be used for display purposes.
type QualifiedPartyIdentification1 struct {

	// Schema ID to be used in IDREF values.
	Identification *ID `xml:"Id"`

	// List of identifications for the same party.
	Party []*SingleQualifiedPartyIdentification1 `xml:"Pty"`

	// Short identification of the resulting party as a control mechanism for humans.
	ShortIdentification *PartyIdentification2Choice `xml:"ShrtId,omitempty"`

	// Formally defined role qualifying the party.
	Role *GenericIdentification1 `xml:"Role,omitempty"`

	// Free form description of the party's role.
	RoleDescription *Max256Text `xml:"RoleDesc,omitempty"`
}

func (q *QualifiedPartyIdentification1) SetIdentification(value string) {
	q.Identification = (*ID)(&value)
}

func (q *QualifiedPartyIdentification1) AddParty() *SingleQualifiedPartyIdentification1 {
	newValue := new(SingleQualifiedPartyIdentification1)
	q.Party = append(q.Party, newValue)
	return newValue
}

func (q *QualifiedPartyIdentification1) AddShortIdentification() *PartyIdentification2Choice {
	q.ShortIdentification = new(PartyIdentification2Choice)
	return q.ShortIdentification
}

func (q *QualifiedPartyIdentification1) AddRole() *GenericIdentification1 {
	q.Role = new(GenericIdentification1)
	return q.Role
}

func (q *QualifiedPartyIdentification1) SetRoleDescription(value string) {
	q.RoleDescription = (*Max256Text)(&value)
}
