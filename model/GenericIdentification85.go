package model

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification85 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Type *GenericIdentification47 `xml:"Tp"`

	// Name or number assigned by an entity to enable recognition of that entity.
	Identification *RestrictedFINXMax30Text `xml:"Id,omitempty"`
}

func (g *GenericIdentification85) AddType() *GenericIdentification47 {
	g.Type = new(GenericIdentification47)
	return g.Type
}

func (g *GenericIdentification85) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax30Text)(&value)
}
