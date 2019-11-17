package model

// Information about a payment against a commercial invoice.
type ReportLine4 struct {

	// Reference to the identification of the underlying commercial document.
	CommercialDocumentReference *InvoiceIdentification1 `xml:"ComrclDocRef"`

	// Specifies the adjustments applied to obtain the net amount.
	Adjustment []*Adjustment4 `xml:"Adjstmnt,omitempty"`

	// Net amount, after adjustments, intended to be paid.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`

	// Specifies how the net amount to be paid is related to different purchase orders.
	BreakdownByPurchaseOrder []*ReportLine2 `xml:"BrkdwnByPurchsOrdr"`
}

func (r *ReportLine4) AddCommercialDocumentReference() *InvoiceIdentification1 {
	r.CommercialDocumentReference = new(InvoiceIdentification1)
	return r.CommercialDocumentReference
}

func (r *ReportLine4) AddAdjustment() *Adjustment4 {
	newValue := new(Adjustment4)
	r.Adjustment = append(r.Adjustment, newValue)
	return newValue
}

func (r *ReportLine4) SetNetAmount(value, currency string) {
	r.NetAmount = NewCurrencyAndAmount(value, currency)
}

func (r *ReportLine4) AddBreakdownByPurchaseOrder() *ReportLine2 {
	newValue := new(ReportLine2)
	r.BreakdownByPurchaseOrder = append(r.BreakdownByPurchaseOrder, newValue)
	return newValue
}
