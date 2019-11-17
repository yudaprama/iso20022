package model

// Provides the details of each individual un
// secured market transaction.
type UnsecuredMarketTransaction3 struct {

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

	// Date on which the amount of money is exchanged by counterparties or on which the purchase or sale of a debt instrument settles.
	// With regard to call accounts and other unsecured borrowing/lending redeemable at notice, it is the date on which the deposit is rolled over, that is on which it would have been paid back if it had been called/not rolled over. In the case of a settlement failure in which settlement takes place on a different date than initially agreed, no transactional amendment needs to be reported.
	SettlementDate *ISODate `xml:"SttlmDt"`

	// Date on which the amount of money is due to be repaid by the borrower to the lender or on which a debt instrument matures and is due to be paid back. In regards to callable and puttable instruments, the final maturity date must be provided. For call accounts and other unsecured borrowing/lending redeemable upon notice, the first date on which the instrument may be redeemed must be provided.
	MaturityDate *ISODate `xml:"MtrtyDt"`

	// Defines whether the transaction is a cash borrowing or cash lending transaction.
	TransactionType *MoneyMarketTransactionType1Code `xml:"TxTp"`

	// Defines the instrument via which the borrowing or lending transaction takes place.
	InstrumentType *FinancialInstrumentProductType1Code `xml:"InstrmTp"`

	// Amount of money initially borrowed or lent on deposits. In the case of debt securities, it is the nominal amount of the security issued or purchased.
	TransactionNominalAmount *ActiveCurrencyAndAmount `xml:"TxNmnlAmt"`

	// Dirty price at which the security is issued or traded in percentage points, and which is to be reported as 100 for unsecured deposits.
	DealPrice *PercentageRate `xml:"DealPric"`

	// Fixed rate for deposits and debt instruments with fixed coupons or variable rate for debt instruments for which the pay out at maturity or period depends on observed value of some underlying reference rate as well as for unsecured deposits paying interest at regular intervals.
	//
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

	// Debt instrument in which the periodic interest payments are calculated on the basis of the value (that is fixing of an underlying reference rate such as EURIBOR) on predefined dates (that is fixing) dates and which has a maturity of no more than one year.
	FloatingRateNote *FloatingRateNote2 `xml:"FltgRateNote,omitempty"`

	// Specifies whether the transaction is arranged via a third party broker or not.
	BrokeredDeal *BrokeredDeal1Code `xml:"BrkrdDeal,omitempty"`

	// Provides the option details, when the transaction reported is a call/put option.
	CallPutOption []*Option12 `xml:"CallPutOptn,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (u *UnsecuredMarketTransaction3) SetReportedTransactionStatus(value string) {
	u.ReportedTransactionStatus = (*TransactionOperationType1Code)(&value)
}

func (u *UnsecuredMarketTransaction3) SetBranchIdentification(value string) {
	u.BranchIdentification = (*LEIIdentifier)(&value)
}

func (u *UnsecuredMarketTransaction3) SetUniqueTransactionIdentifier(value string) {
	u.UniqueTransactionIdentifier = (*Max105Text)(&value)
}

func (u *UnsecuredMarketTransaction3) SetProprietaryTransactionIdentification(value string) {
	u.ProprietaryTransactionIdentification = (*Max105Text)(&value)
}

func (u *UnsecuredMarketTransaction3) SetCounterpartyProprietaryTransactionIdentification(value string) {
	u.CounterpartyProprietaryTransactionIdentification = (*Max105Text)(&value)
}

func (u *UnsecuredMarketTransaction3) AddCounterpartyIdentification() *CounterpartyIdentification2Choice {
	u.CounterpartyIdentification = new(CounterpartyIdentification2Choice)
	return u.CounterpartyIdentification
}

func (u *UnsecuredMarketTransaction3) AddTradeDate() *DateAndDateTimeChoice {
	u.TradeDate = new(DateAndDateTimeChoice)
	return u.TradeDate
}

func (u *UnsecuredMarketTransaction3) SetSettlementDate(value string) {
	u.SettlementDate = (*ISODate)(&value)
}

func (u *UnsecuredMarketTransaction3) SetMaturityDate(value string) {
	u.MaturityDate = (*ISODate)(&value)
}

func (u *UnsecuredMarketTransaction3) SetTransactionType(value string) {
	u.TransactionType = (*MoneyMarketTransactionType1Code)(&value)
}

func (u *UnsecuredMarketTransaction3) SetInstrumentType(value string) {
	u.InstrumentType = (*FinancialInstrumentProductType1Code)(&value)
}

func (u *UnsecuredMarketTransaction3) SetTransactionNominalAmount(value, currency string) {
	u.TransactionNominalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (u *UnsecuredMarketTransaction3) SetDealPrice(value string) {
	u.DealPrice = (*PercentageRate)(&value)
}

func (u *UnsecuredMarketTransaction3) SetRateType(value string) {
	u.RateType = (*InterestRateType1Code)(&value)
}

func (u *UnsecuredMarketTransaction3) AddDealRate() *Rate2 {
	u.DealRate = new(Rate2)
	return u.DealRate
}

func (u *UnsecuredMarketTransaction3) AddFloatingRateNote() *FloatingRateNote2 {
	u.FloatingRateNote = new(FloatingRateNote2)
	return u.FloatingRateNote
}

func (u *UnsecuredMarketTransaction3) SetBrokeredDeal(value string) {
	u.BrokeredDeal = (*BrokeredDeal1Code)(&value)
}

func (u *UnsecuredMarketTransaction3) AddCallPutOption() *Option12 {
	newValue := new(Option12)
	u.CallPutOption = append(u.CallPutOption, newValue)
	return newValue
}

func (u *UnsecuredMarketTransaction3) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	u.SupplementaryData = append(u.SupplementaryData, newValue)
	return newValue
}
