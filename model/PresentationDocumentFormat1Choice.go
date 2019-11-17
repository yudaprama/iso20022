package model

// Choice of format for the type of document format.
type PresentationDocumentFormat1Choice struct {

	// Presentation document.
	//
	Code *ExternalUndertakingDocumentType1Code `xml:"Cd"`

	// Document format expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (p *PresentationDocumentFormat1Choice) SetCode(value string) {
	p.Code = (*ExternalUndertakingDocumentType1Code)(&value)
}

func (p *PresentationDocumentFormat1Choice) AddProprietary() *GenericIdentification1 {
	p.Proprietary = new(GenericIdentification1)
	return p.Proprietary
}
