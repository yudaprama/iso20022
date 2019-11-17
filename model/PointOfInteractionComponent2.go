package model

// Data related to a component of the POI.
type PointOfInteractionComponent2 struct {

	// Type of component belonging to a POI Terminal.
	POIComponentType *POIComponentType2Code `xml:"POICmpntTp"`

	// Identification of the software, hardware or system provider of the POI component.
	ManufacturerIdentification *Max35Text `xml:"ManfctrId,omitempty"`

	// Identification of a model of POI component for a given manufacturer.
	Model *Max35Text `xml:"Mdl,omitempty"`

	// Version of component belonging to a given model.
	VersionNumber *Max16Text `xml:"VrsnNb,omitempty"`

	// Serial number of a component.
	SerialNumber *Max35Text `xml:"SrlNb,omitempty"`

	// Unique approval number for a component, delivered by a certification body.
	// Usage: More than one approval number could be present, when assigned by different bodies. The certification body identification must be provided within the approval number (for example at the beginning of the value).
	ApprovalNumber []*Max70Text `xml:"ApprvlNb,omitempty"`
}

func (p *PointOfInteractionComponent2) SetPOIComponentType(value string) {
	p.POIComponentType = (*POIComponentType2Code)(&value)
}

func (p *PointOfInteractionComponent2) SetManufacturerIdentification(value string) {
	p.ManufacturerIdentification = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponent2) SetModel(value string) {
	p.Model = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponent2) SetVersionNumber(value string) {
	p.VersionNumber = (*Max16Text)(&value)
}

func (p *PointOfInteractionComponent2) SetSerialNumber(value string) {
	p.SerialNumber = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponent2) AddApprovalNumber(value string) {
	p.ApprovalNumber = append(p.ApprovalNumber, (*Max70Text)(&value))
}
