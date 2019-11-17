package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600102 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.006.001.02 Document"`
	Message *StatementOfInvestmentFundTransactionsV02 `xml:"StmtOfInvstmtFndTxsV02"`
}

func (d *Document00600102) AddMessage() *StatementOfInvestmentFundTransactionsV02 {
	d.Message = new(StatementOfInvestmentFundTransactionsV02)
	return d.Message
}

// Scope
// An account servicer, for example, a transfer agent sends the StatementOfInvestmentFundTransactions message to the account owner, for example, an investment manager or its authorised representative to provide detailed transactions (increases and decreases) of holdings which occurred during a specified period of time.
// Usage
// The StatementOfInvestmentFundTransactions message is used to list the holdings transactions of a single (master) account or several sub-accounts.
// This message should be used at a frequency agreed bi-laterally between the account servicer and the account owner.
// This message must not be used in place of confirmation messages.
type StatementOfInvestmentFundTransactionsV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// General information related to the investment fund statement of transactions.
	StatementGeneralDetails *model.Statement8 `xml:"StmtGnlDtls"`

	// Information related to an investment account.
	InvestmentAccountDetails *model.InvestmentAccount25 `xml:"InvstmtAcctDtls"`

	// Creation/cancellation of investment units on the books of the fund or its designated agent, as a result of executing an investment fund order.
	TransactionOnAccount []*model.InvestmentFundTransactionsByFund2 `xml:"TxOnAcct,omitempty"`

	// The sub-account of the safekeeping or investment account.
	SubAccountDetails []*model.SubAccountIdentification6 `xml:"SubAcctDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *StatementOfInvestmentFundTransactionsV02) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *StatementOfInvestmentFundTransactionsV02) AddPreviousReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *StatementOfInvestmentFundTransactionsV02) AddRelatedReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	s.RelatedReference = append(s.RelatedReference, newValue)
	return newValue
}

func (s *StatementOfInvestmentFundTransactionsV02) AddMessagePagination() *model.Pagination {
	s.MessagePagination = new(model.Pagination)
	return s.MessagePagination
}

func (s *StatementOfInvestmentFundTransactionsV02) AddStatementGeneralDetails() *model.Statement8 {
	s.StatementGeneralDetails = new(model.Statement8)
	return s.StatementGeneralDetails
}

func (s *StatementOfInvestmentFundTransactionsV02) AddInvestmentAccountDetails() *model.InvestmentAccount25 {
	s.InvestmentAccountDetails = new(model.InvestmentAccount25)
	return s.InvestmentAccountDetails
}

func (s *StatementOfInvestmentFundTransactionsV02) AddTransactionOnAccount() *model.InvestmentFundTransactionsByFund2 {
	newValue := new(model.InvestmentFundTransactionsByFund2)
	s.TransactionOnAccount = append(s.TransactionOnAccount, newValue)
	return newValue
}

func (s *StatementOfInvestmentFundTransactionsV02) AddSubAccountDetails() *model.SubAccountIdentification6 {
	newValue := new(model.SubAccountIdentification6)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}

func (s *StatementOfInvestmentFundTransactionsV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
