package model

// Information about a payment against a commercial invoice.
type ReportLine6 struct {

	// Reference to the identification of the underlying commercial document.
	CommercialDocumentReference *InvoiceIdentification1 `xml:"ComrclDocRef"`

	// Specifies the adjustments applied to obtain the net amount.
	Adjustment []*Adjustment6 `xml:"Adjstmnt,omitempty"`

	// Net amount, after adjustments, intended to be paid.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`

	// Specifies how the net amount to be paid is related to different purchase orders.
	BreakdownByPurchaseOrder []*ReportLine7 `xml:"BrkdwnByPurchsOrdr"`
}

func (r *ReportLine6) AddCommercialDocumentReference() *InvoiceIdentification1 {
	r.CommercialDocumentReference = new(InvoiceIdentification1)
	return r.CommercialDocumentReference
}

func (r *ReportLine6) AddAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	r.Adjustment = append(r.Adjustment, newValue)
	return newValue
}

func (r *ReportLine6) SetNetAmount(value, currency string) {
	r.NetAmount = NewCurrencyAndAmount(value, currency)
}

func (r *ReportLine6) AddBreakdownByPurchaseOrder() *ReportLine7 {
	newValue := new(ReportLine7)
	r.BreakdownByPurchaseOrder = append(r.BreakdownByPurchaseOrder, newValue)
	return newValue
}
