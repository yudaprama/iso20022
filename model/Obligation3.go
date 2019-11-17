package model

// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
type Obligation3 struct {

	// Defines one of the entities associated with the collateral agreement.
	PartyA *PartyIdentification33Choice `xml:"PtyA"`

	// Specifies the party that is acting on behalf of party A and that offers collateral management services.
	ServicingPartyA *PartyIdentification33Choice `xml:"SvcgPtyA,omitempty"`

	// Defines the other entity associated with the collateral agreement.
	PartyB *PartyIdentification33Choice `xml:"PtyB"`

	// Specifies the party that is acting on behalf of party B and that offers collateral management services.
	ServicingPartyB *PartyIdentification33Choice `xml:"SvcgPtyB,omitempty"`

	// Provides additional information on the Collateral Account of the Party delivering the collateral.
	CollateralAccountIdentification *CollateralAccount1 `xml:"CollAcctId,omitempty"`

	// Specifies the underlying business area or type of trade causing the collateral movement.
	ExposureType *ExposureType5Code `xml:"XpsrTp,omitempty"`

	// Indicates the close of business date on which the initiating party is valuing the margin call.
	ValuationDate *DateAndDateTimeChoice `xml:"ValtnDt"`
}

func (o *Obligation3) AddPartyA() *PartyIdentification33Choice {
	o.PartyA = new(PartyIdentification33Choice)
	return o.PartyA
}

func (o *Obligation3) AddServicingPartyA() *PartyIdentification33Choice {
	o.ServicingPartyA = new(PartyIdentification33Choice)
	return o.ServicingPartyA
}

func (o *Obligation3) AddPartyB() *PartyIdentification33Choice {
	o.PartyB = new(PartyIdentification33Choice)
	return o.PartyB
}

func (o *Obligation3) AddServicingPartyB() *PartyIdentification33Choice {
	o.ServicingPartyB = new(PartyIdentification33Choice)
	return o.ServicingPartyB
}

func (o *Obligation3) AddCollateralAccountIdentification() *CollateralAccount1 {
	o.CollateralAccountIdentification = new(CollateralAccount1)
	return o.CollateralAccountIdentification
}

func (o *Obligation3) SetExposureType(value string) {
	o.ExposureType = (*ExposureType5Code)(&value)
}

func (o *Obligation3) AddValuationDate() *DateAndDateTimeChoice {
	o.ValuationDate = new(DateAndDateTimeChoice)
	return o.ValuationDate
}
