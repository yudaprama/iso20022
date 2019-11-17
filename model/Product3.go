package model

// Purchased item.
type Product3 struct {

	// Identification of the item inside the purchase transaction.
	ItemIdentification *Max35Text `xml:"ItmId,omitempty"`

	// Product code of the item.
	ProductCode *Max70Text `xml:"PdctCd"`

	// Additional product code related to the product.
	AdditionalProductCode *Max70Text `xml:"AddtlPdctCd,omitempty"`

	// Unit of measure of the item purchased.
	UnitOfMeasure *UnitOfMeasure6Code `xml:"UnitOfMeasr,omitempty"`

	// Product quantity.
	ProductQuantity *DecimalNumber `xml:"PdctQty,omitempty"`

	// Price per unit of product.
	UnitPrice *ImpliedCurrencyAndAmount `xml:"UnitPric,omitempty"`

	// Sign of the unit price.
	UnitPriceSign *PlusOrMinusIndicator `xml:"UnitPricSgn,omitempty"`

	// Monetary value of purchased product.
	ProductAmount *ImpliedCurrencyAndAmount `xml:"PdctAmt"`

	// Sign of the product amount.
	ProductAmountSign *PlusOrMinusIndicator `xml:"PdctAmtSgn,omitempty"`

	// Value added tax amount of the item included in the product amount.
	ValueAddedTax *ImpliedCurrencyAndAmount `xml:"ValAddedTax,omitempty"`

	// Information on tax paid on the product.
	TaxType *Max35Text `xml:"TaxTp,omitempty"`

	// Description of the product or item.
	ProductDescription *Max140Text `xml:"PdctDesc,omitempty"`

	// Location of the delivery of the item, for instance pump number or parking bay.
	DeliveryLocation *Max10Text `xml:"DlvryLctn,omitempty"`

	// Identify the method of delivery or distribution of the item.
	DeliveryService *AttendanceContext2Code `xml:"DlvrySvc,omitempty"`
}

func (p *Product3) SetItemIdentification(value string) {
	p.ItemIdentification = (*Max35Text)(&value)
}

func (p *Product3) SetProductCode(value string) {
	p.ProductCode = (*Max70Text)(&value)
}

func (p *Product3) SetAdditionalProductCode(value string) {
	p.AdditionalProductCode = (*Max70Text)(&value)
}

func (p *Product3) SetUnitOfMeasure(value string) {
	p.UnitOfMeasure = (*UnitOfMeasure6Code)(&value)
}

func (p *Product3) SetProductQuantity(value string) {
	p.ProductQuantity = (*DecimalNumber)(&value)
}

func (p *Product3) SetUnitPrice(value, currency string) {
	p.UnitPrice = NewImpliedCurrencyAndAmount(value, currency)
}

func (p *Product3) SetUnitPriceSign(value string) {
	p.UnitPriceSign = (*PlusOrMinusIndicator)(&value)
}

func (p *Product3) SetProductAmount(value, currency string) {
	p.ProductAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (p *Product3) SetProductAmountSign(value string) {
	p.ProductAmountSign = (*PlusOrMinusIndicator)(&value)
}

func (p *Product3) SetValueAddedTax(value, currency string) {
	p.ValueAddedTax = NewImpliedCurrencyAndAmount(value, currency)
}

func (p *Product3) SetTaxType(value string) {
	p.TaxType = (*Max35Text)(&value)
}

func (p *Product3) SetProductDescription(value string) {
	p.ProductDescription = (*Max140Text)(&value)
}

func (p *Product3) SetDeliveryLocation(value string) {
	p.DeliveryLocation = (*Max10Text)(&value)
}

func (p *Product3) SetDeliveryService(value string) {
	p.DeliveryService = (*AttendanceContext2Code)(&value)
}
