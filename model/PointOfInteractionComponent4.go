package model

// Data related to a component of the POI (Point Of Interaction) performing the transaction.
type PointOfInteractionComponent4 struct {

	// Type of component belonging to a POI (Point Of Interaction) Terminal.
	Type *POIComponentType3Code `xml:"Tp"`

	// Identification of the POI (Point Of Interaction) component.
	Identification *PointOfInteractionComponentIdentification1 `xml:"Id"`

	// Status of the POI (Point Of Interaction) component.
	Status *PointOfInteractionComponentStatus1 `xml:"Sts,omitempty"`

	// Identification of the standard for which the component complies with.
	StandardCompliance []*GenericIdentification48 `xml:"StdCmplc,omitempty"`

	// Characteristics of a POI (Point Of Interaction) component.
	Characteristics *PointOfInteractionComponentCharacteristics2 `xml:"Chrtcs,omitempty"`

	// Assessments for the component of the point of interaction.
	Assessment []*PointOfInteractionComponentAssessment1 `xml:"Assmnt,omitempty"`
}

func (p *PointOfInteractionComponent4) SetType(value string) {
	p.Type = (*POIComponentType3Code)(&value)
}

func (p *PointOfInteractionComponent4) AddIdentification() *PointOfInteractionComponentIdentification1 {
	p.Identification = new(PointOfInteractionComponentIdentification1)
	return p.Identification
}

func (p *PointOfInteractionComponent4) AddStatus() *PointOfInteractionComponentStatus1 {
	p.Status = new(PointOfInteractionComponentStatus1)
	return p.Status
}

func (p *PointOfInteractionComponent4) AddStandardCompliance() *GenericIdentification48 {
	newValue := new(GenericIdentification48)
	p.StandardCompliance = append(p.StandardCompliance, newValue)
	return newValue
}

func (p *PointOfInteractionComponent4) AddCharacteristics() *PointOfInteractionComponentCharacteristics2 {
	p.Characteristics = new(PointOfInteractionComponentCharacteristics2)
	return p.Characteristics
}

func (p *PointOfInteractionComponent4) AddAssessment() *PointOfInteractionComponentAssessment1 {
	newValue := new(PointOfInteractionComponentAssessment1)
	p.Assessment = append(p.Assessment, newValue)
	return newValue
}
