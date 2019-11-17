package model

// Goods or services that are part of a commercial trade agreement.
type LineItemDetails7 struct {

	// Identification assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Specifies the quantity of goods on a line in a trade transaction.
	Quantity *Quantity4 `xml:"Qty"`

	// Variance allowed in the quantity of goods.
	QuantityTolerance *PercentageTolerance1 `xml:"QtyTlrnce,omitempty"`

	// Amount of money for which goods or services are offered, sold, or bought.
	UnitPrice *UnitPrice9 `xml:"UnitPric,omitempty"`

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
	ShipmentSchedule *ShipmentSchedule1Choice `xml:"ShipmntSchdl,omitempty"`

	// Information related to the conveyance of goods.
	RoutingSummary *TransportMeans1 `xml:"RtgSummry,omitempty"`

	// Specifies the applicable Incoterms and associated location. Latest version of Incoterms in effect at the date of message creation.
	Incoterms []*Incoterms1 `xml:"Incotrms,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment3 `xml:"Adjstmnt,omitempty"`

	// Maximum charges related to the conveyance of goods.
	FreightCharges *Charge12 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax13 `xml:"Tax,omitempty"`

	// Total amount of the line item after adjustments have been applied.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`
}

func (l *LineItemDetails7) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails7) AddQuantity() *Quantity4 {
	l.Quantity = new(Quantity4)
	return l.Quantity
}

func (l *LineItemDetails7) AddQuantityTolerance() *PercentageTolerance1 {
	l.QuantityTolerance = new(PercentageTolerance1)
	return l.QuantityTolerance
}

func (l *LineItemDetails7) AddUnitPrice() *UnitPrice9 {
	l.UnitPrice = new(UnitPrice9)
	return l.UnitPrice
}

func (l *LineItemDetails7) AddPriceTolerance() *PercentageTolerance1 {
	l.PriceTolerance = new(PercentageTolerance1)
	return l.PriceTolerance
}

func (l *LineItemDetails7) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails7) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails7) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails7) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails7) AddProductOrigin(value string) {
	l.ProductOrigin = append(l.ProductOrigin, (*CountryCode)(&value))
}

func (l *LineItemDetails7) AddShipmentSchedule() *ShipmentSchedule1Choice {
	l.ShipmentSchedule = new(ShipmentSchedule1Choice)
	return l.ShipmentSchedule
}

func (l *LineItemDetails7) AddRoutingSummary() *TransportMeans1 {
	l.RoutingSummary = new(TransportMeans1)
	return l.RoutingSummary
}

func (l *LineItemDetails7) AddIncoterms() *Incoterms1 {
	newValue := new(Incoterms1)
	l.Incoterms = append(l.Incoterms, newValue)
	return newValue
}

func (l *LineItemDetails7) AddAdjustment() *Adjustment3 {
	newValue := new(Adjustment3)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItemDetails7) AddFreightCharges() *Charge12 {
	l.FreightCharges = new(Charge12)
	return l.FreightCharges
}

func (l *LineItemDetails7) AddTax() *Tax13 {
	newValue := new(Tax13)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItemDetails7) SetTotalAmount(value, currency string) {
	l.TotalAmount = NewCurrencyAndAmount(value, currency)
}
