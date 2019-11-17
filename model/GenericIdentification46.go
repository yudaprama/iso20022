package model

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification46 struct {

	// Name or number assigned by an entity to enable recognition of that entity, for example, national registration identification number.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identifier.
	Type *OtherIdentification1Choice `xml:"Tp"`
}

func (g *GenericIdentification46) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification46) AddType() *OtherIdentification1Choice {
	g.Type = new(OtherIdentification1Choice)
	return g.Type
}
