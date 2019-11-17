package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700102 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.007.001.02 Document"`
	Message *TransferInConfirmationV02 `xml:"TrfInConfV02"`
}

func (d *Document00700102) AddMessage() *TransferInConfirmationV02 {
	d.Message = new(TransferInConfirmationV02)
	return d.Message
}

// Scope
// An executing party, eg, a transfer agent, sends the TransferInConfirmation message to the instructing party, eg, an investment manager or its authorised representative, to confirm the receipt of a financial instrument, free of payment, on a given date, from a specified party.
// This message may also be used to confirm the receipt of a financial instrument, free of payment, from another of the instructing parties own accounts or from a third party.
// Usage
// The TransferInConfirmation message is used to confirm receipt of a financial instrument, either from another account owned by the instructing party or from a third party. The reference of the transfer confirmation is identified in TransferConfirmationReference.
// The reference of the original transfer instruction is specified in TransferReference. The message identification of the TransferInInstruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
type TransferInConfirmationV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails *model.Transfer7 `xml:"TrfDtls"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *model.FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *model.InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *model.DeliverInformation4 `xml:"SttlmDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInConfirmationV02) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInConfirmationV02) AddRelatedReference() *model.AdditionalReference2 {
	t.RelatedReference = new(model.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInConfirmationV02) AddPoolReference() *model.AdditionalReference2 {
	t.PoolReference = new(model.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferInConfirmationV02) AddPreviousReference() *model.AdditionalReference2 {
	t.PreviousReference = new(model.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferInConfirmationV02) AddTransferDetails() *model.Transfer7 {
	t.TransferDetails = new(model.Transfer7)
	return t.TransferDetails
}

func (t *TransferInConfirmationV02) AddFinancialInstrumentDetails() *model.FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(model.FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferInConfirmationV02) AddAccountDetails() *model.InvestmentAccount22 {
	t.AccountDetails = new(model.InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferInConfirmationV02) AddSettlementDetails() *model.DeliverInformation4 {
	t.SettlementDetails = new(model.DeliverInformation4)
	return t.SettlementDetails
}

func (t *TransferInConfirmationV02) AddCopyDetails() *model.CopyInformation2 {
	t.CopyDetails = new(model.CopyInformation2)
	return t.CopyDetails
}

func (t *TransferInConfirmationV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
