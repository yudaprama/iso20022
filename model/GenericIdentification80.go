package model

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification80 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Type *GenericIdentification30 `xml:"Tp"`

	// Name or number assigned by an entity to enable recognition of that entity.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (g *GenericIdentification80) AddType() *GenericIdentification30 {
	g.Type = new(GenericIdentification30)
	return g.Type
}

func (g *GenericIdentification80) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}
