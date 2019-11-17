package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400102 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.004.001.02 Document"`
	Message *CustodyStatementOfHoldingsCancellationV02 `xml:"CtdyStmtOfHldgsCxlV02"`
}

func (d *Document00400102) AddMessage() *CustodyStatementOfHoldingsCancellationV02 {
	d.Message = new(CustodyStatementOfHoldingsCancellationV02)
	return d.Message
}

// Scope
// An account servicer, for example, a transfer agent, sends the CustodyStatementOfHoldingsCancellation message to the account owner, for example, an investment manager or its authorised representative to cancel a previously sent CustodyStatementOfHoldings message.
// Usage
// The CustodyStatementOfHoldingsCancellation message is used to cancel a previously sent CustodyStatementOfHoldings message. This message must contain the reference of the message to be cancelled.
// This message may also contain details of the message to be cancelled, but this is not recommended.
type CustodyStatementOfHoldingsCancellationV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Number used to sequence pages when it is not possible for data to be conveyed in a single message and the data has to be split across several pages (messages).
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// The Custody Statement of Holdings message to cancel.
	StatementToBeCancelled *model.CustodyStatementOfHoldings2 `xml:"StmtToBeCanc,omitempty"`
}

func (c *CustodyStatementOfHoldingsCancellationV02) AddMessageIdentification() *model.MessageIdentification1 {
	c.MessageIdentification = new(model.MessageIdentification1)
	return c.MessageIdentification
}

func (c *CustodyStatementOfHoldingsCancellationV02) AddPreviousReference() *model.AdditionalReference2 {
	c.PreviousReference = new(model.AdditionalReference2)
	return c.PreviousReference
}

func (c *CustodyStatementOfHoldingsCancellationV02) AddRelatedReference() *model.AdditionalReference2 {
	c.RelatedReference = new(model.AdditionalReference2)
	return c.RelatedReference
}

func (c *CustodyStatementOfHoldingsCancellationV02) AddMessagePagination() *model.Pagination {
	c.MessagePagination = new(model.Pagination)
	return c.MessagePagination
}

func (c *CustodyStatementOfHoldingsCancellationV02) AddStatementToBeCancelled() *model.CustodyStatementOfHoldings2 {
	c.StatementToBeCancelled = new(model.CustodyStatementOfHoldings2)
	return c.StatementToBeCancelled
}
