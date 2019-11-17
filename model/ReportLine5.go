package model

// Information about a payment against a purchase order.
type ReportLine5 struct {

	// Unique identification of the purchase order, assigned by the buyer.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies the adjustments applied to obtain the net amount.
	Adjustment []*Adjustment6 `xml:"Adjstmnt,omitempty"`

	// Net amount, after adjustments, intended to be paid.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`
}

func (r *ReportLine5) AddPurchaseOrderReference() *DocumentIdentification7 {
	r.PurchaseOrderReference = new(DocumentIdentification7)
	return r.PurchaseOrderReference
}

func (r *ReportLine5) AddAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	r.Adjustment = append(r.Adjustment, newValue)
	return newValue
}

func (r *ReportLine5) SetNetAmount(value, currency string) {
	r.NetAmount = NewCurrencyAndAmount(value, currency)
}
