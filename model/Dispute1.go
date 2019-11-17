package model

// Provides the dispute details.
type Dispute1 struct {

	// Unique identification for the margin call request.
	MarginCallRequestIdentification *Max35Text `xml:"MrgnCallReqId"`

	// Disputed amount.
	DisputedAmount *ActiveCurrencyAndAmount `xml:"DsptdAmt"`

	// Date of dispute.
	DisputeDate *ISODate `xml:"DsptDt"`
}

func (d *Dispute1) SetMarginCallRequestIdentification(value string) {
	d.MarginCallRequestIdentification = (*Max35Text)(&value)
}

func (d *Dispute1) SetDisputedAmount(value, currency string) {
	d.DisputedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *Dispute1) SetDisputeDate(value string) {
	d.DisputeDate = (*ISODate)(&value)
}
