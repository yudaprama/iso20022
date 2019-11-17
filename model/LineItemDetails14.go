package model

// Goods or services that are part of a commercial trade agreement.
type LineItemDetails14 struct {

	// Sequential number assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Specifies the quantity of a product in a trade transaction.
	Quantity *Quantity9 `xml:"Qty"`

	// Amount of money for which goods or services are offered, sold, or bought.
	UnitPrice *UnitPrice18 `xml:"UnitPric,omitempty"`

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
	Adjustment []*Adjustment6 `xml:"Adjstmnt,omitempty"`

	// Charges related to the conveyance of goods.
	FreightCharges *Charge25 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax22 `xml:"Tax,omitempty"`

	// Total amount of the line item after adjustments have been applied.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`
}

func (l *LineItemDetails14) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails14) AddQuantity() *Quantity9 {
	l.Quantity = new(Quantity9)
	return l.Quantity
}

func (l *LineItemDetails14) AddUnitPrice() *UnitPrice18 {
	l.UnitPrice = new(UnitPrice18)
	return l.UnitPrice
}

func (l *LineItemDetails14) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails14) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails14) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails14) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails14) SetProductOrigin(value string) {
	l.ProductOrigin = (*CountryCode)(&value)
}

func (l *LineItemDetails14) AddAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItemDetails14) AddFreightCharges() *Charge25 {
	l.FreightCharges = new(Charge25)
	return l.FreightCharges
}

func (l *LineItemDetails14) AddTax() *Tax22 {
	newValue := new(Tax22)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItemDetails14) SetTotalAmount(value, currency string) {
	l.TotalAmount = NewCurrencyAndAmount(value, currency)
}
