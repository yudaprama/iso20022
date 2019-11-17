package model

// Provides the details of each individual overnight index swap transaction.
type OvernightIndexSwapTransaction3 struct {

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

	// Represents the date as of which the overnight rate of the floating leg is computed.
	StartDate *ISODate `xml:"StartDt"`

	// Last date of the term over which the compounded overnight rate is calculated.
	MaturityDate *ISODate `xml:"MtrtyDt"`

	// Fixed rate used for the calculation of the overnight index swap pay out.
	FixedInterestRate *Rate2 `xml:"FxdIntrstRate"`

	// Defines whether the fixed interest rate is paid or received by the reporting agent.
	TransactionType *OvernightIndexSwapType1Code `xml:"TxTp"`

	// Notional amount of the overnight index swap.
	TransactionNominalAmount *ActiveCurrencyAndAmount `xml:"TxNmnlAmt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (o *OvernightIndexSwapTransaction3) SetReportedTransactionStatus(value string) {
	o.ReportedTransactionStatus = (*TransactionOperationType1Code)(&value)
}

func (o *OvernightIndexSwapTransaction3) SetBranchIdentification(value string) {
	o.BranchIdentification = (*LEIIdentifier)(&value)
}

func (o *OvernightIndexSwapTransaction3) SetUniqueTransactionIdentifier(value string) {
	o.UniqueTransactionIdentifier = (*Max105Text)(&value)
}

func (o *OvernightIndexSwapTransaction3) SetProprietaryTransactionIdentification(value string) {
	o.ProprietaryTransactionIdentification = (*Max105Text)(&value)
}

func (o *OvernightIndexSwapTransaction3) SetCounterpartyProprietaryTransactionIdentification(value string) {
	o.CounterpartyProprietaryTransactionIdentification = (*Max105Text)(&value)
}

func (o *OvernightIndexSwapTransaction3) AddCounterpartyIdentification() *CounterpartyIdentification2Choice {
	o.CounterpartyIdentification = new(CounterpartyIdentification2Choice)
	return o.CounterpartyIdentification
}

func (o *OvernightIndexSwapTransaction3) AddTradeDate() *DateAndDateTimeChoice {
	o.TradeDate = new(DateAndDateTimeChoice)
	return o.TradeDate
}

func (o *OvernightIndexSwapTransaction3) SetStartDate(value string) {
	o.StartDate = (*ISODate)(&value)
}

func (o *OvernightIndexSwapTransaction3) SetMaturityDate(value string) {
	o.MaturityDate = (*ISODate)(&value)
}

func (o *OvernightIndexSwapTransaction3) AddFixedInterestRate() *Rate2 {
	o.FixedInterestRate = new(Rate2)
	return o.FixedInterestRate
}

func (o *OvernightIndexSwapTransaction3) SetTransactionType(value string) {
	o.TransactionType = (*OvernightIndexSwapType1Code)(&value)
}

func (o *OvernightIndexSwapTransaction3) SetTransactionNominalAmount(value, currency string) {
	o.TransactionNominalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (o *OvernightIndexSwapTransaction3) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	o.SupplementaryData = append(o.SupplementaryData, newValue)
	return newValue
}
