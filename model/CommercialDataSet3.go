package model

// Goods or services that are part of a commercial trade agreement.
type CommercialDataSet3 struct {

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
	Goods []*LineItem9 `xml:"Goods"`

	// Specifies the payment terms by means of a code and a limit in time.
	PaymentTerms []*PaymentTerms1 `xml:"PmtTerms"`

	// Specifies how the transaction should be settled.
	SettlementTerms *SettlementTerms2 `xml:"SttlmTerms"`
}

func (c *CommercialDataSet3) AddDataSetIdentification() *DocumentIdentification1 {
	c.DataSetIdentification = new(DocumentIdentification1)
	return c.DataSetIdentification
}

func (c *CommercialDataSet3) AddCommercialDocumentReference() *InvoiceIdentification1 {
	c.CommercialDocumentReference = new(InvoiceIdentification1)
	return c.CommercialDocumentReference
}

func (c *CommercialDataSet3) AddBuyer() *PartyIdentification26 {
	c.Buyer = new(PartyIdentification26)
	return c.Buyer
}

func (c *CommercialDataSet3) AddSeller() *PartyIdentification26 {
	c.Seller = new(PartyIdentification26)
	return c.Seller
}

func (c *CommercialDataSet3) AddBillTo() *PartyIdentification26 {
	c.BillTo = new(PartyIdentification26)
	return c.BillTo
}

func (c *CommercialDataSet3) AddGoods() *LineItem9 {
	newValue := new(LineItem9)
	c.Goods = append(c.Goods, newValue)
	return newValue
}

func (c *CommercialDataSet3) AddPaymentTerms() *PaymentTerms1 {
	newValue := new(PaymentTerms1)
	c.PaymentTerms = append(c.PaymentTerms, newValue)
	return newValue
}

func (c *CommercialDataSet3) AddSettlementTerms() *SettlementTerms2 {
	c.SettlementTerms = new(SettlementTerms2)
	return c.SettlementTerms
}
