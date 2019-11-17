package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600104 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.006.001.04 Document"`
	Message *TransferInCancellationRequestV04 `xml:"TrfInCxlReq"`
}

func (d *Document00600104) AddMessage() *TransferInCancellationRequestV04 {
	d.Message = new(TransferInCancellationRequestV04)
	return d.Message
}

// Scope
// An instructing party, for example, a transfer agent, sends the TransferInCancellationRequest message to the executing party, for example, a transfer agent, to request the cancellation of a previously sent TransferInInstruction.
// Usage
// The TransferInCancellationRequest message is used to request cancellation of a previously sent TransferInInstruction.
// There are two ways to specify the transfer cancellation request. Either:
// - the transfer reference of the original transfer is quoted, or,
// - all the details of the original transfer (this includes TransferReference) are quoted but this is not recommended.
// The message identification of the TransferInInstruction message in which the transfer was conveyed may also be quoted in PreviousReference. It is also possible to request the cancellation of a TransferInInstruction message by quoting its message identification in PreviousReference.
type TransferInCancellationRequestV04 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
	References []*model.References11 `xml:"Refs"`

	// Choice between cancellation by reference or by transfer details.
	Cancellation *model.Cancellation2Choice `xml:"Cxl"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (t *TransferInCancellationRequestV04) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInCancellationRequestV04) AddReferences() *model.References11 {
	newValue := new(model.References11)
	t.References = append(t.References, newValue)
	return newValue
}

func (t *TransferInCancellationRequestV04) AddCancellation() *model.Cancellation2Choice {
	t.Cancellation = new(model.Cancellation2Choice)
	return t.Cancellation
}

func (t *TransferInCancellationRequestV04) AddCopyDetails() *model.CopyInformation2 {
	t.CopyDetails = new(model.CopyInformation2)
	return t.CopyDetails
}
