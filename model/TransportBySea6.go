package model

// Information related for the transportation of goods by sea.
type TransportBySea6 struct {

	// Identifies the port where the goods are loaded on board the ship.
	PortOfLoading []*Max35Text `xml:"PortOfLoadng,omitempty"`

	// Identifies the port where the goods are discharged.
	PortOfDischarge []*Max35Text `xml:"PortOfDschrge"`

	// Name of a vessel.
	VesselName *Max70Text `xml:"VsslNm,omitempty"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	SeaCarrierName *Max70Text `xml:"SeaCrrierNm,omitempty"`

	// Country in which the carrier of the goods, for example, shipping company, is located or registered.
	SeaCarrierCountry *CountryCode `xml:"SeaCrrierCtry,omitempty"`

	// Name of the carrier's (for example, shipping company's) agent that acts on behalf of the carrier and may be the issuer of transport documents relating to the underlying shipment.
	CarrierAgentName *Max70Text `xml:"CrrierAgtNm,omitempty"`

	// Country of registration of the carrier's (for example, shipping company's) agent that acts on behalf of the carrier and may be the issuer of transport documents relating to the underlying shipment.
	CarrierAgentCountry *CountryCode `xml:"CrrierAgtCtry,omitempty"`
}

func (t *TransportBySea6) AddPortOfLoading(value string) {
	t.PortOfLoading = append(t.PortOfLoading, (*Max35Text)(&value))
}

func (t *TransportBySea6) AddPortOfDischarge(value string) {
	t.PortOfDischarge = append(t.PortOfDischarge, (*Max35Text)(&value))
}

func (t *TransportBySea6) SetVesselName(value string) {
	t.VesselName = (*Max70Text)(&value)
}

func (t *TransportBySea6) SetSeaCarrierName(value string) {
	t.SeaCarrierName = (*Max70Text)(&value)
}

func (t *TransportBySea6) SetSeaCarrierCountry(value string) {
	t.SeaCarrierCountry = (*CountryCode)(&value)
}

func (t *TransportBySea6) SetCarrierAgentName(value string) {
	t.CarrierAgentName = (*Max70Text)(&value)
}

func (t *TransportBySea6) SetCarrierAgentCountry(value string) {
	t.CarrierAgentCountry = (*CountryCode)(&value)
}
