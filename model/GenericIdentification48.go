package model

// Information related to an identification.
type GenericIdentification48 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Max35Text `xml:"Id"`

	// Version of the identification.
	Version *Max35Text `xml:"Vrsn"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`
}

func (g *GenericIdentification48) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification48) SetVersion(value string) {
	g.Version = (*Max35Text)(&value)
}

func (g *GenericIdentification48) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
