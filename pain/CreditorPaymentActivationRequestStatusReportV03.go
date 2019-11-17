package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400103 struct {
	XMLName xml.Name                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.03 Document"`
	Message *CreditorPaymentActivationRequestStatusReportV03 `xml:"CdtrPmtActvtnReqStsRpt"`
}

func (d *Document01400103) AddMessage() *CreditorPaymentActivationRequestStatusReportV03 {
	d.Message = new(CreditorPaymentActivationRequestStatusReportV03)
	return d.Message
}

// The CreditorPaymentActivationRequestStatusReport message is sent by a party to the next party in the creditor payment activation request chain. It is used to inform the latter about the positive or negative status of a creditor payment activation request (either single or file).
type CreditorPaymentActivationRequestStatusReportV03 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader46 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus *model.OriginalGroupInformation25 `xml:"OrgnlGrpInfAndSts"`

	// Information concerning the original payment information, to which the status report message refers.
	OriginalPaymentInformationAndStatus []*model.OriginalPaymentInstruction9 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CreditorPaymentActivationRequestStatusReportV03) AddGroupHeader() *model.GroupHeader46 {
	c.GroupHeader = new(model.GroupHeader46)
	return c.GroupHeader
}

func (c *CreditorPaymentActivationRequestStatusReportV03) AddOriginalGroupInformationAndStatus() *model.OriginalGroupInformation25 {
	c.OriginalGroupInformationAndStatus = new(model.OriginalGroupInformation25)
	return c.OriginalGroupInformationAndStatus
}

func (c *CreditorPaymentActivationRequestStatusReportV03) AddOriginalPaymentInformationAndStatus() *model.OriginalPaymentInstruction9 {
	newValue := new(model.OriginalPaymentInstruction9)
	c.OriginalPaymentInformationAndStatus = append(c.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (c *CreditorPaymentActivationRequestStatusReportV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
