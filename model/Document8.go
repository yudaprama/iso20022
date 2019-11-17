package model

// Information about a document.
type Document8 struct {

	// Type of document.
	Type *PresentationDocumentFormat1Choice `xml:"Tp"`

	// Wording for document.
	Wording *Max20000Text `xml:"Wrdg,omitempty"`

	// Details related to an electronic presentation.
	ElectronicDetails []*Presentation3 `xml:"ElctrncDtls,omitempty"`
}

func (d *Document8) AddType() *PresentationDocumentFormat1Choice {
	d.Type = new(PresentationDocumentFormat1Choice)
	return d.Type
}

func (d *Document8) SetWording(value string) {
	d.Wording = (*Max20000Text)(&value)
}

func (d *Document8) AddElectronicDetails() *Presentation3 {
	newValue := new(Presentation3)
	d.ElectronicDetails = append(d.ElectronicDetails, newValue)
	return newValue
}
