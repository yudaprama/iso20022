package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03500103 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.03 Document"`
	Message *ProprietaryFormatInvestigationV03 `xml:"PrtryFrmtInvstgtn"`
}

func (d *Document03500103) AddMessage() *ProprietaryFormatInvestigationV03 {
	d.Message = new(ProprietaryFormatInvestigationV03)
	return d.Message
}

// Scope
// The Proprietary Format Investigation message type is used by financial institutions, with their own offices, and/or with other financial institutions with which they have established bilateral agreements.
// Usage
// The user should ensure that an existing standard message cannot be used before using the proprietary message.
// As defined in the scope, this ProprietaryFormatInvestigation message may only be used when bilaterally agreed.
// It is used as an envelope for a non standard message and provides means to manage an exception or investigation which falls outside the scope or capability of any other formatted message.
// The ProprietaryData element must contain a well formed XML document. This means XML special characters such as '<' must be used in a way that is consistent with XML well-formedness criteria.
type ProprietaryFormatInvestigationV03 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage Rule: the Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment3 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *model.Case3 `xml:"Case"`

	// Proprietary information.
	ProprietaryData *model.ProprietaryData4 `xml:"PrtryData"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *ProprietaryFormatInvestigationV03) AddAssignment() *model.CaseAssignment3 {
	p.Assignment = new(model.CaseAssignment3)
	return p.Assignment
}

func (p *ProprietaryFormatInvestigationV03) AddCase() *model.Case3 {
	p.Case = new(model.Case3)
	return p.Case
}

func (p *ProprietaryFormatInvestigationV03) AddProprietaryData() *model.ProprietaryData4 {
	p.ProprietaryData = new(model.ProprietaryData4)
	return p.ProprietaryData
}

func (p *ProprietaryFormatInvestigationV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}
