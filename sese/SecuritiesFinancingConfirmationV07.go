package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03500107 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.035.001.07 Document"`
	Message *SecuritiesFinancingConfirmationV07 `xml:"SctiesFincgConf"`
}

func (d *Document03500107) AddMessage() *SecuritiesFinancingConfirmationV07 {
	d.Message = new(SecuritiesFinancingConfirmationV07)
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
type SecuritiesFinancingConfirmationV07 struct {

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionIdentificationDetails *model.TransactionTypeAndAdditionalParameters16 `xml:"TxIdDtls"`

	// Additional parameters to the transaction.
	AdditionalParameters *model.AdditionalParameters24 `xml:"AddtlParams,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *model.SecuritiesTradeDetails55 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount40 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *model.SecuritiesFinancingTransactionDetails28 `xml:"SctiesFincgDtls,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails96 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties26 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *model.AmountAndDirection46 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts31 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingConfirmationV07) AddTransactionIdentificationDetails() *model.TransactionTypeAndAdditionalParameters16 {
	s.TransactionIdentificationDetails = new(model.TransactionTypeAndAdditionalParameters16)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesFinancingConfirmationV07) AddAdditionalParameters() *model.AdditionalParameters24 {
	s.AdditionalParameters = new(model.AdditionalParameters24)
	return s.AdditionalParameters
}

func (s *SecuritiesFinancingConfirmationV07) AddTradeDetails() *model.SecuritiesTradeDetails55 {
	s.TradeDetails = new(model.SecuritiesTradeDetails55)
	return s.TradeDetails
}

func (s *SecuritiesFinancingConfirmationV07) AddFinancialInstrumentIdentification() *model.SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingConfirmationV07) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingConfirmationV07) AddQuantityAndAccountDetails() *model.QuantityAndAccount40 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount40)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingConfirmationV07) AddSecuritiesFinancingDetails() *model.SecuritiesFinancingTransactionDetails28 {
	s.SecuritiesFinancingDetails = new(model.SecuritiesFinancingTransactionDetails28)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingConfirmationV07) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingConfirmationV07) AddSettlementParameters() *model.SettlementDetails96 {
	s.SettlementParameters = new(model.SettlementDetails96)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingConfirmationV07) AddDeliveringSettlementParties() *model.SettlementParties36 {
	s.DeliveringSettlementParties = new(model.SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingConfirmationV07) AddReceivingSettlementParties() *model.SettlementParties36 {
	s.ReceivingSettlementParties = new(model.SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingConfirmationV07) AddCashParties() *model.CashParties26 {
	s.CashParties = new(model.CashParties26)
	return s.CashParties
}

func (s *SecuritiesFinancingConfirmationV07) AddSettledAmount() *model.AmountAndDirection46 {
	s.SettledAmount = new(model.AmountAndDirection46)
	return s.SettledAmount
}

func (s *SecuritiesFinancingConfirmationV07) AddOtherAmounts() *model.OtherAmounts31 {
	s.OtherAmounts = new(model.OtherAmounts31)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingConfirmationV07) AddOtherBusinessParties() *model.OtherParties27 {
	s.OtherBusinessParties = new(model.OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingConfirmationV07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
