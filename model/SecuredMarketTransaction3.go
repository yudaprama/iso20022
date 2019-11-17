package model

// Provides the details of each individual secured market transaction.
type SecuredMarketTransaction3 struct {

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

	// Identification of the tri-party agent, when the transaction has been performed via tri-party agent.
	TripartyAgentIdentification *LEIIdentifier `xml:"TrptyAgtId,omitempty"`

	// Date and time on which the parties entered into the reported transaction.
	//
	// Usage: when time is available, it must be reported.
	//
	// It is to be reported with only the date when the time of the transaction is not available.
	//
	// The reported time is the execution time when available or otherwise the time at which the transaction entered the trading system of the reporting agent.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt"`

	// Date on which the amount of money is initially exchanged versus the asset as contractually agreed.
	//
	// Usage:
	// In the case of open basis repurchase transactions, this is the date on which the rollover settles (even if no exchange of an amount of money takes place).
	// In the case of a settlement failure in which settlement takes place on a date different than initially agreed, no transactional amendment needs to be reported.
	SettlementDate *ISODate `xml:"SttlmDt"`

	// Date on which the repurchase will be executed, that is the date on which the amount of money is due to be returned or received versus the asset pledged or received as collateral.
	MaturityDate *ISODate `xml:"MtrtyDt"`

	// Defines whether the transaction is a cash borrowing or cash lending transaction.
	TransactionType *MoneyMarketTransactionType1Code `xml:"TxTp"`

	// Amount of money initially borrowed or lent to be reported as an absolute value.
	TransactionNominalAmount *ActiveCurrencyAndAmount `xml:"TxNmnlAmt"`

	// Specifies whether the transaction interest rate of the repurchase agreements is either fixed or floating (variable rate).
	RateType *InterestRateType1Code `xml:"RateTp"`

	// Interest rate expressed in accordance with the local money market convention at which the repurchase agreement has been concluded and at which the cash lent is remunerated.
	//
	// Usage:
	// When the remuneration for securities lending transactions is represented by a fee amount, the fee amount will be translated into a deal rate per annum based on the ratio between the fee amount and the transaction nominal amount times number of days based on relevant money market convention divided by the number of days between the settlement date and the maturity of the transaction.
	//
	// Only actual values, as opposed to estimated or default values, will be reported for this variable.
	//
	// This value can be either positive or negative irrespective of whether the cash is borrowed or lent. It represents the contractually agreed remuneration rate on the transaction nominal amount regardless of the transaction sign (that whether the transaction type is reported as borrowed or lent).
	DealRate *Rate2 `xml:"DealRate,omitempty"`

	// Repurchase agreement in which the periodic interest payments are calculated on the basis of the value (that is, fixing of an underlying reference rate such as Euribor) on predefined dates (that is, fixing dates) and which has a maturity of no more than one year.
	FloatingRateRepurchaseAgreement *FloatingRateNote2 `xml:"FltgRateRpAgrmt,omitempty"`

	// Specifies whether the transaction is arranged via a third party broker or not.
	BrokeredDeal *BrokeredDeal1Code `xml:"BrkrdDeal,omitempty"`

	// Identifies the asset class pledged as collateral.
	Collateral *Collateral14 `xml:"Coll"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuredMarketTransaction3) SetReportedTransactionStatus(value string) {
	s.ReportedTransactionStatus = (*TransactionOperationType1Code)(&value)
}

func (s *SecuredMarketTransaction3) SetBranchIdentification(value string) {
	s.BranchIdentification = (*LEIIdentifier)(&value)
}

func (s *SecuredMarketTransaction3) SetUniqueTransactionIdentifier(value string) {
	s.UniqueTransactionIdentifier = (*Max105Text)(&value)
}

func (s *SecuredMarketTransaction3) SetProprietaryTransactionIdentification(value string) {
	s.ProprietaryTransactionIdentification = (*Max105Text)(&value)
}

func (s *SecuredMarketTransaction3) SetCounterpartyProprietaryTransactionIdentification(value string) {
	s.CounterpartyProprietaryTransactionIdentification = (*Max105Text)(&value)
}

func (s *SecuredMarketTransaction3) AddCounterpartyIdentification() *CounterpartyIdentification2Choice {
	s.CounterpartyIdentification = new(CounterpartyIdentification2Choice)
	return s.CounterpartyIdentification
}

func (s *SecuredMarketTransaction3) SetTripartyAgentIdentification(value string) {
	s.TripartyAgentIdentification = (*LEIIdentifier)(&value)
}

func (s *SecuredMarketTransaction3) AddTradeDate() *DateAndDateTimeChoice {
	s.TradeDate = new(DateAndDateTimeChoice)
	return s.TradeDate
}

func (s *SecuredMarketTransaction3) SetSettlementDate(value string) {
	s.SettlementDate = (*ISODate)(&value)
}

func (s *SecuredMarketTransaction3) SetMaturityDate(value string) {
	s.MaturityDate = (*ISODate)(&value)
}

func (s *SecuredMarketTransaction3) SetTransactionType(value string) {
	s.TransactionType = (*MoneyMarketTransactionType1Code)(&value)
}

func (s *SecuredMarketTransaction3) SetTransactionNominalAmount(value, currency string) {
	s.TransactionNominalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuredMarketTransaction3) SetRateType(value string) {
	s.RateType = (*InterestRateType1Code)(&value)
}

func (s *SecuredMarketTransaction3) AddDealRate() *Rate2 {
	s.DealRate = new(Rate2)
	return s.DealRate
}

func (s *SecuredMarketTransaction3) AddFloatingRateRepurchaseAgreement() *FloatingRateNote2 {
	s.FloatingRateRepurchaseAgreement = new(FloatingRateNote2)
	return s.FloatingRateRepurchaseAgreement
}

func (s *SecuredMarketTransaction3) SetBrokeredDeal(value string) {
	s.BrokeredDeal = (*BrokeredDeal1Code)(&value)
}

func (s *SecuredMarketTransaction3) AddCollateral() *Collateral14 {
	s.Collateral = new(Collateral14)
	return s.Collateral
}

func (s *SecuredMarketTransaction3) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
