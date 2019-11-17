package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.003.001.02 Document"`
	Message *TransferOutConfirmationV02 `xml:"TrfOutConfV02"`
}

func (d *Document00300102) AddMessage() *TransferOutConfirmationV02 {
	d.Message = new(TransferOutConfirmationV02)
	return d.Message
}

// Scope
// An executing party, eg, a transfer agent, sends the TransferOutConfirmation message to the instructing party, eg, an investment manager or its authorised representative, to confirm the delivery of a financial instrument, free of payment, on a given date, to a specified party.
// This message may also be used to confirm the delivery of a financial instrument, free of payment, to another of the instructing parties own accounts or to a third party.
// Usage
// The TransferOutConfirmation message is used to confirm the withdrawal of a financial instrument from the owner's account and its delivery to another own account, or to a third party, has taken place.
// The reference of the transfer confirmation is identified in TransferConfirmationReference. The reference of the original transfer instruction is specified in TransferReference. The message identification of the TransferOutInstruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
type TransferOutConfirmationV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails *model.Transfer10 `xml:"TrfDtls"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *model.FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *model.InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *model.ReceiveInformation4 `xml:"SttlmDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOutConfirmationV02) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferOutConfirmationV02) AddRelatedReference() *model.AdditionalReference2 {
	t.RelatedReference = new(model.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferOutConfirmationV02) AddPoolReference() *model.AdditionalReference2 {
	t.PoolReference = new(model.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferOutConfirmationV02) AddPreviousReference() *model.AdditionalReference2 {
	t.PreviousReference = new(model.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferOutConfirmationV02) AddTransferDetails() *model.Transfer10 {
	t.TransferDetails = new(model.Transfer10)
	return t.TransferDetails
}

func (t *TransferOutConfirmationV02) AddFinancialInstrumentDetails() *model.FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(model.FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferOutConfirmationV02) AddAccountDetails() *model.InvestmentAccount22 {
	t.AccountDetails = new(model.InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOutConfirmationV02) AddSettlementDetails() *model.ReceiveInformation4 {
	t.SettlementDetails = new(model.ReceiveInformation4)
	return t.SettlementDetails
}

func (t *TransferOutConfirmationV02) AddCopyDetails() *model.CopyInformation2 {
	t.CopyDetails = new(model.CopyInformation2)
	return t.CopyDetails
}

func (t *TransferOutConfirmationV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
