package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400102 struct {
	XMLName xml.Name                                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.014.001.02 Document"`
	Message *PEPOrISAOrPortfolioTransferCancellationRequestV02 `xml:"PEPOrISAOrPrtflTrfCxlReqV02"`
}

func (d *Document01400102) AddMessage() *PEPOrISAOrPortfolioTransferCancellationRequestV02 {
	d.Message = new(PEPOrISAOrPortfolioTransferCancellationRequestV02)
	return d.Message
}

// Scope
// An instructing party, eg, a (new) plan manager, sends the PEPOrISAOrPortfolioTransferCancellationRequest message to the executing party, eg, a (old) plan manager, to request the cancellation of a previously sent PEPOrISAOrPortfolioTransferInstruction.
// Usage
// The PEPOrISAOrPortfolioTransferCancellationRequest message is used to request the cancellation of an entire PEPOrISAOrPortfolioTransferInstruction message, ie, all the product transfers that it contained. The cancellation request can be specified either by:
// - quoting the transfer references of all the product transfers listed in the PEPOrISAOrPortfolioTransferInstruction message, or,
// - quoting the details of all the product transfers (this includes TransferReference) listed in PEPOrISAOrPortfolioTransferInstruction message.
// The message identification of the PEPOrISAOrPortfolioTransferInstruction may also be quoted in PreviousReference. It is also possible to request the cancellation of PEPOrISAOrPortfolioTransferInstruction by just quoting its message identification in PreviousReference.
type PEPOrISAOrPortfolioTransferCancellationRequestV02 struct {

	// Identifies the message.
	MessageReference *model.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information related to the transfer instruction to be cancelled.
	CancellationByTransferInstructionDetails *model.PEPISATransfer7 `xml:"CxlByTrfInstrDtls,omitempty"`

	// Reference of the transfer instruction to be cancelled.
	CancellationByReference *model.TransferReference3 `xml:"CxlByRef,omitempty"`
}

func (p *PEPOrISAOrPortfolioTransferCancellationRequestV02) AddMessageReference() *model.MessageIdentification1 {
	p.MessageReference = new(model.MessageIdentification1)
	return p.MessageReference
}

func (p *PEPOrISAOrPortfolioTransferCancellationRequestV02) AddPoolReference() *model.AdditionalReference3 {
	p.PoolReference = new(model.AdditionalReference3)
	return p.PoolReference
}

func (p *PEPOrISAOrPortfolioTransferCancellationRequestV02) AddPreviousReference() *model.AdditionalReference3 {
	p.PreviousReference = new(model.AdditionalReference3)
	return p.PreviousReference
}

func (p *PEPOrISAOrPortfolioTransferCancellationRequestV02) AddRelatedReference() *model.AdditionalReference3 {
	p.RelatedReference = new(model.AdditionalReference3)
	return p.RelatedReference
}

func (p *PEPOrISAOrPortfolioTransferCancellationRequestV02) AddCancellationByTransferInstructionDetails() *model.PEPISATransfer7 {
	p.CancellationByTransferInstructionDetails = new(model.PEPISATransfer7)
	return p.CancellationByTransferInstructionDetails
}

func (p *PEPOrISAOrPortfolioTransferCancellationRequestV02) AddCancellationByReference() *model.TransferReference3 {
	p.CancellationByReference = new(model.TransferReference3)
	return p.CancellationByReference
}
