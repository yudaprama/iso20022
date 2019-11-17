package model

// Parameters applied to the settlement of a security transfer.
type Transfer3 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Date and time at which the securities are to be delivered or received.
	TransferDate *DateFormat1Choice `xml:"TrfDt,omitempty"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd"`
}

func (t *Transfer3) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer3) AddTransferDate() *DateFormat1Choice {
	t.TransferDate = new(DateFormat1Choice)
	return t.TransferDate
}

func (t *Transfer3) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer3) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}
