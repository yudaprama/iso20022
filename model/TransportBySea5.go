package model

// Information related for the transportation of goods by sea.
type TransportBySea5 struct {

	// Identifies the port where the goods are loaded on board the ship.
	PortOfLoading *Max35Text `xml:"PortOfLoadng"`

	// Identifies the port where the goods are discharged.
	PortOfDischarge *Max35Text `xml:"PortOfDschrge"`

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

	// Name of the master or captain of a vessel that signs the document for example, bill of lading, charter party bill of lading, non-negotiable sea waybill or multimodal transport document that evidences shipment of the goods from a port of loading to a port of discharge.
	MasterName *Max70Text `xml:"MstrNm,omitempty"`

	// Name of the company or individual that signs a charter party bill of lading that evidences shipment of the goods from a port of loading to a port of discharge and acts in the capacity of charterer.
	ChartererName *Max70Text `xml:"ChrtrrNm,omitempty"`

	// Name of the company or individual that signs a charter party bill of lading that evidences shipment of the goods from a port of loading to a port of discharge and acts in the capacity of owner;
	OwnerName *Max70Text `xml:"OwnrNm,omitempty"`

	// International Maritime Organisation identification of a ship. The IMO identification number scheme was introduced in 1987 as a measure aimed at enhancing maritime safety and pollution prevention and to facilitate the prevention of maritime fraud. It assigns a permanent number to each vessel for identification purposes. This number remains unchanged upon transfer of the vessel to other flag(s) and is inserted in all vessel certificates. The IMO identification number is made up of the three letters "IMO" followed by a seven-digit number assigned to all vessels by IHS FairPlay (formerly known as Lloyd's Register-Fairplay). This is a unique seven digit number that is assigned to vessels and aids banks in determining whether a vessel is subject to an order that would not permit a bank to handle a certain transaction under local or international laws.
	IMONumber *Exact7NumericText `xml:"IMONb,omitempty"`

	// Identifies the voyage by sea.
	VoyageNumber *Max35Text `xml:"VygNb,omitempty"`
}

func (t *TransportBySea5) SetPortOfLoading(value string) {
	t.PortOfLoading = (*Max35Text)(&value)
}

func (t *TransportBySea5) SetPortOfDischarge(value string) {
	t.PortOfDischarge = (*Max35Text)(&value)
}

func (t *TransportBySea5) SetVesselName(value string) {
	t.VesselName = (*Max70Text)(&value)
}

func (t *TransportBySea5) SetSeaCarrierName(value string) {
	t.SeaCarrierName = (*Max70Text)(&value)
}

func (t *TransportBySea5) SetSeaCarrierCountry(value string) {
	t.SeaCarrierCountry = (*CountryCode)(&value)
}

func (t *TransportBySea5) SetCarrierAgentName(value string) {
	t.CarrierAgentName = (*Max70Text)(&value)
}

func (t *TransportBySea5) SetCarrierAgentCountry(value string) {
	t.CarrierAgentCountry = (*CountryCode)(&value)
}

func (t *TransportBySea5) SetMasterName(value string) {
	t.MasterName = (*Max70Text)(&value)
}

func (t *TransportBySea5) SetChartererName(value string) {
	t.ChartererName = (*Max70Text)(&value)
}

func (t *TransportBySea5) SetOwnerName(value string) {
	t.OwnerName = (*Max70Text)(&value)
}

func (t *TransportBySea5) SetIMONumber(value string) {
	t.IMONumber = (*Exact7NumericText)(&value)
}

func (t *TransportBySea5) SetVoyageNumber(value string) {
	t.VoyageNumber = (*Max35Text)(&value)
}
