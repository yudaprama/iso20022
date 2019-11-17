package model

// Provides details about the settlement obligation.
type SettlementObligation7 struct {

	// Last reference given by the settlement platform (this is the central securities depository)  to the transaction (non settled instruction).
	CSDTransactionIdentification *Max35Text `xml:"CSDTxId,omitempty"`

	// Reference of the transaction (non settled instruction) given by the central counterparty.
	CentralCounterpartyTransactionIdentification *Max35Text `xml:"CntrlCtrPtyTxId,omitempty"`

	// Original buy-in identification number in case an event causes a generation of a new buy-in identification.
	PreviousBuyInIdentification *Max35Text `xml:"PrvsBuyInId,omitempty"`

	// An account opened by the central counterparty in the name of the clearing member within the account structure, for settlement purposes (gives information about the clearing member account at central counterparty).
	DeliveryAccount *SecuritiesAccount19 `xml:"DlvryAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally. This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat7Choice `xml:"SfkpgPlc,omitempty"`

	// Clearing member account at the central securities depository.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Clearing organisation that will clear the trade.
	//
	// Note: This field allows clearing member firm to segregate flows coming from clearing counterparty's clearing system. Indeed, clearing member firms receive messages from the same system (same sender) and this field allows them to know if the message is related to equities or derivatives.
	ClearingSegment *PartyIdentification35Choice `xml:"ClrSgmt,omitempty"`

	// Provides the identification for the non-clearing member and account.
	NonClearingMember *PartyIdentificationAndAccount31 `xml:"NonClrMmb,omitempty"`

	// Provides the intended settlement date of the position.
	IntendedSettlementDate *ISODate `xml:"IntnddSttlmDt,omitempty"`

	// Provides details about the security identification.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Provides the trade date.
	TradeDate *ISODate `xml:"TradDt,omitempty"`

	// Provides the price of the trade.
	DealPrice *Price4 `xml:"DealPric,omitempty"`

	// Provides the quantity of the trade.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Place where settlement of the securities takes place.
	Depository *PartyIdentification34Choice `xml:"Dpstry,omitempty"`

	// Provides the remaining quantity to be settled.
	RemainingQuantityToBeSettled *FinancialInstrumentQuantity1Choice `xml:"RmngQtyToBeSttld,omitempty"`

	// Provides the amount to be settled.
	SettlementAmount *AmountAndDirection27 `xml:"SttlmAmt"`

	// Provides the remaining amount to be settled.
	RemainingAmountToBeSettled *AmountAndDirection27 `xml:"RmngAmtToBeSttld,omitempty"`
}

func (s *SettlementObligation7) SetCSDTransactionIdentification(value string) {
	s.CSDTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementObligation7) SetCentralCounterpartyTransactionIdentification(value string) {
	s.CentralCounterpartyTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementObligation7) SetPreviousBuyInIdentification(value string) {
	s.PreviousBuyInIdentification = (*Max35Text)(&value)
}

func (s *SettlementObligation7) AddDeliveryAccount() *SecuritiesAccount19 {
	s.DeliveryAccount = new(SecuritiesAccount19)
	return s.DeliveryAccount
}

func (s *SettlementObligation7) AddSafekeepingPlace() *SafekeepingPlaceFormat7Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat7Choice)
	return s.SafekeepingPlace
}

func (s *SettlementObligation7) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SettlementObligation7) AddClearingSegment() *PartyIdentification35Choice {
	s.ClearingSegment = new(PartyIdentification35Choice)
	return s.ClearingSegment
}

func (s *SettlementObligation7) AddNonClearingMember() *PartyIdentificationAndAccount31 {
	s.NonClearingMember = new(PartyIdentificationAndAccount31)
	return s.NonClearingMember
}

func (s *SettlementObligation7) SetIntendedSettlementDate(value string) {
	s.IntendedSettlementDate = (*ISODate)(&value)
}

func (s *SettlementObligation7) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SettlementObligation7) SetTradeDate(value string) {
	s.TradeDate = (*ISODate)(&value)
}

func (s *SettlementObligation7) AddDealPrice() *Price4 {
	s.DealPrice = new(Price4)
	return s.DealPrice
}

func (s *SettlementObligation7) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SettlementObligation7) AddDepository() *PartyIdentification34Choice {
	s.Depository = new(PartyIdentification34Choice)
	return s.Depository
}

func (s *SettlementObligation7) AddRemainingQuantityToBeSettled() *FinancialInstrumentQuantity1Choice {
	s.RemainingQuantityToBeSettled = new(FinancialInstrumentQuantity1Choice)
	return s.RemainingQuantityToBeSettled
}

func (s *SettlementObligation7) AddSettlementAmount() *AmountAndDirection27 {
	s.SettlementAmount = new(AmountAndDirection27)
	return s.SettlementAmount
}

func (s *SettlementObligation7) AddRemainingAmountToBeSettled() *AmountAndDirection27 {
	s.RemainingAmountToBeSettled = new(AmountAndDirection27)
	return s.RemainingAmountToBeSettled
}
