package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200106 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.06 Document"`
	Message *CustomerPaymentStatusReportV06 `xml:"CstmrPmtStsRpt"`
}

func (d *Document00200106) AddMessage() *CustomerPaymentStatusReportV06 {
	d.Message = new(CustomerPaymentStatusReportV06)
	return d.Message
}

// Scope
// The CustomerPaymentStatusReport message is sent by an instructed agent to the previous party in the payment chain. It is used to inform this party about the positive or negative status of an instruction (either single or file). It is also used to report on a pending instruction.
// Usage
// The CustomerPaymentStatusReport message is exchanged between an agent and a non-financial institution customer to provide status information on instructions previously sent. Its usage will always be governed by a bilateral agreement between the agent and the non-financial institution customer.
// The CustomerPaymentStatusReport message can be used to provide information about the status (e.g. rejection, acceptance) of the initiation of a credit transfer, a direct debit, as well as on the initiation of other customer instructions.
// The CustomerPaymentStatusReport message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The CustomerPaymentStatusReport message can be used in domestic and cross-border scenarios.
type CustomerPaymentStatusReportV06 struct {

	// Set of characteristics shared by all individual transactions included in the status report message.
	GroupHeader *model.GroupHeader52 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus *model.OriginalGroupHeader1 `xml:"OrgnlGrpInfAndSts"`

	// Information concerning the original payment information, to which the status report message refers.
	OriginalPaymentInformationAndStatus []*model.OriginalPaymentInstruction12 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CustomerPaymentStatusReportV06) AddGroupHeader() *model.GroupHeader52 {
	c.GroupHeader = new(model.GroupHeader52)
	return c.GroupHeader
}

func (c *CustomerPaymentStatusReportV06) AddOriginalGroupInformationAndStatus() *model.OriginalGroupHeader1 {
	c.OriginalGroupInformationAndStatus = new(model.OriginalGroupHeader1)
	return c.OriginalGroupInformationAndStatus
}

func (c *CustomerPaymentStatusReportV06) AddOriginalPaymentInformationAndStatus() *model.OriginalPaymentInstruction12 {
	newValue := new(model.OriginalPaymentInstruction12)
	c.OriginalPaymentInformationAndStatus = append(c.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (c *CustomerPaymentStatusReportV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
