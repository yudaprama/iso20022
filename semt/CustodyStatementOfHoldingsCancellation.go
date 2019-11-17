package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400101 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.004.001.01 Document"`
	Message *CustodyStatementOfHoldingsCancellation `xml:"semt.004.001.01"`
}

func (d *Document00400101) AddMessage() *CustodyStatementOfHoldingsCancellation {
	d.Message = new(CustodyStatementOfHoldingsCancellation)
	return d.Message
}

// Scope
// The CustodyStatementOfHoldingsCancellation message is sent by an account servicer to the account owner or the account owner's designated agent. The account servicer may be a local agent (sub-custodian) acting on behalf of its global custodian customer, a custodian acting on behalf of an investment management institution or a broker/dealer, a fund administrator or fund intermediary, trustee or registrar, etc.
// This message is used to cancel a previously sent CustodyStatementOfHoldings message.
// Usage
// The CustodyStatementOfHoldingsCancellation message is sent by an account servicer to the account owner to cancel a previously sent CustodyStatementOfHoldings message.
// This message must contain the reference of the message to be cancelled. This message may also contain details of the message to be cancelled, but this is not recommended.
type CustodyStatementOfHoldingsCancellation struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Number used to sequence pages when it is not possible for data to be conveyed in a single message and the data has to be split across several pages (messages).
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// The Custody Statement of Holdings message to cancel.
	StatementToBeCancelled *model.CustodyStatementOfHoldings1 `xml:"StmtToBeCanc,omitempty"`
}

func (c *CustodyStatementOfHoldingsCancellation) AddPreviousReference() *model.AdditionalReference2 {
	c.PreviousReference = new(model.AdditionalReference2)
	return c.PreviousReference
}

func (c *CustodyStatementOfHoldingsCancellation) AddRelatedReference() *model.AdditionalReference2 {
	c.RelatedReference = new(model.AdditionalReference2)
	return c.RelatedReference
}

func (c *CustodyStatementOfHoldingsCancellation) AddMessagePagination() *model.Pagination {
	c.MessagePagination = new(model.Pagination)
	return c.MessagePagination
}

func (c *CustodyStatementOfHoldingsCancellation) AddStatementToBeCancelled() *model.CustodyStatementOfHoldings1 {
	c.StatementToBeCancelled = new(model.CustodyStatementOfHoldings1)
	return c.StatementToBeCancelled
}
