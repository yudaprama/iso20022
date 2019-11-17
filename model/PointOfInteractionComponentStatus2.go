package model

// Status of a POI component (Point of Interaction).
type PointOfInteractionComponentStatus2 struct {

	// Current version of the component that might include the release number.
	VersionNumber *Max256Text `xml:"VrsnNb,omitempty"`

	// Current status of the component.
	Status *POIComponentStatus1Code `xml:"Sts,omitempty"`
}

func (p *PointOfInteractionComponentStatus2) SetVersionNumber(value string) {
	p.VersionNumber = (*Max256Text)(&value)
}

func (p *PointOfInteractionComponentStatus2) SetStatus(value string) {
	p.Status = (*POIComponentStatus1Code)(&value)
}
