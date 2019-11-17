package model

// Information about a document.
type Document11 struct {

	// Type of document.
	Type *PresentationDocumentFormat1Choice `xml:"Tp,omitempty"`

	// Wording for document.
	Wording *Max20000Text `xml:"Wrdg,omitempty"`

	// Details related to an electronic presentation.
	ElectronicDetails []*Presentation3 `xml:"ElctrncDtls,omitempty"`
}

func (d *Document11) AddType() *PresentationDocumentFormat1Choice {
	d.Type = new(PresentationDocumentFormat1Choice)
	return d.Type
}

func (d *Document11) SetWording(value string) {
	d.Wording = (*Max20000Text)(&value)
}

func (d *Document11) AddElectronicDetails() *Presentation3 {
	newValue := new(Presentation3)
	d.ElectronicDetails = append(d.ElectronicDetails, newValue)
	return newValue
}
