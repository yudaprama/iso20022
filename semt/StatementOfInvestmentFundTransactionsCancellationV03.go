package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700103 struct {
	XMLName xml.Name                                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.007.001.03 Document"`
	Message *StatementOfInvestmentFundTransactionsCancellationV03 `xml:"StmtOfInvstmtFndTxsCxl"`
}

func (d *Document00700103) AddMessage() *StatementOfInvestmentFundTransactionsCancellationV03 {
	d.Message = new(StatementOfInvestmentFundTransactionsCancellationV03)
	return d.Message
}

// Scope
// An account servicer, for example, a transfer agent sends the StatementOfInvestmentFundTransactionsCancellation message to the account owner, for example, an investment manager or its authorised representative to cancel a previously sent StatementOfInvestmentFundTransactions message.
// Usage
// The StatementOfInvestmentFundTransactionsCancellation message is used to cancel a previously sent StatementOfInvestmentFundTransactions message. This message must contain the reference of the message to be cancelled.
// This message may also contain all the details of the message to be cancelled, but this is not recommended.
type StatementOfInvestmentFundTransactionsCancellationV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// The Statement of Investment Fund Transactions message to cancel.
	StatementToBeCancelled *model.StatementOfInvestmentFundTransactions3 `xml:"StmtToBeCanc,omitempty"`
}

func (s *StatementOfInvestmentFundTransactionsCancellationV03) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *StatementOfInvestmentFundTransactionsCancellationV03) AddPreviousReference() *model.AdditionalReference2 {
	s.PreviousReference = new(model.AdditionalReference2)
	return s.PreviousReference
}

func (s *StatementOfInvestmentFundTransactionsCancellationV03) AddRelatedReference() *model.AdditionalReference2 {
	s.RelatedReference = new(model.AdditionalReference2)
	return s.RelatedReference
}

func (s *StatementOfInvestmentFundTransactionsCancellationV03) AddMessagePagination() *model.Pagination {
	s.MessagePagination = new(model.Pagination)
	return s.MessagePagination
}

func (s *StatementOfInvestmentFundTransactionsCancellationV03) AddStatementToBeCancelled() *model.StatementOfInvestmentFundTransactions3 {
	s.StatementToBeCancelled = new(model.StatementOfInvestmentFundTransactions3)
	return s.StatementToBeCancelled
}
