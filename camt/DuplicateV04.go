package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03400104 struct {
	XMLName xml.Name      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.034.001.04 Document"`
	Message *DuplicateV04 `xml:"Dplct"`
}

func (d *Document03400104) AddMessage() *DuplicateV04 {
	d.Message = new(DuplicateV04)
	return d.Message
}

// Scope
// The Duplicate message is used by financial institutions, with their own offices, and/or with other financial institutions with which they have established bilateral agreements. It allows to exchange duplicate payment instructions.
// Usage
// This message must be sent in response to a Request For Duplicate message.
// The Duplicate Data element must contain a well formed XML document. This means XML special characters such as '<' must be used in a way that is consistent with XML well-formedness criteria.
type DuplicateV04 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment3 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *model.Case3 `xml:"Case"`

	// Duplicate of a previously sent message.
	Duplicate *model.ProprietaryData4 `xml:"Dplct"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DuplicateV04) AddAssignment() *model.CaseAssignment3 {
	d.Assignment = new(model.CaseAssignment3)
	return d.Assignment
}

func (d *DuplicateV04) AddCase() *model.Case3 {
	d.Case = new(model.Case3)
	return d.Case
}

func (d *DuplicateV04) AddDuplicate() *model.ProprietaryData4 {
	d.Duplicate = new(model.ProprietaryData4)
	return d.Duplicate
}

func (d *DuplicateV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}
