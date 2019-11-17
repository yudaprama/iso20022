package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400102 struct {
	XMLName xml.Name                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.02 Document"`
	Message *CreditorPaymentActivationRequestStatusReportV02 `xml:"CdtrPmtActvtnReqStsRpt"`
}

func (d *Document01400102) AddMessage() *CreditorPaymentActivationRequestStatusReportV02 {
	d.Message = new(CreditorPaymentActivationRequestStatusReportV02)
	return d.Message
}

// Scope
// This message is sent by a party to the next party in the creditor payment activation request chain.
// It is used to inform the latter about the positive or negative status of a creditor payment activation request (either single or file).
type CreditorPaymentActivationRequestStatusReportV02 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader46 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus *model.OriginalGroupInformation25 `xml:"OrgnlGrpInfAndSts"`

	// Information concerning the original payment information, to which the status report message refers.
	OriginalPaymentInformationAndStatus []*model.OriginalPaymentInstruction5 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CreditorPaymentActivationRequestStatusReportV02) AddGroupHeader() *model.GroupHeader46 {
	c.GroupHeader = new(model.GroupHeader46)
	return c.GroupHeader
}

func (c *CreditorPaymentActivationRequestStatusReportV02) AddOriginalGroupInformationAndStatus() *model.OriginalGroupInformation25 {
	c.OriginalGroupInformationAndStatus = new(model.OriginalGroupInformation25)
	return c.OriginalGroupInformationAndStatus
}

func (c *CreditorPaymentActivationRequestStatusReportV02) AddOriginalPaymentInformationAndStatus() *model.OriginalPaymentInstruction5 {
	newValue := new(model.OriginalPaymentInstruction5)
	c.OriginalPaymentInformationAndStatus = append(c.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (c *CreditorPaymentActivationRequestStatusReportV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
