package model

// Commercial details of a trade transaction between a buyer and a seller.
type LineItem9 struct {

	// Reference to the purchase order of the underlying transaction.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies whether this current invoice is the last submission against the purchase order referenced.
	FinalSubmission *YesNoIndicator `xml:"FnlSubmissn"`

	// Goods which are the subject of the invoice.
	CommercialLineItems []*LineItemDetails9 `xml:"ComrclLineItms"`

	// Specifies the total amount of all line items (incl. their adjustments).
	LineItemsTotalAmount *CurrencyAndAmount `xml:"LineItmsTtlAmt"`

	// Specifies the applicable Incoterm and associated location.
	Incoterms *Incoterms2 `xml:"Incotrms,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment4 `xml:"Adjstmnt,omitempty"`

	// Charges related to the conveyance of goods.
	FreightCharges *Charge13 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax12 `xml:"Tax,omitempty"`

	// Total net amount of a trade transaction. Total amount resulting from the gross amount plus freight charges, tax and plus/minus Adjustments.
	TotalNetAmount *CurrencyAndAmount `xml:"TtlNetAmt"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	BuyerDefinedInformation []*UserDefinedInformation1 `xml:"BuyrDfndInf,omitempty"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	SellerDefinedInformation []*UserDefinedInformation1 `xml:"SellrDfndInf,omitempty"`
}

func (l *LineItem9) AddPurchaseOrderReference() *DocumentIdentification7 {
	l.PurchaseOrderReference = new(DocumentIdentification7)
	return l.PurchaseOrderReference
}

func (l *LineItem9) SetFinalSubmission(value string) {
	l.FinalSubmission = (*YesNoIndicator)(&value)
}

func (l *LineItem9) AddCommercialLineItems() *LineItemDetails9 {
	newValue := new(LineItemDetails9)
	l.CommercialLineItems = append(l.CommercialLineItems, newValue)
	return newValue
}

func (l *LineItem9) SetLineItemsTotalAmount(value, currency string) {
	l.LineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem9) AddIncoterms() *Incoterms2 {
	l.Incoterms = new(Incoterms2)
	return l.Incoterms
}

func (l *LineItem9) AddAdjustment() *Adjustment4 {
	newValue := new(Adjustment4)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItem9) AddFreightCharges() *Charge13 {
	l.FreightCharges = new(Charge13)
	return l.FreightCharges
}

func (l *LineItem9) AddTax() *Tax12 {
	newValue := new(Tax12)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItem9) SetTotalNetAmount(value, currency string) {
	l.TotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem9) AddBuyerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.BuyerDefinedInformation = append(l.BuyerDefinedInformation, newValue)
	return newValue
}

func (l *LineItem9) AddSellerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.SellerDefinedInformation = append(l.SellerDefinedInformation, newValue)
	return newValue
}
