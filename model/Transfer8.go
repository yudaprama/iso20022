package model

// Parameters applied to the settlement of a security transfer.
type Transfer8 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Total quantity of securities to be settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer8) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *Transfer8) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer8) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer8) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *Transfer8) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer8) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer8) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer8) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer8) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer8) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer8) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}
