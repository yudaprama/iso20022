package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300104 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.003.001.04 Document"`
	Message *TransferOutConfirmationV04 `xml:"TrfOutConf"`
}

func (d *Document00300104) AddMessage() *TransferOutConfirmationV04 {
	d.Message = new(TransferOutConfirmationV04)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the TransferOutConfirmation message to the instructing party, for example, an investment manager or its authorised representative, to confirm the delivery of a financial instrument, free of payment, on a given date, to a specified party.
// This message may also be used to confirm the delivery of a financial instrument, free of payment, to another of the instructing parties own accounts or to a third party.
// Usage
// The TransferOutConfirmation message is used to confirm the withdrawal of a financial instrument from the owner's account and its delivery to another own account, or to a third party, has taken place.
// The reference of the transfer confirmation is identified in TransferConfirmationReference. The reference of the original transfer instruction is specified in TransferReference. The message identification of the TransferOutInstruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
type TransferOutConfirmationV04 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *model.Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*model.Transfer23 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *model.InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *model.ReceiveInformation11 `xml:"SttlmDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOutConfirmationV04) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferOutConfirmationV04) AddPoolReference() *model.AdditionalReference2 {
	t.PoolReference = new(model.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferOutConfirmationV04) AddPreviousReference() *model.AdditionalReference2 {
	t.PreviousReference = new(model.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferOutConfirmationV04) AddRelatedReference() *model.AdditionalReference2 {
	t.RelatedReference = new(model.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferOutConfirmationV04) SetMasterReference(value string) {
	t.MasterReference = (*model.Max35Text)(&value)
}

func (t *TransferOutConfirmationV04) AddTransferDetails() *model.Transfer23 {
	newValue := new(model.Transfer23)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOutConfirmationV04) AddAccountDetails() *model.InvestmentAccount22 {
	t.AccountDetails = new(model.InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOutConfirmationV04) AddSettlementDetails() *model.ReceiveInformation11 {
	t.SettlementDetails = new(model.ReceiveInformation11)
	return t.SettlementDetails
}

func (t *TransferOutConfirmationV04) AddCopyDetails() *model.CopyInformation2 {
	t.CopyDetails = new(model.CopyInformation2)
	return t.CopyDetails
}

func (t *TransferOutConfirmationV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
