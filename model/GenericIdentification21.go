package model

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification21 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Type *GenericIdentification20 `xml:"Tp"`

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (g *GenericIdentification21) AddType() *GenericIdentification20 {
	g.Type = new(GenericIdentification20)
	return g.Type
}

func (g *GenericIdentification21) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}
