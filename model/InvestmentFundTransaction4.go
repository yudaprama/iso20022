package model

// Creation/cancellation of investment units on the books of the fund or its designated agent, as a result of executing an investment fund order.
type InvestmentFundTransaction4 struct {

	// Underlying transaction or corporate action.
	EventType *TransactionType1Choice `xml:"EvtTp"`

	// Status of an investment fund transaction.
	BookingStatus *TransactionStatus1Code `xml:"BookgSts,omitempty"`

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef,omitempty"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Unique technical identifier for an instance of a leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Unique identifier for an instance of a leg execution within a switch confirmation.
	LegExecutionIdentification *Max35Text `xml:"LegExctnId,omitempty"`

	// Date and time at which the order was placed by the investor or its agent.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Indicates whether the cash payment with respect to the executed order is settled.
	SettledTransactionIndicator *YesNoIndicator `xml:"SttldTxInd"`

	// Indicates whether the executed order has a registered status on the books of the transfer agent.
	RegisteredTransactionIndicator *YesNoIndicator `xml:"RegdTxInd"`

	// Number of investment funds units.
	UnitsQuantity *FinancialInstrumentQuantity1 `xml:"UnitsQty"`

	// Direction of the transaction being reported, is, securities are received (credited) or delivered (debited).
	CreditDebit *CreditDebitCode `xml:"CdtDbt"`

	// Transaction being reported is a reversal of previously reported transaction.
	Reversal *ReversalCode `xml:"Rvsl,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the debtor expects the amount of money to be available to the creditor.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Indicates whether the dividend is included, that is, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Indicates whether the order has been partially executed, that is, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Price at which the order was executed.
	PriceDetails *UnitPrice20 `xml:"PricDtls,omitempty"`
}

func (i *InvestmentFundTransaction4) AddEventType() *TransactionType1Choice {
	i.EventType = new(TransactionType1Choice)
	return i.EventType
}

func (i *InvestmentFundTransaction4) SetBookingStatus(value string) {
	i.BookingStatus = (*TransactionStatus1Code)(&value)
}

func (i *InvestmentFundTransaction4) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *InvestmentFundTransaction4) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundTransaction4) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundTransaction4) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *InvestmentFundTransaction4) SetLegIdentification(value string) {
	i.LegIdentification = (*Max35Text)(&value)
}

func (i *InvestmentFundTransaction4) SetLegExecutionIdentification(value string) {
	i.LegExecutionIdentification = (*Max35Text)(&value)
}

func (i *InvestmentFundTransaction4) SetOrderDateTime(value string) {
	i.OrderDateTime = (*ISODateTime)(&value)
}

func (i *InvestmentFundTransaction4) SetSettledTransactionIndicator(value string) {
	i.SettledTransactionIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentFundTransaction4) SetRegisteredTransactionIndicator(value string) {
	i.RegisteredTransactionIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentFundTransaction4) AddUnitsQuantity() *FinancialInstrumentQuantity1 {
	i.UnitsQuantity = new(FinancialInstrumentQuantity1)
	return i.UnitsQuantity
}

func (i *InvestmentFundTransaction4) SetCreditDebit(value string) {
	i.CreditDebit = (*CreditDebitCode)(&value)
}

func (i *InvestmentFundTransaction4) SetReversal(value string) {
	i.Reversal = (*ReversalCode)(&value)
}

func (i *InvestmentFundTransaction4) SetSettlementAmount(value, currency string) {
	i.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvestmentFundTransaction4) SetSettlementDate(value string) {
	i.SettlementDate = (*ISODate)(&value)
}

func (i *InvestmentFundTransaction4) AddTradeDateTime() *DateAndDateTimeChoice {
	i.TradeDateTime = new(DateAndDateTimeChoice)
	return i.TradeDateTime
}

func (i *InvestmentFundTransaction4) SetCumDividendIndicator(value string) {
	i.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentFundTransaction4) SetPartiallyExecutedIndicator(value string) {
	i.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentFundTransaction4) AddPriceDetails() *UnitPrice20 {
	i.PriceDetails = new(UnitPrice20)
	return i.PriceDetails
}
