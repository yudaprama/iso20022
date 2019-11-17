package model

// Calculation of the current situation of a line item as a result of the submission of a commercial dataset.
type LineItemDetails12 struct {

	// Sequential number assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Name of the product detailed in the corresponding line item.
	ProductName *Max70Text `xml:"PdctNm,omitempty"`

	// Identifies the product of the corresponding line item.
	ProductIdentifier []*ProductIdentifier2Choice `xml:"PdctIdr,omitempty"`

	// Identifies the characteristic of a product.
	ProductCharacteristics []*ProductCharacteristics1Choice `xml:"PdctChrtcs,omitempty"`

	// Identifies the category of product.
	ProductCategory []*ProductCategory1Choice `xml:"PdctCtgy,omitempty"`

	// Quantity ordered for a line as indicated in the baseline.
	OrderedQuantity *Quantity9 `xml:"OrdrdQty"`

	// Quantity accepted by data set submission.
	AcceptedQuantity *Quantity9 `xml:"AccptdQty"`

	// Difference between the ordered quantity and the accepted one.
	OutstandingQuantity *Quantity9 `xml:"OutsdngQty"`

	// Quantity of a product for which a mismatched data set has been submitted and has not yet been accepted or rejected..
	PendingQuantity *Quantity9 `xml:"PdgQty"`

	// Variance allowed on the quantity of goods.
	QuantityTolerance *PercentageTolerance1 `xml:"QtyTlrnce,omitempty"`

	// Price multiplied by the ordered quantity for a line as indicated in the baseline.
	OrderedAmount *CurrencyAndAmount `xml:"OrdrdAmt"`

	// Price multiplied by the accepted quantity for a line.
	AcceptedAmount *CurrencyAndAmount `xml:"AccptdAmt"`

	// Price multiplied by the outstanding quantity for a line.
	OutstandingAmount *CurrencyAndAmount `xml:"OutsdngAmt"`

	// Price multiplied by the pending quantity for a line.
	PendingAmount *CurrencyAndAmount `xml:"PdgAmt"`

	// Variance on price for the goods.
	PriceTolerance *PercentageTolerance1 `xml:"PricTlrnce,omitempty"`
}

func (l *LineItemDetails12) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails12) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails12) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails12) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails12) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails12) AddOrderedQuantity() *Quantity9 {
	l.OrderedQuantity = new(Quantity9)
	return l.OrderedQuantity
}

func (l *LineItemDetails12) AddAcceptedQuantity() *Quantity9 {
	l.AcceptedQuantity = new(Quantity9)
	return l.AcceptedQuantity
}

func (l *LineItemDetails12) AddOutstandingQuantity() *Quantity9 {
	l.OutstandingQuantity = new(Quantity9)
	return l.OutstandingQuantity
}

func (l *LineItemDetails12) AddPendingQuantity() *Quantity9 {
	l.PendingQuantity = new(Quantity9)
	return l.PendingQuantity
}

func (l *LineItemDetails12) AddQuantityTolerance() *PercentageTolerance1 {
	l.QuantityTolerance = new(PercentageTolerance1)
	return l.QuantityTolerance
}

func (l *LineItemDetails12) SetOrderedAmount(value, currency string) {
	l.OrderedAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails12) SetAcceptedAmount(value, currency string) {
	l.AcceptedAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails12) SetOutstandingAmount(value, currency string) {
	l.OutstandingAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails12) SetPendingAmount(value, currency string) {
	l.PendingAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails12) AddPriceTolerance() *PercentageTolerance1 {
	l.PriceTolerance = new(PercentageTolerance1)
	return l.PriceTolerance
}
