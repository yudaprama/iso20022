package model

// Assessments for the component of the point of interaction.
type PointOfInteractionComponentAssessment1 struct {

	// Type of assessment of the component.
	Type *POIComponentAssessment1Code `xml:"Tp"`

	// Body which has delivered the assessment.
	Assigner []*Max35Text `xml:"Assgnr"`

	// Date when the assessment has been delivered.
	DeliveryDate *ISODateTime `xml:"DlvryDt,omitempty"`

	// Date when the assessment will expire.
	ExpirationDate *ISODateTime `xml:"XprtnDt,omitempty"`

	// Unique assessment number for the component.
	Number *Max35Text `xml:"Nb"`
}

func (p *PointOfInteractionComponentAssessment1) SetType(value string) {
	p.Type = (*POIComponentAssessment1Code)(&value)
}

func (p *PointOfInteractionComponentAssessment1) AddAssigner(value string) {
	p.Assigner = append(p.Assigner, (*Max35Text)(&value))
}

func (p *PointOfInteractionComponentAssessment1) SetDeliveryDate(value string) {
	p.DeliveryDate = (*ISODateTime)(&value)
}

func (p *PointOfInteractionComponentAssessment1) SetExpirationDate(value string) {
	p.ExpirationDate = (*ISODateTime)(&value)
}

func (p *PointOfInteractionComponentAssessment1) SetNumber(value string) {
	p.Number = (*Max35Text)(&value)
}
