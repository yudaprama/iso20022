package model

// Point of interaction (POI) performing the transaction.
type PointOfInteraction5 struct {

	// Identification of the POI (Point Of Interaction) for the acquirer or its agent.
	Identification *GenericIdentification32 `xml:"Id"`

	// Common name assigned by the acquirer to the POI (Point Of Interaction) system.
	SystemName *Max70Text `xml:"SysNm,omitempty"`

	// Identifier assigned by the merchant identifying a set of POI (Point Of Interaction) terminals performing some categories of transactions.
	GroupIdentification *Max35Text `xml:"GrpId,omitempty"`

	// Capabilities of the POI (Point Of Interaction) performing the transaction.
	Capabilities *PointOfInteractionCapabilities6 `xml:"Cpblties,omitempty"`

	// Time zone name as defined by IANA (Internet Assigned Numbers Authority) in the time zone data base. America/Chicago or Europe/Paris are examples of time zone names.
	TimeZone *Max70Text `xml:"TmZone,omitempty"`

	// Indicates the type of integration of the POI terminal in the sale environment.
	TerminalIntegration *LocationCategory3Code `xml:"TermnlIntgtn,omitempty"`

	// Data related to a component of the POI (Point Of Interaction) performing the transaction.
	Component []*PointOfInteractionComponent6 `xml:"Cmpnt,omitempty"`
}

func (p *PointOfInteraction5) AddIdentification() *GenericIdentification32 {
	p.Identification = new(GenericIdentification32)
	return p.Identification
}

func (p *PointOfInteraction5) SetSystemName(value string) {
	p.SystemName = (*Max70Text)(&value)
}

func (p *PointOfInteraction5) SetGroupIdentification(value string) {
	p.GroupIdentification = (*Max35Text)(&value)
}

func (p *PointOfInteraction5) AddCapabilities() *PointOfInteractionCapabilities6 {
	p.Capabilities = new(PointOfInteractionCapabilities6)
	return p.Capabilities
}

func (p *PointOfInteraction5) SetTimeZone(value string) {
	p.TimeZone = (*Max70Text)(&value)
}

func (p *PointOfInteraction5) SetTerminalIntegration(value string) {
	p.TerminalIntegration = (*LocationCategory3Code)(&value)
}

func (p *PointOfInteraction5) AddComponent() *PointOfInteractionComponent6 {
	newValue := new(PointOfInteractionComponent6)
	p.Component = append(p.Component, newValue)
	return newValue
}
