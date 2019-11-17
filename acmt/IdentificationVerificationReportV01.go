package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02400101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.024.001.01 Document"`
	Message *IdentificationVerificationReportV01 `xml:"IdVrfctnRpt"`
}

func (d *Document02400101) AddMessage() *IdentificationVerificationReportV01 {
	d.Message = new(IdentificationVerificationReportV01)
	return d.Message
}

// Scope
// The IdentificationVerificationReport message is sent by an assigner to an assignee. It is used to confirm whether or not the presented party and/or account identification information is correct.
// Usage
// The IdentificationVerificationReport message is sent as a response to an IdentificationVerificationRequest message.
// The IdentificationVerificationReport message can contain one or more reports.
// The IdentificationVerificationReport message may include a reason if the presented party and/or account identification information is confirmed to be incorrect.
// The IdentificationVerificationReport message may include the correct party and/or account identification information.
type IdentificationVerificationReportV01 struct {

	// Identifies the identification assignment.
	Assignment *model.IdentificationAssignment1 `xml:"Assgnmt"`

	// Provides for the reference to the original identification assignment.
	OriginalAssignment *model.MessageIdentification4 `xml:"OrgnlAssgnmt,omitempty"`

	// Information concerning the verification of the identification data for which verification was requested.
	Report []*model.VerificationReport1 `xml:"Rpt"`
}

func (i *IdentificationVerificationReportV01) AddAssignment() *model.IdentificationAssignment1 {
	i.Assignment = new(model.IdentificationAssignment1)
	return i.Assignment
}

func (i *IdentificationVerificationReportV01) AddOriginalAssignment() *model.MessageIdentification4 {
	i.OriginalAssignment = new(model.MessageIdentification4)
	return i.OriginalAssignment
}

func (i *IdentificationVerificationReportV01) AddReport() *model.VerificationReport1 {
	newValue := new(model.VerificationReport1)
	i.Report = append(i.Report, newValue)
	return newValue
}
