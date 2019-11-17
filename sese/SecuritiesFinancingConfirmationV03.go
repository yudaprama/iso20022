package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03500103 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.035.001.03 Document"`
	Message *SecuritiesFinancingConfirmationV03 `xml:"SctiesFincgConf"`
}

func (d *Document03500103) AddMessage() *SecuritiesFinancingConfirmationV03 {
	d.Message = new(SecuritiesFinancingConfirmationV03)
	return d.Message
}

// Scope
// A securities financing transaction account servicer sends a SecuritiesFinancingConfirmation to an account owner to confirm or advise of the partial or full settlement of the opening or closing leg of a securities financing transaction.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure managing securities financing transactions on behalf of their participants
// - an agent (sub-custodian) managing securities financing transactions on behalf of their global custodian customer, or
// - a custodian managing securities financing transactions on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesFinancingConfirmationV03 struct {

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionIdentificationDetails *model.TransactionTypeAndAdditionalParameters3 `xml:"TxIdDtls"`

	// Additional parameters to the transaction.
	AdditionalParameters *model.AdditionalParameters12 `xml:"AddtlParams,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *model.SecuritiesTradeDetails6 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount18 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *model.SecuritiesFinancingTransactionDetails11 `xml:"SctiesFincgDtls,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction3 `xml:"StgSttlmInstrDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails60 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties10 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties10 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties7 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *model.AmountAndDirection2 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts17 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingConfirmationV03) AddTransactionIdentificationDetails() *model.TransactionTypeAndAdditionalParameters3 {
	s.TransactionIdentificationDetails = new(model.TransactionTypeAndAdditionalParameters3)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesFinancingConfirmationV03) AddAdditionalParameters() *model.AdditionalParameters12 {
	s.AdditionalParameters = new(model.AdditionalParameters12)
	return s.AdditionalParameters
}

func (s *SecuritiesFinancingConfirmationV03) AddTradeDetails() *model.SecuritiesTradeDetails6 {
	s.TradeDetails = new(model.SecuritiesTradeDetails6)
	return s.TradeDetails
}

func (s *SecuritiesFinancingConfirmationV03) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingConfirmationV03) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingConfirmationV03) AddQuantityAndAccountDetails() *model.QuantityAndAccount18 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount18)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingConfirmationV03) AddSecuritiesFinancingDetails() *model.SecuritiesFinancingTransactionDetails11 {
	s.SecuritiesFinancingDetails = new(model.SecuritiesFinancingTransactionDetails11)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingConfirmationV03) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction3 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction3)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingConfirmationV03) AddSettlementParameters() *model.SettlementDetails60 {
	s.SettlementParameters = new(model.SettlementDetails60)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingConfirmationV03) AddDeliveringSettlementParties() *model.SettlementParties10 {
	s.DeliveringSettlementParties = new(model.SettlementParties10)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingConfirmationV03) AddReceivingSettlementParties() *model.SettlementParties10 {
	s.ReceivingSettlementParties = new(model.SettlementParties10)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingConfirmationV03) AddCashParties() *model.CashParties7 {
	s.CashParties = new(model.CashParties7)
	return s.CashParties
}

func (s *SecuritiesFinancingConfirmationV03) AddSettledAmount() *model.AmountAndDirection2 {
	s.SettledAmount = new(model.AmountAndDirection2)
	return s.SettledAmount
}

func (s *SecuritiesFinancingConfirmationV03) AddOtherAmounts() *model.OtherAmounts17 {
	s.OtherAmounts = new(model.OtherAmounts17)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingConfirmationV03) AddOtherBusinessParties() *model.OtherParties19 {
	s.OtherBusinessParties = new(model.OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingConfirmationV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
