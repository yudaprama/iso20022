package model

// Goods or services that are part of a commercial trade agreement.
type CommercialDataSet5 struct {

	// Identifies the commercial data set
	DataSetIdentification *DocumentIdentification1 `xml:"DataSetId"`

	// Reference to the identification of  the underlying commercial document.
	CommercialDocumentReference *InvoiceIdentification1 `xml:"ComrclDocRef"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *PartyIdentification26 `xml:"Buyr"`

	// Party that sells goods or services, or a financial instrument.
	Seller *PartyIdentification26 `xml:"Sellr"`

	// Party to be invoiced for the purchase.
	BillTo *PartyIdentification26 `xml:"BllTo,omitempty"`

	// Information about the goods and/or services of the underlying transaction.
	Goods []*LineItem15 `xml:"Goods"`

	// Specifies the payment terms by means of a code and a limit in time.
	PaymentTerms []*PaymentTerms4 `xml:"PmtTerms"`

	// Specifies how the transaction should be settled.
	SettlementTerms *SettlementTerms3 `xml:"SttlmTerms"`
}

func (c *CommercialDataSet5) AddDataSetIdentification() *DocumentIdentification1 {
	c.DataSetIdentification = new(DocumentIdentification1)
	return c.DataSetIdentification
}

func (c *CommercialDataSet5) AddCommercialDocumentReference() *InvoiceIdentification1 {
	c.CommercialDocumentReference = new(InvoiceIdentification1)
	return c.CommercialDocumentReference
}

func (c *CommercialDataSet5) AddBuyer() *PartyIdentification26 {
	c.Buyer = new(PartyIdentification26)
	return c.Buyer
}

func (c *CommercialDataSet5) AddSeller() *PartyIdentification26 {
	c.Seller = new(PartyIdentification26)
	return c.Seller
}

func (c *CommercialDataSet5) AddBillTo() *PartyIdentification26 {
	c.BillTo = new(PartyIdentification26)
	return c.BillTo
}

func (c *CommercialDataSet5) AddGoods() *LineItem15 {
	newValue := new(LineItem15)
	c.Goods = append(c.Goods, newValue)
	return newValue
}

func (c *CommercialDataSet5) AddPaymentTerms() *PaymentTerms4 {
	newValue := new(PaymentTerms4)
	c.PaymentTerms = append(c.PaymentTerms, newValue)
	return newValue
}

func (c *CommercialDataSet5) AddSettlementTerms() *SettlementTerms3 {
	c.SettlementTerms = new(SettlementTerms3)
	return c.SettlementTerms
}
