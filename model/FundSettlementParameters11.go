package model

// Parameters applied to the settlement of a security.
type FundSettlementParameters11 struct {

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Place where the settlement of the transaction will take place. In the context of investment funds, the place of settlement is the transfer agent, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SettlementPlace *PartyIdentification113 `xml:"SttlmPlc"`

	// Place where the securities are safe-kept, physically or notionally. This place can be, for example, a local custodian, a Central Securities Depository or an International Central Securities Depository.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Identification of a specific system or set of rules and/or processes to be applied at the settlement place.
	SecuritiesSettlementSystemIdentification *Max35Text `xml:"SctiesSttlmSysId,omitempty"`

	// Condition under which the order/trade is to be/was executed. This may be required for settlement through T2S.
	TradeTransactionCondition []*TradeTransactionCondition8Choice `xml:"TradTxCond,omitempty"`

	// Condition under which the order/trade is to be settled. This may be required for settlement through T2S.
	SettlementTransactionCondition []*SettlementTransactionCondition30Choice `xml:"SttlmTxCond,omitempty"`

	// Chain of parties involved in the settlement of a transaction resulting in the movement of a security from one account to another.
	ReceivingSideDetails *ReceivingPartiesAndAccount16 `xml:"RcvgSdDtls"`

	// Chain of parties involved in the settlement of a transaction resulting in the movement of a security from one account to another.
	DeliveringSideDetails *DeliveringPartiesAndAccount16 `xml:"DlvrgSdDtls,omitempty"`
}

func (f *FundSettlementParameters11) SetSettlementDate(value string) {
	f.SettlementDate = (*ISODate)(&value)
}

func (f *FundSettlementParameters11) AddSettlementPlace() *PartyIdentification113 {
	f.SettlementPlace = new(PartyIdentification113)
	return f.SettlementPlace
}

func (f *FundSettlementParameters11) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	f.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return f.SafekeepingPlace
}

func (f *FundSettlementParameters11) SetSecuritiesSettlementSystemIdentification(value string) {
	f.SecuritiesSettlementSystemIdentification = (*Max35Text)(&value)
}

func (f *FundSettlementParameters11) AddTradeTransactionCondition() *TradeTransactionCondition8Choice {
	newValue := new(TradeTransactionCondition8Choice)
	f.TradeTransactionCondition = append(f.TradeTransactionCondition, newValue)
	return newValue
}

func (f *FundSettlementParameters11) AddSettlementTransactionCondition() *SettlementTransactionCondition30Choice {
	newValue := new(SettlementTransactionCondition30Choice)
	f.SettlementTransactionCondition = append(f.SettlementTransactionCondition, newValue)
	return newValue
}

func (f *FundSettlementParameters11) AddReceivingSideDetails() *ReceivingPartiesAndAccount16 {
	f.ReceivingSideDetails = new(ReceivingPartiesAndAccount16)
	return f.ReceivingSideDetails
}

func (f *FundSettlementParameters11) AddDeliveringSideDetails() *DeliveringPartiesAndAccount16 {
	f.DeliveringSideDetails = new(DeliveringPartiesAndAccount16)
	return f.DeliveringSideDetails
}
