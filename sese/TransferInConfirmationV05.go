package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700105 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.007.001.05 Document"`
	Message *TransferInConfirmationV05 `xml:"TrfInConf"`
}

func (d *Document00700105) AddMessage() *TransferInConfirmationV05 {
	d.Message = new(TransferInConfirmationV05)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the TransferInConfirmation message to the instructing party, for example, an investment manager or its authorised representative, to confirm the receipt of a financial instrument, free of payment, on a given date, from a specified party.
// This message may also be used to confirm the receipt of a financial instrument, free of payment, from another of the instructing parties own accounts or from a third party.
// Usage
// The TransferInConfirmation message is used to confirm receipt of a financial instrument, either from another account owned by the instructing party or from a third party. The reference of the transfer confirmation is identified in TransferConfirmationReference.
// The reference of the original transfer instruction is specified in TransferReference. The message identification of the TransferInInstruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
type TransferInConfirmationV05 struct {

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
	TransferDetails []*model.Transfer29 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *model.InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *model.DeliverInformation12 `xml:"SttlmDtls,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInConfirmationV05) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInConfirmationV05) AddPoolReference() *model.AdditionalReference2 {
	t.PoolReference = new(model.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferInConfirmationV05) AddPreviousReference() *model.AdditionalReference2 {
	t.PreviousReference = new(model.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferInConfirmationV05) AddRelatedReference() *model.AdditionalReference2 {
	t.RelatedReference = new(model.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInConfirmationV05) SetMasterReference(value string) {
	t.MasterReference = (*model.Max35Text)(&value)
}

func (t *TransferInConfirmationV05) AddTransferDetails() *model.Transfer29 {
	newValue := new(model.Transfer29)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferInConfirmationV05) AddAccountDetails() *model.InvestmentAccount40 {
	t.AccountDetails = new(model.InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferInConfirmationV05) AddSettlementDetails() *model.DeliverInformation12 {
	t.SettlementDetails = new(model.DeliverInformation12)
	return t.SettlementDetails
}

func (t *TransferInConfirmationV05) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	t.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return t.MarketPracticeVersion
}

func (t *TransferInConfirmationV05) AddCopyDetails() *model.CopyInformation2 {
	t.CopyDetails = new(model.CopyInformation2)
	return t.CopyDetails
}

func (t *TransferInConfirmationV05) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
