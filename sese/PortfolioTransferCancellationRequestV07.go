package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400107 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.014.001.07 Document"`
	Message *PortfolioTransferCancellationRequestV07 `xml:"PrtflTrfCxlReq"`
}

func (d *Document01400107) AddMessage() *PortfolioTransferCancellationRequestV07 {
	d.Message = new(PortfolioTransferCancellationRequestV07)
	return d.Message
}

// Scope
// An instructing party, for example, a (new) plan manager (Transferee), sends the PortfolioTransferCancellationRequest message to the executing party, for example, a (old) plan manager (Transferor), to request the cancellation of a previously sent PortfolioTransferInstruction.
// Usage
// The PortfolioTransferCancellationRequest message is used to request the cancellation of an entire PortfolioTransferInstruction message, that is, all the product transfers that it contained. The cancellation request can be specified either by:
// - quoting the transfer references of all the product transfers listed in the PortfolioTransferInstruction message, or,
// - quoting the details of all the product transfers (this includes TransferReference) listed in PortfolioTransferInstruction message.
// The message identification of the PortfolioTransferInstruction may also be quoted in PreviousReference. It is also possible to request the cancellation of PortfolioTransferInstruction by just quoting its message identification in PreviousReference.
type PortfolioTransferCancellationRequestV07 struct {

	// Identifies the message.
	MessageReference *model.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference6 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference6 `xml:"RltdRef,omitempty"`

	// Choice between cancellation by transfer details or reference.
	Cancellation *model.Cancellation11Choice `xml:"Cxl"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`
}

func (p *PortfolioTransferCancellationRequestV07) AddMessageReference() *model.MessageIdentification1 {
	p.MessageReference = new(model.MessageIdentification1)
	return p.MessageReference
}

func (p *PortfolioTransferCancellationRequestV07) AddPoolReference() *model.AdditionalReference6 {
	p.PoolReference = new(model.AdditionalReference6)
	return p.PoolReference
}

func (p *PortfolioTransferCancellationRequestV07) AddPreviousReference() *model.AdditionalReference6 {
	p.PreviousReference = new(model.AdditionalReference6)
	return p.PreviousReference
}

func (p *PortfolioTransferCancellationRequestV07) AddRelatedReference() *model.AdditionalReference6 {
	p.RelatedReference = new(model.AdditionalReference6)
	return p.RelatedReference
}

func (p *PortfolioTransferCancellationRequestV07) AddCancellation() *model.Cancellation11Choice {
	p.Cancellation = new(model.Cancellation11Choice)
	return p.Cancellation
}

func (p *PortfolioTransferCancellationRequestV07) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	p.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return p.MarketPracticeVersion
}
