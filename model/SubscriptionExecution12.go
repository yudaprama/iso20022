package model

// Execution of a subscription order.
type SubscriptionExecution12 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order execution, as assigned by the confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Category of the investment fund order.
	OrderType []*FundOrderType4Choice `xml:"OrdrTp,omitempty"`

	// Account impacted by the investment fund order execution.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails []*IndividualPerson31 `xml:"BnfcryDtls,omitempty"`

	// Number of investment fund units subscribed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money invested in the fund.
	// Net Amount = Quantity * Price.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Amount of money to be paid by the investor when subscribing to fund units.
	// Gross amount = (Quantity * Price) + (Fees + Taxes).
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	DealingPriceDetails *UnitPrice22 `xml:"DealgPricDtls"`

	// Other quoted price than the one at which the order was executed.
	InformativePriceDetails []*UnitPrice22 `xml:"InftvPricDtls,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Indicates whether the order has been partially executed, that is, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Specifies that the execution was subject to best execution rules as defined by MiFID.
	BestExecution *BestExecution1Code `xml:"BestExctn,omitempty"`

	// Indicates whether the dividend is included, that is, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss2Choice `xml:"IntrmPrftAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	// How the exchange rate is expressed determines which currency is the Unit Currency and Quoted Currency. If the amounts concerned are EUR 1000 and USD 1300, the exchange rate may be expressed as per either of the following examples:
	// EXAMPLE 1
	// UnitCurrency  EUR
	// QuotedCurrency  USD
	// ExchangeRate  1.300
	// EXAMPLE 2
	// UnitCurrency  USD
	// QuotedCurrency  EUR
	// ExchangeRate  0.769
	ForeignExchangeDetails []*ForeignExchangeTerms33 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Fees (charges/commission) and taxes that are taken into consideration for the transaction, so that the total difference between the net amount and gross amount is known, without taking into account equalisation.
	TransactionOverhead *TotalFeesAndTaxes40 `xml:"TxOvrhd,omitempty"`

	// Additional information about tax that does not have an impact on the transaction overhead.
	InformativeTaxDetails *InformativeTax1 `xml:"InftvTaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters12 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown2 `xml:"StffClntBrkdwn,omitempty"`

	// Return of cash that has been overpaid for a subscription.
	Refund *ActiveCurrencyAndAmount `xml:"Rfnd,omitempty"`

	// Interest received when a subscription amount is paid in advance and then invested by the fund.
	SubscriptionInterest *ActiveCurrencyAndAmount `xml:"SbcptIntrst,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction70 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Percentage of units partially settled.
	PartialSettlementOfUnits *PercentageRate `xml:"PrtlSttlmOfUnits,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`

	// Percentage of cash partially settled.
	PartialSettlementOfCash *PercentageRate `xml:"PrtlSttlmOfCsh,omitempty"`

	// Party related to the transaction.
	RelatedPartyDetails []*Intermediary39 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`

	// Source of cash used for the settlement of the subscription.
	SourceOfCash []*SourceOfCash1Choice `xml:"SrcOfCsh,omitempty"`

	// Assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Type of signature.
	SignatureType *SignatureType1Choice `xml:"SgntrTp,omitempty"`

	// Information about a non-standard order.
	OrderWaiverDetails *OrderWaiver1 `xml:"OrdrWvrDtls,omitempty"`
}

func (s *SubscriptionExecution12) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution12) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution12) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution12) AddOrderType() *FundOrderType4Choice {
	newValue := new(FundOrderType4Choice)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionExecution12) AddInvestmentAccountDetails() *InvestmentAccount58 {
	s.InvestmentAccountDetails = new(InvestmentAccount58)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionExecution12) AddBeneficiaryDetails() *IndividualPerson31 {
	newValue := new(IndividualPerson31)
	s.BeneficiaryDetails = append(s.BeneficiaryDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution12) SetUnitsNumber(value string) {
	s.UnitsNumber = (*DecimalNumber)(&value)
}

func (s *SubscriptionExecution12) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionExecution12) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution12) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution12) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SubscriptionExecution12) AddDealingPriceDetails() *UnitPrice22 {
	s.DealingPriceDetails = new(UnitPrice22)
	return s.DealingPriceDetails
}

func (s *SubscriptionExecution12) AddInformativePriceDetails() *UnitPrice22 {
	newValue := new(UnitPrice22)
	s.InformativePriceDetails = append(s.InformativePriceDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution12) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution12) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionExecution12) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SubscriptionExecution12) SetPartiallyExecutedIndicator(value string) {
	s.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution12) SetBestExecution(value string) {
	s.BestExecution = (*BestExecution1Code)(&value)
}

func (s *SubscriptionExecution12) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution12) AddInterimProfitAmount() *ProfitAndLoss2Choice {
	s.InterimProfitAmount = new(ProfitAndLoss2Choice)
	return s.InterimProfitAmount
}

func (s *SubscriptionExecution12) AddForeignExchangeDetails() *ForeignExchangeTerms33 {
	newValue := new(ForeignExchangeTerms33)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution12) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionExecution12) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution12) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution12) AddTransactionOverhead() *TotalFeesAndTaxes40 {
	s.TransactionOverhead = new(TotalFeesAndTaxes40)
	return s.TransactionOverhead
}

func (s *SubscriptionExecution12) AddInformativeTaxDetails() *InformativeTax1 {
	s.InformativeTaxDetails = new(InformativeTax1)
	return s.InformativeTaxDetails
}

func (s *SubscriptionExecution12) AddSettlementAndCustodyDetails() *FundSettlementParameters12 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters12)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionExecution12) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution12) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionExecution12) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown2 {
	newValue := new(InvestmentFundsOrderBreakdown2)
	s.StaffClientBreakdown = append(s.StaffClientBreakdown, newValue)
	return newValue
}

func (s *SubscriptionExecution12) SetRefund(value, currency string) {
	s.Refund = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution12) SetSubscriptionInterest(value, currency string) {
	s.SubscriptionInterest = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution12) AddCashSettlementDetails() *PaymentTransaction70 {
	s.CashSettlementDetails = new(PaymentTransaction70)
	return s.CashSettlementDetails
}

func (s *SubscriptionExecution12) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SubscriptionExecution12) SetPartialSettlementOfUnits(value string) {
	s.PartialSettlementOfUnits = (*PercentageRate)(&value)
}

func (s *SubscriptionExecution12) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SubscriptionExecution12) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SubscriptionExecution12) SetLateReport(value string) {
	s.LateReport = (*LateReport1Code)(&value)
}

func (s *SubscriptionExecution12) SetPartialSettlementOfCash(value string) {
	s.PartialSettlementOfCash = (*PercentageRate)(&value)
}

func (s *SubscriptionExecution12) AddRelatedPartyDetails() *Intermediary39 {
	newValue := new(Intermediary39)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution12) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

func (s *SubscriptionExecution12) AddSourceOfCash() *SourceOfCash1Choice {
	newValue := new(SourceOfCash1Choice)
	s.SourceOfCash = append(s.SourceOfCash, newValue)
	return newValue
}

func (s *SubscriptionExecution12) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	s.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return s.CustomerConductClassification
}

func (s *SubscriptionExecution12) AddTransactionChannelType() *TransactionChannelType1Choice {
	s.TransactionChannelType = new(TransactionChannelType1Choice)
	return s.TransactionChannelType
}

func (s *SubscriptionExecution12) AddSignatureType() *SignatureType1Choice {
	s.SignatureType = new(SignatureType1Choice)
	return s.SignatureType
}

func (s *SubscriptionExecution12) AddOrderWaiverDetails() *OrderWaiver1 {
	s.OrderWaiverDetails = new(OrderWaiver1)
	return s.OrderWaiverDetails
}
