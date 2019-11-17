package model

// Goods or services that are part of a commercial trade agreement.
type LineItemDetails13 struct {

	// Identification assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Specifies the quantity of goods on a line in a trade transaction.
	Quantity *Quantity9 `xml:"Qty"`

	// Variance allowed in the quantity of goods.
	QuantityTolerance *PercentageTolerance1 `xml:"QtyTlrnce,omitempty"`

	// Amount of money for which goods or services are offered, sold, or bought.
	UnitPrice *UnitPrice18 `xml:"UnitPric,omitempty"`

	// Variance allowed on a price.
	PriceTolerance *PercentageTolerance1 `xml:"PricTlrnce,omitempty"`

	// Name of the product detailed in the corresponding line item.
	ProductName *Max70Text `xml:"PdctNm,omitempty"`

	// Identifies the product of the corresponding line item.
	ProductIdentifier []*ProductIdentifier2Choice `xml:"PdctIdr,omitempty"`

	// Identifies the characteristics of a product.
	ProductCharacteristics []*ProductCharacteristics1Choice `xml:"PdctChrtcs,omitempty"`

	// Identifies the category of product.
	ProductCategory []*ProductCategory1Choice `xml:"PdctCtgy,omitempty"`

	// Country from which the product originates.
	ProductOrigin []*CountryCode `xml:"PdctOrgn,omitempty"`

	// Specifies the shipment schedule for the goods.
	ShipmentSchedule *ShipmentSchedule2Choice `xml:"ShipmntSchdl,omitempty"`

	// Information related to the conveyance of goods.
	RoutingSummary *TransportMeans5 `xml:"RtgSummry,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment7 `xml:"Adjstmnt,omitempty"`

	// Maximum charges related to the conveyance of goods.
	FreightCharges *Charge24 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax23 `xml:"Tax,omitempty"`

	// Total amount of the line item after adjustments have been applied.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`

	// Specifies the applicable Incoterms and associated location. Latest version of Incoterms in effect at the date of message creation.
	Incoterms *Incoterms4 `xml:"Incotrms,omitempty"`
}

func (l *LineItemDetails13) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails13) AddQuantity() *Quantity9 {
	l.Quantity = new(Quantity9)
	return l.Quantity
}

func (l *LineItemDetails13) AddQuantityTolerance() *PercentageTolerance1 {
	l.QuantityTolerance = new(PercentageTolerance1)
	return l.QuantityTolerance
}

func (l *LineItemDetails13) AddUnitPrice() *UnitPrice18 {
	l.UnitPrice = new(UnitPrice18)
	return l.UnitPrice
}

func (l *LineItemDetails13) AddPriceTolerance() *PercentageTolerance1 {
	l.PriceTolerance = new(PercentageTolerance1)
	return l.PriceTolerance
}

func (l *LineItemDetails13) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails13) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails13) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails13) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails13) AddProductOrigin(value string) {
	l.ProductOrigin = append(l.ProductOrigin, (*CountryCode)(&value))
}

func (l *LineItemDetails13) AddShipmentSchedule() *ShipmentSchedule2Choice {
	l.ShipmentSchedule = new(ShipmentSchedule2Choice)
	return l.ShipmentSchedule
}

func (l *LineItemDetails13) AddRoutingSummary() *TransportMeans5 {
	l.RoutingSummary = new(TransportMeans5)
	return l.RoutingSummary
}

func (l *LineItemDetails13) AddAdjustment() *Adjustment7 {
	newValue := new(Adjustment7)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItemDetails13) AddFreightCharges() *Charge24 {
	l.FreightCharges = new(Charge24)
	return l.FreightCharges
}

func (l *LineItemDetails13) AddTax() *Tax23 {
	newValue := new(Tax23)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItemDetails13) SetTotalAmount(value, currency string) {
	l.TotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails13) AddIncoterms() *Incoterms4 {
	l.Incoterms = new(Incoterms4)
	return l.Incoterms
}
