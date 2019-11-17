package model

// Identification of a point of interaction.
type PointOfInteraction6 struct {

	// Identifier of the terminal manufacturer.
	ManufacturerIdentifier *Max35Text `xml:"ManfctrIdr"`

	// Identifier of the terminal model.
	Model *Max35Text `xml:"Mdl"`

	// Serial number of the terminal manufacturer.
	SerialNumber *Max35Text `xml:"SrlNb"`
}

func (p *PointOfInteraction6) SetManufacturerIdentifier(value string) {
	p.ManufacturerIdentifier = (*Max35Text)(&value)
}

func (p *PointOfInteraction6) SetModel(value string) {
	p.Model = (*Max35Text)(&value)
}

func (p *PointOfInteraction6) SetSerialNumber(value string) {
	p.SerialNumber = (*Max35Text)(&value)
}
