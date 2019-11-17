package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03500207 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.035.002.07 Document"`
	Message *SecuritiesFinancingConfirmation002V07 `xml:"SctiesFincgConf"`
}

func (d *Document03500207) AddMessage() *SecuritiesFinancingConfirmation002V07 {
	d.Message = new(SecuritiesFinancingConfirmation002V07)
	return d.Message
}

// Scope
// A securities financing transaction account servicer sends a SecuritiesFinancingConfirmation to an account owner to confirm or advise of the partial or full settlement of the opening or closing leg of a securities financing transaction.
//
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure managing securities financing transactions on behalf of their participants
// - an agent (sub-custodian) managing securities financing transactions on behalf of their global custodian customer, or
// - a custodian managing securities financing transactions on behalf of an investment management institution or a broker/dealer.
//
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesFinancingConfirmation002V07 struct {

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionIdentificationDetails *model.TransactionTypeAndAdditionalParameters19 `xml:"TxIdDtls"`

	// Additional parameters to the transaction.
	AdditionalParameters *model.AdditionalParameters26 `xml:"AddtlParams,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *model.SecuritiesTradeDetails58 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount62 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *model.SecuritiesFinancingTransactionDetails30 `xml:"SctiesFincgDtls,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails104 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties30 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *model.AmountAndDirection60 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts34 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingConfirmation002V07) AddTransactionIdentificationDetails() *model.TransactionTypeAndAdditionalParameters19 {
	s.TransactionIdentificationDetails = new(model.TransactionTypeAndAdditionalParameters19)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesFinancingConfirmation002V07) AddAdditionalParameters() *model.AdditionalParameters26 {
	s.AdditionalParameters = new(model.AdditionalParameters26)
	return s.AdditionalParameters
}

func (s *SecuritiesFinancingConfirmation002V07) AddTradeDetails() *model.SecuritiesTradeDetails58 {
	s.TradeDetails = new(model.SecuritiesTradeDetails58)
	return s.TradeDetails
}

func (s *SecuritiesFinancingConfirmation002V07) AddFinancialInstrumentIdentification() *model.SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingConfirmation002V07) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingConfirmation002V07) AddQuantityAndAccountDetails() *model.QuantityAndAccount62 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount62)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingConfirmation002V07) AddSecuritiesFinancingDetails() *model.SecuritiesFinancingTransactionDetails30 {
	s.SecuritiesFinancingDetails = new(model.SecuritiesFinancingTransactionDetails30)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingConfirmation002V07) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingConfirmation002V07) AddSettlementParameters() *model.SettlementDetails104 {
	s.SettlementParameters = new(model.SettlementDetails104)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingConfirmation002V07) AddDeliveringSettlementParties() *model.SettlementParties44 {
	s.DeliveringSettlementParties = new(model.SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingConfirmation002V07) AddReceivingSettlementParties() *model.SettlementParties44 {
	s.ReceivingSettlementParties = new(model.SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingConfirmation002V07) AddCashParties() *model.CashParties30 {
	s.CashParties = new(model.CashParties30)
	return s.CashParties
}

func (s *SecuritiesFinancingConfirmation002V07) AddSettledAmount() *model.AmountAndDirection60 {
	s.SettledAmount = new(model.AmountAndDirection60)
	return s.SettledAmount
}

func (s *SecuritiesFinancingConfirmation002V07) AddOtherAmounts() *model.OtherAmounts34 {
	s.OtherAmounts = new(model.OtherAmounts34)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingConfirmation002V07) AddOtherBusinessParties() *model.OtherParties29 {
	s.OtherBusinessParties = new(model.OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingConfirmation002V07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
