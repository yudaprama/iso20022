package model

// Commercial details of a trade transaction between a buyer and a seller.
type LineItem11 struct {

	// Information about the goods and/or services of a trade transaction.
	GoodsDescription *Max70Text `xml:"GoodsDesc,omitempty"`

	// Indicates whether or not partial shipments are allowed.
	PartialShipment *YesNoIndicator `xml:"PrtlShipmnt"`

	// Indicates whether or not transshipment of goods is allowed.
	TransShipment *YesNoIndicator `xml:"TrnsShipmnt,omitempty"`

	// Specifies an earliest shipment date and a latest shipment date.
	ShipmentDateRange *ShipmentDateRange1 `xml:"ShipmntDtRg,omitempty"`

	// Goods or services that are purchased.
	LineItemDetails []*LineItemDetails10 `xml:"LineItmDtls"`

	// Specifies the total amount of all line items (incl. their adjustments).
	LineItemsTotalAmount *CurrencyAndAmount `xml:"LineItmsTtlAmt"`

	// Information related to the conveyance of goods.
	RoutingSummary *TransportMeans5 `xml:"RtgSummry,omitempty"`

	// Specifies the applicable Incoterms and associated location. Latest version of Incoterms in effect at the date of message creation.
	Incoterms *Incoterms4 `xml:"Incotrms,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment7 `xml:"Adjstmnt,omitempty"`

	// Maximum charges related to the conveyance of goods.
	FreightCharges *Charge24 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax23 `xml:"Tax,omitempty"`

	// Total net amount of a trade transaction. Total amount resulting from the gross amount plus freight charges, tax and plus/minus Adjustments.
	TotalNetAmount *CurrencyAndAmount `xml:"TtlNetAmt"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	BuyerDefinedInformation []*UserDefinedInformation1 `xml:"BuyrDfndInf,omitempty"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	SellerDefinedInformation []*UserDefinedInformation1 `xml:"SellrDfndInf,omitempty"`
}

func (l *LineItem11) SetGoodsDescription(value string) {
	l.GoodsDescription = (*Max70Text)(&value)
}

func (l *LineItem11) SetPartialShipment(value string) {
	l.PartialShipment = (*YesNoIndicator)(&value)
}

func (l *LineItem11) SetTransShipment(value string) {
	l.TransShipment = (*YesNoIndicator)(&value)
}

func (l *LineItem11) AddShipmentDateRange() *ShipmentDateRange1 {
	l.ShipmentDateRange = new(ShipmentDateRange1)
	return l.ShipmentDateRange
}

func (l *LineItem11) AddLineItemDetails() *LineItemDetails10 {
	newValue := new(LineItemDetails10)
	l.LineItemDetails = append(l.LineItemDetails, newValue)
	return newValue
}

func (l *LineItem11) SetLineItemsTotalAmount(value, currency string) {
	l.LineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem11) AddRoutingSummary() *TransportMeans5 {
	l.RoutingSummary = new(TransportMeans5)
	return l.RoutingSummary
}

func (l *LineItem11) AddIncoterms() *Incoterms4 {
	l.Incoterms = new(Incoterms4)
	return l.Incoterms
}

func (l *LineItem11) AddAdjustment() *Adjustment7 {
	newValue := new(Adjustment7)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItem11) AddFreightCharges() *Charge24 {
	l.FreightCharges = new(Charge24)
	return l.FreightCharges
}

func (l *LineItem11) AddTax() *Tax23 {
	newValue := new(Tax23)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItem11) SetTotalNetAmount(value, currency string) {
	l.TotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem11) AddBuyerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.BuyerDefinedInformation = append(l.BuyerDefinedInformation, newValue)
	return newValue
}

func (l *LineItem11) AddSellerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.SellerDefinedInformation = append(l.SellerDefinedInformation, newValue)
	return newValue
}
