package model

// Goods or services that are part of a commercial trade agreement.
type LineItemDetails9 struct {

	// Sequential number assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Specifies the quantity of a product in a trade transaction.
	Quantity *Quantity4 `xml:"Qty"`

	// Amount of money for which goods or services are offered, sold, or bought.
	UnitPrice *UnitPrice9 `xml:"UnitPric,omitempty"`

	// Name of the product detailed in the corresponding line item.
	ProductName *Max70Text `xml:"PdctNm,omitempty"`

	// Identifies the product of the corresponding line item.
	ProductIdentifier []*ProductIdentifier2Choice `xml:"PdctIdr,omitempty"`

	// Identifies the characteristics of product.
	ProductCharacteristics []*ProductCharacteristics1Choice `xml:"PdctChrtcs,omitempty"`

	// Identifies the category of product.
	ProductCategory []*ProductCategory1Choice `xml:"PdctCtgy,omitempty"`

	// Country of origin of the goods.
	ProductOrigin *CountryCode `xml:"PdctOrgn,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment4 `xml:"Adjstmnt,omitempty"`

	// Charges related to the conveyance of goods.
	FreightCharges *Charge13 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax12 `xml:"Tax,omitempty"`

	// Total amount of the line item after adjustments have been applied.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`
}

func (l *LineItemDetails9) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails9) AddQuantity() *Quantity4 {
	l.Quantity = new(Quantity4)
	return l.Quantity
}

func (l *LineItemDetails9) AddUnitPrice() *UnitPrice9 {
	l.UnitPrice = new(UnitPrice9)
	return l.UnitPrice
}

func (l *LineItemDetails9) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails9) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails9) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails9) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails9) SetProductOrigin(value string) {
	l.ProductOrigin = (*CountryCode)(&value)
}

func (l *LineItemDetails9) AddAdjustment() *Adjustment4 {
	newValue := new(Adjustment4)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItemDetails9) AddFreightCharges() *Charge13 {
	l.FreightCharges = new(Charge13)
	return l.FreightCharges
}

func (l *LineItemDetails9) AddTax() *Tax12 {
	newValue := new(Tax12)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItemDetails9) SetTotalAmount(value, currency string) {
	l.TotalAmount = NewCurrencyAndAmount(value, currency)
}
