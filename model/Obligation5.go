package model

// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
type Obligation5 struct {

	// Defines one of the entities associated with the collateral agreement.
	PartyA *PartyIdentification100Choice `xml:"PtyA"`

	// Specifies the party that is acting on behalf of party A and that offers collateral management services.
	ServicingPartyA *PartyIdentification100Choice `xml:"SvcgPtyA,omitempty"`

	// Defines the other entity associated with the collateral agreement.
	PartyB *PartyIdentification100Choice `xml:"PtyB"`

	// Specifies the party that is acting on behalf of party B and that offers collateral management services.
	ServicingPartyB *PartyIdentification100Choice `xml:"SvcgPtyB,omitempty"`

	// Provides additional information on the collateral account of the party delivering/receiving the collateral.
	CollateralAccountIdentification *CollateralAccount3 `xml:"CollAcctId,omitempty"`

	// Specifies the underlying business area or type of trade causing the collateral movement.
	ExposureType *ExposureType5Code `xml:"XpsrTp,omitempty"`

	// Indicates the close of business date on which the initiating party is valuing the margin call.
	ValuationDate *DateAndDateTimeChoice `xml:"ValtnDt"`
}

func (o *Obligation5) AddPartyA() *PartyIdentification100Choice {
	o.PartyA = new(PartyIdentification100Choice)
	return o.PartyA
}

func (o *Obligation5) AddServicingPartyA() *PartyIdentification100Choice {
	o.ServicingPartyA = new(PartyIdentification100Choice)
	return o.ServicingPartyA
}

func (o *Obligation5) AddPartyB() *PartyIdentification100Choice {
	o.PartyB = new(PartyIdentification100Choice)
	return o.PartyB
}

func (o *Obligation5) AddServicingPartyB() *PartyIdentification100Choice {
	o.ServicingPartyB = new(PartyIdentification100Choice)
	return o.ServicingPartyB
}

func (o *Obligation5) AddCollateralAccountIdentification() *CollateralAccount3 {
	o.CollateralAccountIdentification = new(CollateralAccount3)
	return o.CollateralAccountIdentification
}

func (o *Obligation5) SetExposureType(value string) {
	o.ExposureType = (*ExposureType5Code)(&value)
}

func (o *Obligation5) AddValuationDate() *DateAndDateTimeChoice {
	o.ValuationDate = new(DateAndDateTimeChoice)
	return o.ValuationDate
}
