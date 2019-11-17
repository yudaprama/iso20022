package model

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification58 struct {

	// Name or number assigned by an entity to enable recognition of that entity.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Type *GenericIdentification40 `xml:"Tp"`
}

func (g *GenericIdentification58) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification58) AddType() *GenericIdentification40 {
	g.Type = new(GenericIdentification40)
	return g.Type
}
