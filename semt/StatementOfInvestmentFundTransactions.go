package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.006.001.01 Document"`
	Message *StatementOfInvestmentFundTransactions `xml:"semt.006.001.01"`
}

func (d *Document00600101) AddMessage() *StatementOfInvestmentFundTransactions {
	d.Message = new(StatementOfInvestmentFundTransactions)
	return d.Message
}

// Scope
// The StatementOfInvestmentFundTransactions is sent by an account servicer to the account owner or the account owner's designated agent. The account servicer may be a fund administrator or fund intermediary, trustee or registrar.
// This message provides the details of increases and decreases of holdings which occurred during a specified period.
// This message can also be used for information purposes, eg, tax information.
// Usage
// The StatementOfInvestmentFundTransactions message can be sent:
// - At a frequency agreed bi-laterally between the Sender and the Receiver and/or
// - As a response to a request for statement sent by the account owner.
// The StatementOfInvestmentFundTransactions message can only be used to list the transactions of a single (master) account. However, it is possible to break down these transactions into one or several sub-accounts. Therefore, the message can be used to either specify transactions at
// - the main account level, or
// - the sub-account level.
// This message must not be used in place of confirmation messages.
// Since a SWIFT message as sent is restricted to the maximum input message length, several messages may be needed to accommodate all the information.
type StatementOfInvestmentFundTransactions struct {

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// General information related to the investment fund statement of transactions.
	StatementGeneralDetails *model.Statement5 `xml:"StmtGnlDtls"`

	// Information related to an investment account.
	InvestmentAccountDetails *model.InvestmentAccount12 `xml:"InvstmtAcctDtls"`

	// Creation/cancellation of investment units on the books of the fund or its designated agent, as a result of executing an investment fund order.
	TransactionOnAccount []*model.InvestmentFundTransactionsByFund1 `xml:"TxOnAcct,omitempty"`

	// The sub-account of the safekeeping or investment account.
	SubAccountDetails []*model.SubAccountIdentification4 `xml:"SubAcctDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *StatementOfInvestmentFundTransactions) AddPreviousReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *StatementOfInvestmentFundTransactions) AddRelatedReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	s.RelatedReference = append(s.RelatedReference, newValue)
	return newValue
}

func (s *StatementOfInvestmentFundTransactions) AddMessagePagination() *model.Pagination {
	s.MessagePagination = new(model.Pagination)
	return s.MessagePagination
}

func (s *StatementOfInvestmentFundTransactions) AddStatementGeneralDetails() *model.Statement5 {
	s.StatementGeneralDetails = new(model.Statement5)
	return s.StatementGeneralDetails
}

func (s *StatementOfInvestmentFundTransactions) AddInvestmentAccountDetails() *model.InvestmentAccount12 {
	s.InvestmentAccountDetails = new(model.InvestmentAccount12)
	return s.InvestmentAccountDetails
}

func (s *StatementOfInvestmentFundTransactions) AddTransactionOnAccount() *model.InvestmentFundTransactionsByFund1 {
	newValue := new(model.InvestmentFundTransactionsByFund1)
	s.TransactionOnAccount = append(s.TransactionOnAccount, newValue)
	return newValue
}

func (s *StatementOfInvestmentFundTransactions) AddSubAccountDetails() *model.SubAccountIdentification4 {
	newValue := new(model.SubAccountIdentification4)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}

func (s *StatementOfInvestmentFundTransactions) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
