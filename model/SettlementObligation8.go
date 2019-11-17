package model

// Provides details about the settlement obligation.
type SettlementObligation8 struct {

	// Provides the identification of the settlement obligation.
	SettlementObligationIdentification *Max35Text `xml:"SttlmOblgtnId"`

	// Provides details about the security identification.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Intended settlement date of the position.
	IntendedSettlementDate *DateFormat11Choice `xml:"IntnddSttlmDt"`

	// Specifies the quantity related to the settlement obligation.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Provides the total amount to be settled.
	SettlementAmount *AmountAndDirection27 `xml:"SttlmAmt"`

	// Place at which the security is traded.
	PlaceOfTrade *MarketIdentification84 `xml:"PlcOfTrad"`

	// Provides the trade date.
	TradeDate *TradeDate3Choice `xml:"TradDt,omitempty"`

	// Identifies the trading capacity of seller.
	TradingCapacity *TradingCapacity5Code `xml:"TradgCpcty,omitempty"`

	// Specifies the clearing account type.
	ClearingAccountType *ClearingAccountType1Code `xml:"ClrAcctTp,omitempty"`

	// Place where the securities are safe-kept, physically or notionally. This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat7Choice `xml:"SfkpgPlc,omitempty"`

	// Clearing member account at the central securities depository.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Indicates if the obligation will result in a receive or a delivery of securities.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp,omitempty"`

	// Specifies the amount to be paid under a specific obligation.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Provides details on the parties that are part of the settlement chain.
	SettlementParties *SettlementParties4Choice `xml:"SttlmPties,omitempty"`

	// Provides additional information about the settlement obligation details.
	AdditionalSettlementObligationDetails []*SettlementObligation5 `xml:"AddtlSttlmOblgtnDtls,omitempty"`
}

func (s *SettlementObligation8) SetSettlementObligationIdentification(value string) {
	s.SettlementObligationIdentification = (*Max35Text)(&value)
}

func (s *SettlementObligation8) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SettlementObligation8) AddIntendedSettlementDate() *DateFormat11Choice {
	s.IntendedSettlementDate = new(DateFormat11Choice)
	return s.IntendedSettlementDate
}

func (s *SettlementObligation8) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SettlementObligation8) AddSettlementAmount() *AmountAndDirection27 {
	s.SettlementAmount = new(AmountAndDirection27)
	return s.SettlementAmount
}

func (s *SettlementObligation8) AddPlaceOfTrade() *MarketIdentification84 {
	s.PlaceOfTrade = new(MarketIdentification84)
	return s.PlaceOfTrade
}

func (s *SettlementObligation8) AddTradeDate() *TradeDate3Choice {
	s.TradeDate = new(TradeDate3Choice)
	return s.TradeDate
}

func (s *SettlementObligation8) SetTradingCapacity(value string) {
	s.TradingCapacity = (*TradingCapacity5Code)(&value)
}

func (s *SettlementObligation8) SetClearingAccountType(value string) {
	s.ClearingAccountType = (*ClearingAccountType1Code)(&value)
}

func (s *SettlementObligation8) AddSafekeepingPlace() *SafekeepingPlaceFormat7Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat7Choice)
	return s.SafekeepingPlace
}

func (s *SettlementObligation8) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SettlementObligation8) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementObligation8) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementObligation8) AddSettlementParties() *SettlementParties4Choice {
	s.SettlementParties = new(SettlementParties4Choice)
	return s.SettlementParties
}

func (s *SettlementObligation8) AddAdditionalSettlementObligationDetails() *SettlementObligation5 {
	newValue := new(SettlementObligation5)
	s.AdditionalSettlementObligationDetails = append(s.AdditionalSettlementObligationDetails, newValue)
	return newValue
}
