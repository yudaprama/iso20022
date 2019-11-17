package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.006.001.02 Document"`
	Message *AccountManagementStatusReportV02 `xml:"AcctMgmtStsRptV02"`
}

func (d *Document00600102) AddMessage() *AccountManagementStatusReportV02 {
	d.Message = new(AccountManagementStatusReportV02)
	return d.Message
}

// Scope
// An account servicer, eg, a registrar, transfer agent or custodian bank sends the AccountManagementStatusReport message to an account owner or its designated agent, eg, an investor to report on the receipt or the processing status of a previously received AccountOpeningInstruction or AccountModificationInstruction or GetAccountDetails message.
// Usage
// The AccountManagementStatusReport message is used to provide the processing status of a previously received AccountOpeningInstruction or of an AccountModificationInstruction message.
// The AccountManagementStatusReport message is also used by an account servicer to reject an AccountOpeningInstruction or AccountModificationInstruction or GetAccountDetails message when the message is not compliant with the agreed SLA or when the account cannot be uniquely identified.
// The account owner may report that the status of the instruction is either rejected, accepted, that the instruction is being processed or that the instruction has been forwarded to the next intermediary party for further processing.
type AccountManagementStatusReportV02 struct {

	// Identifies the message.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference3 `xml:"RltdRef"`

	// Status report details of an account opening instruction or account modification instruction that was previously received.
	StatusReport *model.AccountManagementStatusAndReason1 `xml:"StsRpt"`
}

func (a *AccountManagementStatusReportV02) AddMessageIdentification() *model.MessageIdentification1 {
	a.MessageIdentification = new(model.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountManagementStatusReportV02) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	a.RelatedReference = append(a.RelatedReference, newValue)
	return newValue
}

func (a *AccountManagementStatusReportV02) AddStatusReport() *model.AccountManagementStatusAndReason1 {
	a.StatusReport = new(model.AccountManagementStatusAndReason1)
	return a.StatusReport
}
