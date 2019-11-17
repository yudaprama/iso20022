package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400104 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.014.001.04 Document"`
	Message *PortfolioTransferCancellationRequestV04 `xml:"PrtflTrfCxlReq"`
}

func (d *Document01400104) AddMessage() *PortfolioTransferCancellationRequestV04 {
	d.Message = new(PortfolioTransferCancellationRequestV04)
	return d.Message
}

// Scope
// An instructing party, for example, a (new) plan manager (Transferee), sends the PortfolioTransferCancellationRequest message to the executing party, for example, a (old) plan manager (Transferor), to request the cancellation of a previously sent PortfolioTransferInstruction.
// Usage
// The PortfolioTransferCancellationRequest message is used to request the cancellation of an entire PortfolioTransferInstruction message, ie, all the product transfers that it contained. The cancellation request can be specified either by:
// - quoting the transfer references of all the product transfers listed in the PortfolioTransferInstruction message, or,
// - quoting the details of all the product transfers (this includes TransferReference) listed in PortfolioTransferInstruction message.
// The message identification of the PortfolioTransferInstruction may also be quoted in PreviousReference. It is also possible to request the cancellation of PortfolioTransferInstruction by just quoting its message identification in PreviousReference.
type PortfolioTransferCancellationRequestV04 struct {

	// Identifies the message.
	MessageReference *model.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Choice between cancellation by transfer details or reference.
	Cancellation *model.Cancellation3Choice `xml:"Cxl"`
}

func (p *PortfolioTransferCancellationRequestV04) AddMessageReference() *model.MessageIdentification1 {
	p.MessageReference = new(model.MessageIdentification1)
	return p.MessageReference
}

func (p *PortfolioTransferCancellationRequestV04) AddPoolReference() *model.AdditionalReference3 {
	p.PoolReference = new(model.AdditionalReference3)
	return p.PoolReference
}

func (p *PortfolioTransferCancellationRequestV04) AddPreviousReference() *model.AdditionalReference3 {
	p.PreviousReference = new(model.AdditionalReference3)
	return p.PreviousReference
}

func (p *PortfolioTransferCancellationRequestV04) AddRelatedReference() *model.AdditionalReference3 {
	p.RelatedReference = new(model.AdditionalReference3)
	return p.RelatedReference
}

func (p *PortfolioTransferCancellationRequestV04) AddCancellation() *model.Cancellation3Choice {
	p.Cancellation = new(model.Cancellation3Choice)
	return p.Cancellation
}
