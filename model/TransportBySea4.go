package model

// Information related for the transportation of goods by sea.
type TransportBySea4 struct {

	// Identifies the port where the goods are loaded on board the ship.
	PortOfLoading *Max35Text `xml:"PortOfLoadng"`

	// Identifies the port where the goods are discharged.
	PortOfDischarge *Max35Text `xml:"PortOfDschrge"`

	// Name of a vessel.
	VesselName *Max35Text `xml:"VsslNm,omitempty"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	SeaCarrierName *Max35Text `xml:"SeaCrrierNm,omitempty"`
}

func (t *TransportBySea4) SetPortOfLoading(value string) {
	t.PortOfLoading = (*Max35Text)(&value)
}

func (t *TransportBySea4) SetPortOfDischarge(value string) {
	t.PortOfDischarge = (*Max35Text)(&value)
}

func (t *TransportBySea4) SetVesselName(value string) {
	t.VesselName = (*Max35Text)(&value)
}

func (t *TransportBySea4) SetSeaCarrierName(value string) {
	t.SeaCarrierName = (*Max35Text)(&value)
}
