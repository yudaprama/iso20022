package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01900105 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.019.001.05 Document"`
	Message *InitialBaselineSubmissionV05 `xml:"InitlBaselnSubmissn"`
}

func (d *Document01900105) AddMessage() *InitialBaselineSubmissionV05 {
	d.Message = new(InitialBaselineSubmissionV05)
	return d.Message
}

// Scope
// The InitialBaselineSubmission message is sent by the initiator of a transaction to the matching application.
// This message is used to initiate a transaction.
// Usage
// The InitialBaselineSubmission message can be sent by a party to register a transaction in the matching application. The message can be submitted with either lodge or push-through instruction.
// When the push-through instruction is present, the matching application acknowledges the receipt of the message to the sender by sending an Acknowledgement message, stores the submitted information and informs the counterparty about the registration of the transaction by sending a FullPushThroughReport message. With the BaselineReSubmission message the counterparty responds with matching baseline information in order to establish the transaction (baseline).
// When the lodge instruction is present, the matching application acknowledges the receipt of the message to the sender by sending an Acknowledgement message and stores the submitted information. No matching of the submitted baseline data with other baseline information will take place. For example the submission of an InitialBaselineSubmission message containing a lodge instruction establishes the transaction (baseline) in the matching application.
// The InitialBaselineSubmission message consists of data which relates to the purchasing agreement covered by the transaction, for example line item details, shipping details.
type InitialBaselineSubmissionV05 struct {

	// Identifies the submitted information
	SubmissionIdentification *model.MessageIdentification1 `xml:"SubmissnId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *model.SimpleIdentificationInformation `xml:"SubmitrTxRef"`

	// Specifies the instruction requested by the submitter by means of a code.
	Instruction *model.InstructionType1 `xml:"Instr"`

	// Specifies the commercial details of the underlying transaction.
	Baseline *model.Baseline5 `xml:"Baseln"`

	// Person to be contacted in the organisation of the buyer.
	BuyerContactPerson []*model.ContactIdentification1 `xml:"BuyrCtctPrsn,omitempty"`

	// Person to be contacted in the organisation of the seller.
	SellerContactPerson []*model.ContactIdentification1 `xml:"SellrCtctPrsn,omitempty"`

	// Person to be contacted in the seller's bank or buyer's bank.
	BankContactPerson *model.BankContactPerson1Choice `xml:"BkCtctPrsn"`

	// Person to be contacted in another bank than seller or buyer's bank.
	OtherBankContactPerson []*model.ContactIdentification3 `xml:"OthrBkCtctPrsn,omitempty"`
}

func (i *InitialBaselineSubmissionV05) AddSubmissionIdentification() *model.MessageIdentification1 {
	i.SubmissionIdentification = new(model.MessageIdentification1)
	return i.SubmissionIdentification
}

func (i *InitialBaselineSubmissionV05) AddSubmitterTransactionReference() *model.SimpleIdentificationInformation {
	i.SubmitterTransactionReference = new(model.SimpleIdentificationInformation)
	return i.SubmitterTransactionReference
}

func (i *InitialBaselineSubmissionV05) AddInstruction() *model.InstructionType1 {
	i.Instruction = new(model.InstructionType1)
	return i.Instruction
}

func (i *InitialBaselineSubmissionV05) AddBaseline() *model.Baseline5 {
	i.Baseline = new(model.Baseline5)
	return i.Baseline
}

func (i *InitialBaselineSubmissionV05) AddBuyerContactPerson() *model.ContactIdentification1 {
	newValue := new(model.ContactIdentification1)
	i.BuyerContactPerson = append(i.BuyerContactPerson, newValue)
	return newValue
}

func (i *InitialBaselineSubmissionV05) AddSellerContactPerson() *model.ContactIdentification1 {
	newValue := new(model.ContactIdentification1)
	i.SellerContactPerson = append(i.SellerContactPerson, newValue)
	return newValue
}

func (i *InitialBaselineSubmissionV05) AddBankContactPerson() *model.BankContactPerson1Choice {
	i.BankContactPerson = new(model.BankContactPerson1Choice)
	return i.BankContactPerson
}

func (i *InitialBaselineSubmissionV05) AddOtherBankContactPerson() *model.ContactIdentification3 {
	newValue := new(model.ContactIdentification3)
	i.OtherBankContactPerson = append(i.OtherBankContactPerson, newValue)
	return newValue
}
