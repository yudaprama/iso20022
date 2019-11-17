package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02400102 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.024.001.02 Document"`
	Message *IdentificationVerificationReportV02 `xml:"IdVrfctnRpt"`
}

func (d *Document02400102) AddMessage() *IdentificationVerificationReportV02 {
	d.Message = new(IdentificationVerificationReportV02)
	return d.Message
}

// Scope
// The IdentificationVerificationReport message is sent by an assigner to an assignee. It is used to confirm whether or not the presented party and/or account identification information is correct.
// Usage
// The IdentificationVerificationReport message is sent as a response to an IdentificationVerificationRequest message.
// The IdentificationVerificationReport message can contain one or more reports.
// The IdentificationVerificationReport message may include a reason if the presented party and/or account identification information is confirmed to be incorrect.
// The IdentificationVerificationReport message may include the correct party and/or account identification information.
type IdentificationVerificationReportV02 struct {

	// Identifies the identification assignment.
	Assignment *model.IdentificationAssignment2 `xml:"Assgnmt"`

	// Provides for the reference to the original identification assignment.
	OriginalAssignment *model.MessageIdentification5 `xml:"OrgnlAssgnmt,omitempty"`

	// Information concerning the verification of the identification data for which verification was requested.
	Report []*model.VerificationReport2 `xml:"Rpt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IdentificationVerificationReportV02) AddAssignment() *model.IdentificationAssignment2 {
	i.Assignment = new(model.IdentificationAssignment2)
	return i.Assignment
}

func (i *IdentificationVerificationReportV02) AddOriginalAssignment() *model.MessageIdentification5 {
	i.OriginalAssignment = new(model.MessageIdentification5)
	return i.OriginalAssignment
}

func (i *IdentificationVerificationReportV02) AddReport() *model.VerificationReport2 {
	newValue := new(model.VerificationReport2)
	i.Report = append(i.Report, newValue)
	return newValue
}

func (i *IdentificationVerificationReportV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
