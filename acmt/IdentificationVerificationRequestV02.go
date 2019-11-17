package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02300102 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.023.001.02 Document"`
	Message *IdentificationVerificationRequestV02 `xml:"IdVrfctnReq"`
}

func (d *Document02300102) AddMessage() *IdentificationVerificationRequestV02 {
	d.Message = new(IdentificationVerificationRequestV02)
	return d.Message
}

// Scope
// The IdentificationVerificationRequest message is sent by an assigner to an assignee. It is used to request the verification of party and/or account identification information.
// Usage
// The IdentificationVerificationRequest message is sent before the sending of one or several transactions messages.
// The IdentificationVerificationRequest message can contain one or more verification requests.
type IdentificationVerificationRequestV02 struct {

	// Identifies the identification assignment.
	Assignment *model.IdentificationAssignment2 `xml:"Assgnmt"`

	// Information concerning the identification data that is requested to be verified.
	Verification []*model.IdentificationVerification2 `xml:"Vrfctn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IdentificationVerificationRequestV02) AddAssignment() *model.IdentificationAssignment2 {
	i.Assignment = new(model.IdentificationAssignment2)
	return i.Assignment
}

func (i *IdentificationVerificationRequestV02) AddVerification() *model.IdentificationVerification2 {
	newValue := new(model.IdentificationVerification2)
	i.Verification = append(i.Verification, newValue)
	return newValue
}

func (i *IdentificationVerificationRequestV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
