package model

// Creation/cancellation of investment units on the books of the fund or its designated agent, as a result of executing an investment fund order.
type InvestmentFundTransaction2 struct {

	// Type of investment fund transaction.
	TransactionType *TransactionType1CodeChoice `xml:"TxTp"`

	// Type of corporate action event.
	CorporateActionEventType *CorporateActionEventType1CodeChoice `xml:"CorpActnEvtTp"`

	// Status of an investment fund transaction.
	BookingStatus *TransactionStatus1Code `xml:"BookgSts,omitempty"`

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique identifier for an order, as assigned by the sell-side. The identifier must be unique within a single trading day.
	OrderReference *Max35Text `xml:"OrdrRef,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Unique technical identifier for an instance of a leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Unique identifier for an instance of a leg execution within a switch confirmation.
	LegExecutionIdentification *Max35Text `xml:"LegExctnId,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Indicates whether the cash payment with respect to the executed order is settled.
	SettledTransactionIndicator *YesNoIndicator `xml:"SttldTxInd"`

	// Indicates whether the executed order has a registered status on the books of the transfer agent.
	RegisteredTransactionIndicator *YesNoIndicator `xml:"RegdTxInd"`

	// Number of investment funds units.
	UnitsQuantity *FinancialInstrumentQuantity1 `xml:"UnitsQty"`

	// Direction of the transaction being reported, ie, securities are received (credited) or delivered (debited).
	CreditDebit *CreditDebitCode `xml:"CdtDbt"`

	// Transaction being reported is a reversal of previously reported transaction.
	Reversal *ReversalCode `xml:"Rvsl,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	GrossSettlementAmount *ActiveCurrencyAndAmount `xml:"GrssSttlmAmt,omitempty"`

	// Date on which the debtor expects the amount of money to be available to the creditor.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Indicates whether the order has been partially executed, ie, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Price at which the order was executed.
	PriceDetails *UnitPrice1 `xml:"PricDtls,omitempty"`
}

func (i *InvestmentFundTransaction2) AddTransactionType() *TransactionType1CodeChoice {
	i.TransactionType = new(TransactionType1CodeChoice)
	return i.TransactionType
}

func (i *InvestmentFundTransaction2) AddCorporateActionEventType() *CorporateActionEventType1CodeChoice {
	i.CorporateActionEventType = new(CorporateActionEventType1CodeChoice)
	return i.CorporateActionEventType
}

func (i *InvestmentFundTransaction2) SetBookingStatus(value string) {
	i.BookingStatus = (*TransactionStatus1Code)(&value)
}

func (i *InvestmentFundTransaction2) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *InvestmentFundTransaction2) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundTransaction2) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *InvestmentFundTransaction2) SetLegIdentification(value string) {
	i.LegIdentification = (*Max35Text)(&value)
}

func (i *InvestmentFundTransaction2) SetLegExecutionIdentification(value string) {
	i.LegExecutionIdentification = (*Max35Text)(&value)
}

func (i *InvestmentFundTransaction2) SetOrderDateTime(value string) {
	i.OrderDateTime = (*ISODateTime)(&value)
}

func (i *InvestmentFundTransaction2) SetSettledTransactionIndicator(value string) {
	i.SettledTransactionIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentFundTransaction2) SetRegisteredTransactionIndicator(value string) {
	i.RegisteredTransactionIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentFundTransaction2) AddUnitsQuantity() *FinancialInstrumentQuantity1 {
	i.UnitsQuantity = new(FinancialInstrumentQuantity1)
	return i.UnitsQuantity
}

func (i *InvestmentFundTransaction2) SetCreditDebit(value string) {
	i.CreditDebit = (*CreditDebitCode)(&value)
}

func (i *InvestmentFundTransaction2) SetReversal(value string) {
	i.Reversal = (*ReversalCode)(&value)
}

func (i *InvestmentFundTransaction2) SetGrossSettlementAmount(value, currency string) {
	i.GrossSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvestmentFundTransaction2) SetSettlementDate(value string) {
	i.SettlementDate = (*ISODate)(&value)
}

func (i *InvestmentFundTransaction2) AddTradeDateTime() *DateAndDateTimeChoice {
	i.TradeDateTime = new(DateAndDateTimeChoice)
	return i.TradeDateTime
}

func (i *InvestmentFundTransaction2) SetCumDividendIndicator(value string) {
	i.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentFundTransaction2) SetPartiallyExecutedIndicator(value string) {
	i.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentFundTransaction2) AddPriceDetails() *UnitPrice1 {
	i.PriceDetails = new(UnitPrice1)
	return i.PriceDetails
}
