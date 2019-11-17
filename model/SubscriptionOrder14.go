package model

// Order to invest the investor's principal in an investment fund.
type SubscriptionOrder14 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Category of the investment fund order.
	OrderType []*FundOrderType4Choice `xml:"OrdrTp,omitempty"`

	// Investment fund class related to the order.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls"`

	// Subdivision of the account used to segregate specific holdings.
	SubAccountForHolding *SubAccount6 `xml:"SubAcctForHldg,omitempty"`

	// Amount of money or the number of units for the subscription order.
	AmountOrUnits *FinancialInstrumentQuantity27Choice `xml:"AmtOrUnits"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

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
	ForeignExchangeDetails *ForeignExchangeTerms32 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Fees (charges/commission) and tax to be applied to the net amount.
	TransactionOverhead *FeeAndTax1 `xml:"TxOvrhd,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters11 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction70 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown2 `xml:"StffClntBrkdwn,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Party related to the transaction.
	RelatedPartyDetails []*Intermediary40 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive/performance fees at the end of the fiscal year.
	//
	//
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

func (s *SubscriptionOrder14) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder14) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder14) AddOrderType() *FundOrderType4Choice {
	newValue := new(FundOrderType4Choice)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionOrder14) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	s.FinancialInstrumentDetails = new(FinancialInstrument57)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionOrder14) AddSubAccountForHolding() *SubAccount6 {
	s.SubAccountForHolding = new(SubAccount6)
	return s.SubAccountForHolding
}

func (s *SubscriptionOrder14) AddAmountOrUnits() *FinancialInstrumentQuantity27Choice {
	s.AmountOrUnits = new(FinancialInstrumentQuantity27Choice)
	return s.AmountOrUnits
}

func (s *SubscriptionOrder14) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionOrder14) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder14) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionOrder14) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SubscriptionOrder14) AddForeignExchangeDetails() *ForeignExchangeTerms32 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms32)
	return s.ForeignExchangeDetails
}

func (s *SubscriptionOrder14) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionOrder14) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder14) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder14) AddTransactionOverhead() *FeeAndTax1 {
	s.TransactionOverhead = new(FeeAndTax1)
	return s.TransactionOverhead
}

func (s *SubscriptionOrder14) AddSettlementAndCustodyDetails() *FundSettlementParameters11 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters11)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionOrder14) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionOrder14) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionOrder14) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SubscriptionOrder14) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SubscriptionOrder14) AddCashSettlementDetails() *PaymentTransaction70 {
	s.CashSettlementDetails = new(PaymentTransaction70)
	return s.CashSettlementDetails
}

func (s *SubscriptionOrder14) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SubscriptionOrder14) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown2 {
	newValue := new(InvestmentFundsOrderBreakdown2)
	s.StaffClientBreakdown = append(s.StaffClientBreakdown, newValue)
	return newValue
}

func (s *SubscriptionOrder14) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SubscriptionOrder14) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SubscriptionOrder14) AddRelatedPartyDetails() *Intermediary40 {
	newValue := new(Intermediary40)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder14) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

func (s *SubscriptionOrder14) AddSourceOfCash() *SourceOfCash1Choice {
	newValue := new(SourceOfCash1Choice)
	s.SourceOfCash = append(s.SourceOfCash, newValue)
	return newValue
}

func (s *SubscriptionOrder14) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	s.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return s.CustomerConductClassification
}

func (s *SubscriptionOrder14) AddTransactionChannelType() *TransactionChannelType1Choice {
	s.TransactionChannelType = new(TransactionChannelType1Choice)
	return s.TransactionChannelType
}

func (s *SubscriptionOrder14) AddSignatureType() *SignatureType1Choice {
	s.SignatureType = new(SignatureType1Choice)
	return s.SignatureType
}

func (s *SubscriptionOrder14) AddOrderWaiverDetails() *OrderWaiver1 {
	s.OrderWaiverDetails = new(OrderWaiver1)
	return s.OrderWaiverDetails
}
