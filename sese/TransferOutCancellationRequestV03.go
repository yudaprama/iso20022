package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200103 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.002.001.03 Document"`
	Message *TransferOutCancellationRequestV03 `xml:"TrfOutCxlReq"`
}

func (d *Document00200103) AddMessage() *TransferOutCancellationRequestV03 {
	d.Message = new(TransferOutCancellationRequestV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the TransferOutCancellationRequest message to the executing party, for example, a transfer agent, to request the cancellation of a previously sent TransferOutInstruction.
// Usage
// The TransferOutCancellationRequest message is used to request cancellation of a previously sent TransferOutInstruction. There are two ways to specify the transfer cancellation request. Either:
// - the transfer reference of the original transfer is quoted, or,
// - all the details of the original transfer (this includes TransferReference) are quoted but this is not recommended.
// The message identification of the TransferOutInstruction message in which the original transfer was conveyed may also be quoted in PreviousReference. It is also possible to request the cancellation of a TransferOutInstruction message by quoting its message identification in PreviousReference.
type TransferOutCancellationRequestV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
	References []*model.References11 `xml:"Refs"`

	// Reference of the transfer to be cancelled.
	CancellationByReference *model.TransferReference1 `xml:"CxlByRef,omitempty"`

	// The transfer out request message to cancel.
	CancellationByTransferOutDetails *model.TransferOut7 `xml:"CxlByTrfOutDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (t *TransferOutCancellationRequestV03) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferOutCancellationRequestV03) AddReferences() *model.References11 {
	newValue := new(model.References11)
	t.References = append(t.References, newValue)
	return newValue
}

func (t *TransferOutCancellationRequestV03) AddCancellationByReference() *model.TransferReference1 {
	t.CancellationByReference = new(model.TransferReference1)
	return t.CancellationByReference
}

func (t *TransferOutCancellationRequestV03) AddCancellationByTransferOutDetails() *model.TransferOut7 {
	t.CancellationByTransferOutDetails = new(model.TransferOut7)
	return t.CancellationByTransferOutDetails
}

func (t *TransferOutCancellationRequestV03) AddCopyDetails() *model.CopyInformation2 {
	t.CopyDetails = new(model.CopyInformation2)
	return t.CopyDetails
}
