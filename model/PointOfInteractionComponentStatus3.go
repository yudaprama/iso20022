package model

// Status of a POI component (Point of Interaction).
type PointOfInteractionComponentStatus3 struct {

	// Current version of the component that might include the release number.
	VersionNumber *Max256Text `xml:"VrsnNb,omitempty"`

	// Current status of the component.
	Status *POIComponentStatus1Code `xml:"Sts,omitempty"`

	// Expiration date of the component.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`
}

func (p *PointOfInteractionComponentStatus3) SetVersionNumber(value string) {
	p.VersionNumber = (*Max256Text)(&value)
}

func (p *PointOfInteractionComponentStatus3) SetStatus(value string) {
	p.Status = (*POIComponentStatus1Code)(&value)
}

func (p *PointOfInteractionComponentStatus3) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISODate)(&value)
}
