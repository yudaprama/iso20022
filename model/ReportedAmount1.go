package model

// Amount reported.
type ReportedAmount1 struct {

	// Identification of the reported amount.
	Identification *Max35Text `xml:"Id"`

	// Type of reported amount.
	Type *ExternalUndertakingAmountType1Code `xml:"Tp"`

	// Amount reported.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (r *ReportedAmount1) SetIdentification(value string) {
	r.Identification = (*Max35Text)(&value)
}

func (r *ReportedAmount1) SetType(value string) {
	r.Type = (*ExternalUndertakingAmountType1Code)(&value)
}

func (r *ReportedAmount1) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAndAmount(value, currency)
}
