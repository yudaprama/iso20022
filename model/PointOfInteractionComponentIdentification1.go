package model

// Identification of a POI (Point of Interaction) component.
type PointOfInteractionComponentIdentification1 struct {

	// Hierarchical identification of a hardware component inside all the hardware component of the POI. It is composed of all item numbers of the upper level components, separated by the '.' character, ended by the item number of the current component.
	ItemNumber *Max35Text `xml:"ItmNb,omitempty"`

	// Identifies the provider of the software, hardware or parameters of the POI component.
	ProviderIdentification *Max35Text `xml:"PrvdrId,omitempty"`

	// Identification of the POI component assigned by its provider.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Serial number identifying an occurrence of an hardware component.
	SerialNumber *Max35Text `xml:"SrlNb,omitempty"`
}

func (p *PointOfInteractionComponentIdentification1) SetItemNumber(value string) {
	p.ItemNumber = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponentIdentification1) SetProviderIdentification(value string) {
	p.ProviderIdentification = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponentIdentification1) SetIdentification(value string) {
	p.Identification = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponentIdentification1) SetSerialNumber(value string) {
	p.SerialNumber = (*Max35Text)(&value)
}
