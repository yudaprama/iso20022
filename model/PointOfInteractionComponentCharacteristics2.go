package model

// Physical and logical characteristics of a POI component (Point of Interaction).
type PointOfInteractionComponentCharacteristics2 struct {

	// Memory characteristics of the component.
	Memory []*MemoryCharacteristics1 `xml:"Mmry,omitempty"`

	// Low level communication of the hardware or software component toward another component or an external entity.
	Communication []*CommunicationCharacteristics2 `xml:"Com,omitempty"`

	// Number of security access modules (SAM).
	SecurityAccessModules *Number `xml:"SctyAccsMdls,omitempty"`

	// Number of subscriber identity modules (SIM).
	SubscriberIdentityModules *Number `xml:"SbcbrIdntyMdls,omitempty"`

	// Value for checking a cryptographic key security parameter.
	KeyCheckValue *Max35Binary `xml:"KeyChckVal,omitempty"`
}

func (p *PointOfInteractionComponentCharacteristics2) AddMemory() *MemoryCharacteristics1 {
	newValue := new(MemoryCharacteristics1)
	p.Memory = append(p.Memory, newValue)
	return newValue
}

func (p *PointOfInteractionComponentCharacteristics2) AddCommunication() *CommunicationCharacteristics2 {
	newValue := new(CommunicationCharacteristics2)
	p.Communication = append(p.Communication, newValue)
	return newValue
}

func (p *PointOfInteractionComponentCharacteristics2) SetSecurityAccessModules(value string) {
	p.SecurityAccessModules = (*Number)(&value)
}

func (p *PointOfInteractionComponentCharacteristics2) SetSubscriberIdentityModules(value string) {
	p.SubscriberIdentityModules = (*Number)(&value)
}

func (p *PointOfInteractionComponentCharacteristics2) SetKeyCheckValue(value string) {
	p.KeyCheckValue = (*Max35Binary)(&value)
}
