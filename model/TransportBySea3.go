package model

// Information related for the transportation of goods by sea.
type TransportBySea3 struct {

	// Identifies the port where the goods are loaded on board the ship.
	PortOfLoading []*Max35Text `xml:"PortOfLoadng,omitempty"`

	// Identifies the port where the goods are discharged.
	PortOfDischarge []*Max35Text `xml:"PortOfDschrge"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	SeaCarrierName *Max35Text `xml:"SeaCrrierNm,omitempty"`
}

func (t *TransportBySea3) AddPortOfLoading(value string) {
	t.PortOfLoading = append(t.PortOfLoading, (*Max35Text)(&value))
}

func (t *TransportBySea3) AddPortOfDischarge(value string) {
	t.PortOfDischarge = append(t.PortOfDischarge, (*Max35Text)(&value))
}

func (t *TransportBySea3) SetSeaCarrierName(value string) {
	t.SeaCarrierName = (*Max35Text)(&value)
}
