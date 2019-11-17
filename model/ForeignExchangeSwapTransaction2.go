package model

// Provides the details of each individual foreign exchange swap transaction.
type ForeignExchangeSwapTransaction2 struct {

	// Defines the status of the reported transaction, that is details on whether the transaction is a new transaction, an amendment of a previously reported transaction, a cancellation of a previously reported transaction or a correction to a previously reported and rejected transaction.
	ReportedTransactionStatus *TransactionOperationType1Code `xml:"RptdTxSts"`

	// Unique and unambiguous legal entity identification of  the branch of the reporting agent in which the transaction has been booked.
	//
	// Usage: This field must only be provided if the transaction has been conducted and booked by a branch of the reporting agent and only if this branch has its own LEI that the reporting agent can clearly identify.
	// Where the transaction has been booked by the head office or the reporting agent cannot be identified by a unique branch-specific LEI, the reporting agent must provide the LEI of the head office.
	BranchIdentification *LEIIdentifier `xml:"BrnchId,omitempty"`

	// Unique transaction identifier will be created at the time a transaction is first executed, shared with all registered entities and counterparties involved in the transaction, and used to track that particular transaction during its lifetime.
	UniqueTransactionIdentifier *Max105Text `xml:"UnqTxIdr,omitempty"`

	// Internal unique transaction identifier used by the reporting agent for each transaction.
	ProprietaryTransactionIdentification *Max105Text `xml:"PrtryTxId"`

	// Internal unique proprietary transaction identifier as assigned by the counterparty of the reporting agent for each transaction.
	CounterpartyProprietaryTransactionIdentification *Max105Text `xml:"CtrPtyPrtryTxId,omitempty"`

	// Identification of the counterparty of the reporting agent for the reported transaction.
	CounterpartyIdentification *CounterpartyIdentification2Choice `xml:"CtrPtyId"`

	// Date and time on which the parties entered into the reported transaction.
	//
	// Usage: when time is available, it must be reported.
	//
	// It is to be reported with only the date when the time of the transaction is not available.
	//
	// The reported time is the execution time when available or otherwise the time at which the transaction entered the trading system of the reporting agent.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt"`

	// Date on which one party sells to the other a specified amount of a specified currency against payment of an agreed amount of a specified different currency based on an agreed foreign exchange rate known as foreign exchange spot rate.
	SpotValueDate *ISODate `xml:"SpotValDt"`

	// Date on which the foreign exchange swap transaction expires and the currency sold on the value date is repurchased.
	MaturityDate *ISODate `xml:"MtrtyDt"`

	// Defines whether the amount of money reported under the transaction nominal amount is bought or sold on the spot value date.
	TransactionType *SecuritiesTransactionType15Code `xml:"TxTp"`

	// Specifies the nominal amount of the foreign exchange swap, that is the amount bought/sold on the spot value date and reported as an absolute value.
	TransactionNominalAmount *ActiveCurrencyAndAmount `xml:"TxNmnlAmt"`

	// Provides the details of the foreign exchange transaction.
	ForeignExchange *ForeignExchange1 `xml:"FX"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *ForeignExchangeSwapTransaction2) SetReportedTransactionStatus(value string) {
	f.ReportedTransactionStatus = (*TransactionOperationType1Code)(&value)
}

func (f *ForeignExchangeSwapTransaction2) SetBranchIdentification(value string) {
	f.BranchIdentification = (*LEIIdentifier)(&value)
}

func (f *ForeignExchangeSwapTransaction2) SetUniqueTransactionIdentifier(value string) {
	f.UniqueTransactionIdentifier = (*Max105Text)(&value)
}

func (f *ForeignExchangeSwapTransaction2) SetProprietaryTransactionIdentification(value string) {
	f.ProprietaryTransactionIdentification = (*Max105Text)(&value)
}

func (f *ForeignExchangeSwapTransaction2) SetCounterpartyProprietaryTransactionIdentification(value string) {
	f.CounterpartyProprietaryTransactionIdentification = (*Max105Text)(&value)
}

func (f *ForeignExchangeSwapTransaction2) AddCounterpartyIdentification() *CounterpartyIdentification2Choice {
	f.CounterpartyIdentification = new(CounterpartyIdentification2Choice)
	return f.CounterpartyIdentification
}

func (f *ForeignExchangeSwapTransaction2) AddTradeDate() *DateAndDateTimeChoice {
	f.TradeDate = new(DateAndDateTimeChoice)
	return f.TradeDate
}

func (f *ForeignExchangeSwapTransaction2) SetSpotValueDate(value string) {
	f.SpotValueDate = (*ISODate)(&value)
}

func (f *ForeignExchangeSwapTransaction2) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *ForeignExchangeSwapTransaction2) SetTransactionType(value string) {
	f.TransactionType = (*SecuritiesTransactionType15Code)(&value)
}

func (f *ForeignExchangeSwapTransaction2) SetTransactionNominalAmount(value, currency string) {
	f.TransactionNominalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *ForeignExchangeSwapTransaction2) AddForeignExchange() *ForeignExchange1 {
	f.ForeignExchange = new(ForeignExchange1)
	return f.ForeignExchange
}

func (f *ForeignExchangeSwapTransaction2) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
