package model

// Commercial details of a trade transaction between a buyer and a seller.
type LineItem15 struct {

	// Reference to the purchase order of the underlying transaction.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies whether this current invoice is the last submission against the purchase order referenced.
	FinalSubmission *YesNoIndicator `xml:"FnlSubmissn"`

	// Goods which are the subject of the invoice.
	CommercialLineItems []*LineItemDetails14 `xml:"ComrclLineItms"`

	// Specifies the total amount of all line items (incl. their adjustments).
	LineItemsTotalAmount *CurrencyAndAmount `xml:"LineItmsTtlAmt"`

	// Variance on price for the goods.
	Adjustment []*Adjustment6 `xml:"Adjstmnt,omitempty"`

	// Charges related to the conveyance of goods.
	FreightCharges *Charge25 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax22 `xml:"Tax,omitempty"`

	// Total net amount of a trade transaction. Total amount resulting from the gross amount plus freight charges, tax and plus/minus Adjustments.
	TotalNetAmount *CurrencyAndAmount `xml:"TtlNetAmt"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	BuyerDefinedInformation []*UserDefinedInformation1 `xml:"BuyrDfndInf,omitempty"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	SellerDefinedInformation []*UserDefinedInformation1 `xml:"SellrDfndInf,omitempty"`

	// Specifies the applicable Incoterm and associated location.
	Incoterms *Incoterms4 `xml:"Incotrms,omitempty"`
}

func (l *LineItem15) AddPurchaseOrderReference() *DocumentIdentification7 {
	l.PurchaseOrderReference = new(DocumentIdentification7)
	return l.PurchaseOrderReference
}

func (l *LineItem15) SetFinalSubmission(value string) {
	l.FinalSubmission = (*YesNoIndicator)(&value)
}

func (l *LineItem15) AddCommercialLineItems() *LineItemDetails14 {
	newValue := new(LineItemDetails14)
	l.CommercialLineItems = append(l.CommercialLineItems, newValue)
	return newValue
}

func (l *LineItem15) SetLineItemsTotalAmount(value, currency string) {
	l.LineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem15) AddAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItem15) AddFreightCharges() *Charge25 {
	l.FreightCharges = new(Charge25)
	return l.FreightCharges
}

func (l *LineItem15) AddTax() *Tax22 {
	newValue := new(Tax22)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItem15) SetTotalNetAmount(value, currency string) {
	l.TotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem15) AddBuyerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.BuyerDefinedInformation = append(l.BuyerDefinedInformation, newValue)
	return newValue
}

func (l *LineItem15) AddSellerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.SellerDefinedInformation = append(l.SellerDefinedInformation, newValue)
	return newValue
}

func (l *LineItem15) AddIncoterms() *Incoterms4 {
	l.Incoterms = new(Incoterms4)
	return l.Incoterms
}
