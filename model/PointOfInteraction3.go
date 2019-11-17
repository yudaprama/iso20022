package model

// Point of interaction (POI) performing the transaction.
type PointOfInteraction3 struct {

	// Identification of the POI (Point Of Interaction) for the acquirer or its agent.
	Identification *GenericIdentification32 `xml:"Id"`

	// Common name assigned by the acquirer to the POI (Point Of Interaction) system.
	SystemName *Max70Text `xml:"SysNm,omitempty"`

	// Identifier assigned by the merchant identifying a set ofPOI (Point Of Interaction) terminals performing some categories of transactions.
	GroupIdentification *Max35Text `xml:"GrpId,omitempty"`

	// Capabilities of the POI (Point Of Interaction) performing the transaction.
	Capabilities *PointOfInteractionCapabilities2 `xml:"Cpblties,omitempty"`

	// Data related to a component of the POI (Point Of Interaction) performing the transaction.
	Component []*PointOfInteractionComponent4 `xml:"Cmpnt,omitempty"`
}

func (p *PointOfInteraction3) AddIdentification() *GenericIdentification32 {
	p.Identification = new(GenericIdentification32)
	return p.Identification
}

func (p *PointOfInteraction3) SetSystemName(value string) {
	p.SystemName = (*Max70Text)(&value)
}

func (p *PointOfInteraction3) SetGroupIdentification(value string) {
	p.GroupIdentification = (*Max35Text)(&value)
}

func (p *PointOfInteraction3) AddCapabilities() *PointOfInteractionCapabilities2 {
	p.Capabilities = new(PointOfInteractionCapabilities2)
	return p.Capabilities
}

func (p *PointOfInteraction3) AddComponent() *PointOfInteractionComponent4 {
	newValue := new(PointOfInteractionComponent4)
	p.Component = append(p.Component, newValue)
	return newValue
}
