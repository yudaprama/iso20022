package model

// Information about the units to settle.
type Unit3 struct {

	// Total quantity of securities to be settled.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Date upon which the investor purchased the financial instrument.
	AcquisitionDate *ISODate `xml:"AcqstnDt,omitempty"`

	// Certificate representing the security that is delivered.
	CertificateNumber []*Max35Text `xml:"CertNb,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Information related to the price of the transferred units.
	PriceDetails *UnitPrice12 `xml:"PricDtls,omitempty"`
}

func (u *Unit3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	u.UnitsNumber = new(FinancialInstrumentQuantity1)
	return u.UnitsNumber
}

func (u *Unit3) SetAcquisitionDate(value string) {
	u.AcquisitionDate = (*ISODate)(&value)
}

func (u *Unit3) AddCertificateNumber(value string) {
	u.CertificateNumber = append(u.CertificateNumber, (*Max35Text)(&value))
}

func (u *Unit3) SetGroup1Or2Units(value string) {
	u.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (u *Unit3) AddPriceDetails() *UnitPrice12 {
	u.PriceDetails = new(UnitPrice12)
	return u.PriceDetails
}
