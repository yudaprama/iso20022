package model

// Information about a payment against a purchase order.
type ReportLine3 struct {

	// Unique identification of the purchase order, assigned by the buyer.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies the adjustments applied to obtain the net amount.
	Adjustment []*Adjustment4 `xml:"Adjstmnt,omitempty"`

	// Net amount, after adjustments, intended to be paid.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`
}

func (r *ReportLine3) AddPurchaseOrderReference() *DocumentIdentification7 {
	r.PurchaseOrderReference = new(DocumentIdentification7)
	return r.PurchaseOrderReference
}

func (r *ReportLine3) AddAdjustment() *Adjustment4 {
	newValue := new(Adjustment4)
	r.Adjustment = append(r.Adjustment, newValue)
	return newValue
}

func (r *ReportLine3) SetNetAmount(value, currency string) {
	r.NetAmount = NewCurrencyAndAmount(value, currency)
}
