package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200102 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.002.001.02 Document"`
	Message *TransferOutCancellationRequestV02 `xml:"TrfOutCxlReqV02"`
}

func (d *Document00200102) AddMessage() *TransferOutCancellationRequestV02 {
	d.Message = new(TransferOutCancellationRequestV02)
	return d.Message
}

// Scope
// An instructing party, eg, an investment manager or its authorised representative, sends the TransferOutCancellationRequest message to the executing party, eg, a transfer agent, to request the cancellation of a previously sent TransferOutInstruction.
// Usage
// The TransferOutCancellationRequest message is used to request cancellation of a previously sent TransferOutInstruction. There are two ways to specify the transfer cancellation request. Either:
// - the transfer reference of the original transfer is quoted, or,
// - all the details of the original transfer (this includes TransferReference) are quoted but this is not recommended.
// The message identification of the TransferOutInstruction message in which the original transfer was conveyed may also be quoted in PreviousReference. It is also possible to request the cancellation of a TransferOutInstruction message by quoting its message identification in PreviousReference.
type TransferOutCancellationRequestV02 struct {

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

	// The transfer out request message to cancel.
	CancellationByTransferOutDetails *model.TransferOut5 `xml:"CxlByTrfOutDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (t *TransferOutCancellationRequestV02) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferOutCancellationRequestV02) AddPreviousReference() *model.AdditionalReference2 {
	t.PreviousReference = new(model.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferOutCancellationRequestV02) AddPoolReference() *model.AdditionalReference2 {
	t.PoolReference = new(model.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferOutCancellationRequestV02) AddRelatedReference() *model.AdditionalReference2 {
	t.RelatedReference = new(model.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferOutCancellationRequestV02) AddCancellationByReference() *model.TransferReference1 {
	t.CancellationByReference = new(model.TransferReference1)
	return t.CancellationByReference
}

func (t *TransferOutCancellationRequestV02) AddCancellationByTransferOutDetails() *model.TransferOut5 {
	t.CancellationByTransferOutDetails = new(model.TransferOut5)
	return t.CancellationByTransferOutDetails
}

func (t *TransferOutCancellationRequestV02) AddCopyDetails() *model.CopyInformation2 {
	t.CopyDetails = new(model.CopyInformation2)
	return t.CopyDetails
}
