package model

// Allows to specify other collateral type by providing a description and the quantity.
type OtherTypeOfCollateral2 struct {

	// Provides details about the collateral.
	Description *Max140Text `xml:"Desc"`

	// Quantity of other collateral.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty,omitempty"`
}

func (o *OtherTypeOfCollateral2) SetDescription(value string) {
	o.Description = (*Max140Text)(&value)
}

func (o *OtherTypeOfCollateral2) AddQuantity() *FinancialInstrumentQuantity1Choice {
	o.Quantity = new(FinancialInstrumentQuantity1Choice)
	return o.Quantity
}
