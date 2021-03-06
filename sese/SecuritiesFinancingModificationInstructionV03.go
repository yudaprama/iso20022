package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03600103 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.03 Document"`
	Message *SecuritiesFinancingModificationInstructionV03 `xml:"SctiesFincgModInstr"`
}

func (d *Document03600103) AddMessage() *SecuritiesFinancingModificationInstructionV03 {
	d.Message = new(SecuritiesFinancingModificationInstructionV03)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesFinancingModificationInstruction to a securities financing transaction account servicer to notify the securities financing transaction account servicer of an update in the details of a repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing transaction that does not impact the original transaction securities quantity.
// Such a change may be:
// - the providing of closing details not available at the time of the sending of the Securities Financing Instruction, for example, termination date for an open repo,
// - the providing of a new rate, for example, a repo rate,
// - the rollover of a position extending the closing or maturity date.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with a local custodian, or
// - an investment management institution which manage a fund account opened at a custodian, or
// - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure, or
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct the settlement of securities financing transactions to a central securities depository or another settlement market infrastructure.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesFinancingModificationInstructionV03 struct {

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing), modification information and other parameters.
	TransactionTypeAndModificationAdditionalParameters *model.TransactionTypeAndAdditionalParameters7 `xml:"TxTpAndModAddtlParams"`

	// Details of the securities financing deal.
	TradeDetails *model.SecuritiesTradeDetails5 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification14 `xml:"FinInstrmId"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount16 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingAdditionalDetails *model.SecuritiesFinancingTransactionDetails19 `xml:"SctiesFincgAddtlDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails62 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties10 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties10 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *model.AmountAndDirection10 `xml:"OpngSttlmAmt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingModificationInstructionV03) AddTransactionTypeAndModificationAdditionalParameters() *model.TransactionTypeAndAdditionalParameters7 {
	s.TransactionTypeAndModificationAdditionalParameters = new(model.TransactionTypeAndAdditionalParameters7)
	return s.TransactionTypeAndModificationAdditionalParameters
}

func (s *SecuritiesFinancingModificationInstructionV03) AddTradeDetails() *model.SecuritiesTradeDetails5 {
	s.TradeDetails = new(model.SecuritiesTradeDetails5)
	return s.TradeDetails
}

func (s *SecuritiesFinancingModificationInstructionV03) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingModificationInstructionV03) AddQuantityAndAccountDetails() *model.QuantityAndAccount16 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount16)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingModificationInstructionV03) AddSecuritiesFinancingAdditionalDetails() *model.SecuritiesFinancingTransactionDetails19 {
	s.SecuritiesFinancingAdditionalDetails = new(model.SecuritiesFinancingTransactionDetails19)
	return s.SecuritiesFinancingAdditionalDetails
}

func (s *SecuritiesFinancingModificationInstructionV03) AddSettlementParameters() *model.SettlementDetails62 {
	s.SettlementParameters = new(model.SettlementDetails62)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingModificationInstructionV03) AddDeliveringSettlementParties() *model.SettlementParties10 {
	s.DeliveringSettlementParties = new(model.SettlementParties10)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingModificationInstructionV03) AddReceivingSettlementParties() *model.SettlementParties10 {
	s.ReceivingSettlementParties = new(model.SettlementParties10)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingModificationInstructionV03) AddOpeningSettlementAmount() *model.AmountAndDirection10 {
	s.OpeningSettlementAmount = new(model.AmountAndDirection10)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingModificationInstructionV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
