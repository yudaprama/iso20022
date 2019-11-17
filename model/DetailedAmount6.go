package model

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount6 struct {

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Date and time of the payment.
	DateTime *ISODateTime `xml:"DtTm"`

	// Card data entry mode for the related payment.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd,omitempty"`

	// Data of an integrated circuit card application for the related payment.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Short description of the amount to display or print.
	Label *Max140Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount6) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount6) SetDateTime(value string) {
	d.DateTime = (*ISODateTime)(&value)
}

func (d *DetailedAmount6) SetCardDataEntryMode(value string) {
	d.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (d *DetailedAmount6) SetICCRelatedData(value string) {
	d.ICCRelatedData = (*Max10000Binary)(&value)
}

func (d *DetailedAmount6) SetLabel(value string) {
	d.Label = (*Max140Text)(&value)
}
