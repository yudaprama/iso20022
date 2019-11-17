package model

// Data related to a component of the POI (Point Of Interaction) performing the transaction.
type PointOfInteractionComponent6 struct {

	// Type of component belonging to a POI (Point Of Interaction) Terminal.
	Type *POIComponentType4Code `xml:"Tp"`

	// Identification of the POI (Point Of Interaction) component.
	Identification *PointOfInteractionComponentIdentification1 `xml:"Id"`

	// Status of the POI (Point Of Interaction) component.
	Status *PointOfInteractionComponentStatus3 `xml:"Sts,omitempty"`

	// Identification of the standard for which the component complies with.
	StandardCompliance []*GenericIdentification48 `xml:"StdCmplc,omitempty"`

	// Characteristics of a POI (Point Of Interaction) component.
	Characteristics *PointOfInteractionComponentCharacteristics2 `xml:"Chrtcs,omitempty"`

	// Assessments for the component of the point of interaction.
	Assessment []*PointOfInteractionComponentAssessment1 `xml:"Assmnt,omitempty"`
}

func (p *PointOfInteractionComponent6) SetType(value string) {
	p.Type = (*POIComponentType4Code)(&value)
}

func (p *PointOfInteractionComponent6) AddIdentification() *PointOfInteractionComponentIdentification1 {
	p.Identification = new(PointOfInteractionComponentIdentification1)
	return p.Identification
}

func (p *PointOfInteractionComponent6) AddStatus() *PointOfInteractionComponentStatus3 {
	p.Status = new(PointOfInteractionComponentStatus3)
	return p.Status
}

func (p *PointOfInteractionComponent6) AddStandardCompliance() *GenericIdentification48 {
	newValue := new(GenericIdentification48)
	p.StandardCompliance = append(p.StandardCompliance, newValue)
	return newValue
}

func (p *PointOfInteractionComponent6) AddCharacteristics() *PointOfInteractionComponentCharacteristics2 {
	p.Characteristics = new(PointOfInteractionComponentCharacteristics2)
	return p.Characteristics
}

func (p *PointOfInteractionComponent6) AddAssessment() *PointOfInteractionComponentAssessment1 {
	newValue := new(PointOfInteractionComponentAssessment1)
	p.Assessment = append(p.Assessment, newValue)
	return newValue
}
