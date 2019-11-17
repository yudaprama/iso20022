package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700101 struct {
	XMLName xml.Name                                           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.007.001.01 Document"`
	Message *StatementOfInvestmentFundTransactionsCancellation `xml:"semt.007.001.01"`
}

func (d *Document00700101) AddMessage() *StatementOfInvestmentFundTransactionsCancellation {
	d.Message = new(StatementOfInvestmentFundTransactionsCancellation)
	return d.Message
}

// Scope
// The StatementOfInvestmentFundTransactionsCancellation message is sent by an account servicer to the account owner or the account owner's designated agent. The account servicer may be a fund administrator or fund intermediary, trustee or registrar, etc.
// This message is used to cancel a previously sent StatementOfInvestmentFundTransactions message.
// Usage
// The StatementOfInvestmentFundTransactionsCancellation message is sent by an account servicer to the account owner to cancel a previously sent StatementOfInvestmentFundTransactions message.
// This message must contain the reference of the message to be cancelled. This message may also contain all the details of the message to be cancelled, but this is not recommended.
type StatementOfInvestmentFundTransactionsCancellation struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// The Statement of Investment Fund Transactions message to cancel.
	StatementToBeCancelled *model.StatementOfInvestmentFundTransactions1 `xml:"StmtToBeCanc,omitempty"`
}

func (s *StatementOfInvestmentFundTransactionsCancellation) AddPreviousReference() *model.AdditionalReference2 {
	s.PreviousReference = new(model.AdditionalReference2)
	return s.PreviousReference
}

func (s *StatementOfInvestmentFundTransactionsCancellation) AddRelatedReference() *model.AdditionalReference2 {
	s.RelatedReference = new(model.AdditionalReference2)
	return s.RelatedReference
}

func (s *StatementOfInvestmentFundTransactionsCancellation) AddMessagePagination() *model.Pagination {
	s.MessagePagination = new(model.Pagination)
	return s.MessagePagination
}

func (s *StatementOfInvestmentFundTransactionsCancellation) AddStatementToBeCancelled() *model.StatementOfInvestmentFundTransactions1 {
	s.StatementToBeCancelled = new(model.StatementOfInvestmentFundTransactions1)
	return s.StatementToBeCancelled
}
