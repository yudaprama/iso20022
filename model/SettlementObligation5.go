package model

// Provides details about the settlement obligation.
type SettlementObligation5 struct {

	// Provides the identification of an existing obligation that is linked to the new obligation.
	RelatedSettlementObligationIdentification *Max35Text `xml:"RltdSttlmOblgtnId,omitempty"`

	// Indicates the type of the obligation.
	ObligationType *ObligationType1Choice `xml:"OblgtnTp,omitempty"`

	// Provides additional information related to the linked obligation.
	Description *Max35Text `xml:"Desc,omitempty"`

	// Provides the original trade date.
	TradeDate *ISODate `xml:"TradDt,omitempty"`

	// Specifies the quantity related to the settlement obligation.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Provides the price applied to that net position.
	NetPositionPrice *Price4 `xml:"NetPosPric,omitempty"`

	// Specifies the ISO code of the trade currency.
	TradingCurrency *CurrencyCode `xml:"TradgCcy,omitempty"`

	// Provides the total amount to be settled.
	SettlementAmount *AmountAndDirection27 `xml:"SttlmAmt"`

	// Provides the contractual settlement date.
	SettlementDate *ISODate `xml:"SttlmDt"`

	// Indicates if the obligation will result in a receive or a delivery of securities.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Provides the references of the underlying trade leg(s) and/or the reference to the related net position message.
	References *Reference19 `xml:"Refs,omitempty"`
}

func (s *SettlementObligation5) SetRelatedSettlementObligationIdentification(value string) {
	s.RelatedSettlementObligationIdentification = (*Max35Text)(&value)
}

func (s *SettlementObligation5) AddObligationType() *ObligationType1Choice {
	s.ObligationType = new(ObligationType1Choice)
	return s.ObligationType
}

func (s *SettlementObligation5) SetDescription(value string) {
	s.Description = (*Max35Text)(&value)
}

func (s *SettlementObligation5) SetTradeDate(value string) {
	s.TradeDate = (*ISODate)(&value)
}

func (s *SettlementObligation5) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SettlementObligation5) AddNetPositionPrice() *Price4 {
	s.NetPositionPrice = new(Price4)
	return s.NetPositionPrice
}

func (s *SettlementObligation5) SetTradingCurrency(value string) {
	s.TradingCurrency = (*CurrencyCode)(&value)
}

func (s *SettlementObligation5) AddSettlementAmount() *AmountAndDirection27 {
	s.SettlementAmount = new(AmountAndDirection27)
	return s.SettlementAmount
}

func (s *SettlementObligation5) SetSettlementDate(value string) {
	s.SettlementDate = (*ISODate)(&value)
}

func (s *SettlementObligation5) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementObligation5) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementObligation5) AddReferences() *Reference19 {
	s.References = new(Reference19)
	return s.References
}
