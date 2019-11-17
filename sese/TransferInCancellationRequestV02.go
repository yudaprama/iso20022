package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.006.001.02 Document"`
	Message *TransferInCancellationRequestV02 `xml:"TrfInCxlReqV02"`
}

func (d *Document00600102) AddMessage() *TransferInCancellationRequestV02 {
	d.Message = new(TransferInCancellationRequestV02)
	return d.Message
}

// Scope
// An instructing party, eg, a transfer agent, sends the TransferInCancellationRequest message to the executing party, eg, a transfer agent, to request the cancellation of a previously sent TransferInInstruction.
// Usage
// The TransferInCancellationRequest message is used to request cancellation of a previously sent TransferInInstruction.
// There are two ways to specify the transfer cancellation request. Either:
// - the transfer reference of the original transfer is quoted, or,
// - all the details of the original transfer (this includes TransferReference) are quoted but this is not recommended.
// The message identification of the TransferInInstruction message in which the transfer was conveyed may also be quoted in PreviousReference. It is also possible to request the cancellation of a TransferInInstruction message by quoting its message identification in PreviousReference.
type TransferInCancellationRequestV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Reference of the transfer to be cancelled.
	CancellationByReference *model.TransferReference1 `xml:"CxlByRef,omitempty"`

	// The transfer in request message to cancel.
	CancellationByTransferInDetails *model.TransferIn3 `xml:"CxlByTrfInDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (t *TransferInCancellationRequestV02) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInCancellationRequestV02) AddPreviousReference() *model.AdditionalReference2 {
	t.PreviousReference = new(model.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferInCancellationRequestV02) AddPoolReference() *model.AdditionalReference2 {
	t.PoolReference = new(model.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferInCancellationRequestV02) AddRelatedReference() *model.AdditionalReference2 {
	t.RelatedReference = new(model.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInCancellationRequestV02) AddCancellationByReference() *model.TransferReference1 {
	t.CancellationByReference = new(model.TransferReference1)
	return t.CancellationByReference
}

func (t *TransferInCancellationRequestV02) AddCancellationByTransferInDetails() *model.TransferIn3 {
	t.CancellationByTransferInDetails = new(model.TransferIn3)
	return t.CancellationByTransferInDetails
}

func (t *TransferInCancellationRequestV02) AddCopyDetails() *model.CopyInformation2 {
	t.CopyDetails = new(model.CopyInformation2)
	return t.CopyDetails
}
