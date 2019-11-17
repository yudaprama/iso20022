package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02200102 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Document"`
	Message *IdentificationModificationAdviceV02 `xml:"IdModAdvc"`
}

func (d *Document02200102) AddMessage() *IdentificationModificationAdviceV02 {
	d.Message = new(IdentificationModificationAdviceV02)
	return d.Message
}

// Scope
// The IdentificationModificationAdvice message is sent by an assigner to an assignee. The message is used to advice on the correct party and/or account identification information.
// Usage
// The IdentificationModificationAdvice message is sent after the receipt of one or several transaction messages that included no longer valid party and/or account identification information.
// The IdentificationModificationAdvice message is exchanged between financial institutions and between financial institutions and non financial institutions and can contain one or more modification advises.
// There is no time limit on the time between the sending of an IdentificationModificationAdvice message and the receipt of the transaction messages that the IdentificationModificationAdvice refers to.
// The IdentificationModificationAdvice includes the correct party and/or account identification information, the IdentificationModificationAdvice or the included information may be forwarded to the initiating party of the transaction messages.
type IdentificationModificationAdviceV02 struct {

	// Identifies the identification assignment.
	Assignment *model.IdentificationAssignment2 `xml:"Assgnmt"`

	// Provides reference information on the original message.
	OriginalTransactionReference *model.OriginalTransactionReference18 `xml:"OrgnlTxRef,omitempty"`

	// Information concerning the identification data that is advised to be modified.
	Modification []*model.IdentificationModification2 `xml:"Mod"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IdentificationModificationAdviceV02) AddAssignment() *model.IdentificationAssignment2 {
	i.Assignment = new(model.IdentificationAssignment2)
	return i.Assignment
}

func (i *IdentificationModificationAdviceV02) AddOriginalTransactionReference() *model.OriginalTransactionReference18 {
	i.OriginalTransactionReference = new(model.OriginalTransactionReference18)
	return i.OriginalTransactionReference
}

func (i *IdentificationModificationAdviceV02) AddModification() *model.IdentificationModification2 {
	newValue := new(model.IdentificationModification2)
	i.Modification = append(i.Modification, newValue)
	return newValue
}

func (i *IdentificationModificationAdviceV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
