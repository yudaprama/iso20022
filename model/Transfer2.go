package model

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer2 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Reference that identifies the whole transfer out transaction.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Date and optionally  time at which a transaction is completed and cleared, ie, securities are delivered.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date when the transfer was received and processed.
	TradeDate *ISODate `xml:"TradDt"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit1 `xml:"UnitsDtls,omitempty"`

	// Total quantity of securities settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`
}

func (t *Transfer2) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer2) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer2) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer2) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *Transfer2) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer2) AddUnitsDetails() *Unit1 {
	newValue := new(Unit1)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer2) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer2) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer2) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}
