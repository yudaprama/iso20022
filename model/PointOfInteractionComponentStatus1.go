package model

// Status of a POI component (Point of Interaction).
type PointOfInteractionComponentStatus1 struct {

	// Current version of the component that might include the release number.
	VersionNumber *Max35Text `xml:"VrsnNb,omitempty"`

	// Current status of the component.
	Status *POIComponentStatus1Code `xml:"Sts,omitempty"`
}

func (p *PointOfInteractionComponentStatus1) SetVersionNumber(value string) {
	p.VersionNumber = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponentStatus1) SetStatus(value string) {
	p.Status = (*POIComponentStatus1Code)(&value)
}
