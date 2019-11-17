package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400101 struct {
	XMLName xml.Name                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01 Document"`
	Message *CreditorPaymentActivationRequestStatusReportV01 `xml:"CdtrPmtActvtnReqStsRpt"`
}

func (d *Document01400101) AddMessage() *CreditorPaymentActivationRequestStatusReportV01 {
	d.Message = new(CreditorPaymentActivationRequestStatusReportV01)
	return d.Message
}

// Scope
// This message is sent by a party to the next party in the creditor payment activation request chain.
// It is used to inform the latter about the positive or negative status of a creditor payment activation request (either single or file).
type CreditorPaymentActivationRequestStatusReportV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader46 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus *model.OriginalGroupInformation25 `xml:"OrgnlGrpInfAndSts"`

	// Information concerning the original payment information, to which the status report message refers.
	OriginalPaymentInformationAndStatus []*model.OriginalPaymentInformation5 `xml:"OrgnlPmtInfAndSts,omitempty"`
}

func (c *CreditorPaymentActivationRequestStatusReportV01) AddGroupHeader() *model.GroupHeader46 {
	c.GroupHeader = new(model.GroupHeader46)
	return c.GroupHeader
}

func (c *CreditorPaymentActivationRequestStatusReportV01) AddOriginalGroupInformationAndStatus() *model.OriginalGroupInformation25 {
	c.OriginalGroupInformationAndStatus = new(model.OriginalGroupInformation25)
	return c.OriginalGroupInformationAndStatus
}

func (c *CreditorPaymentActivationRequestStatusReportV01) AddOriginalPaymentInformationAndStatus() *model.OriginalPaymentInformation5 {
	newValue := new(model.OriginalPaymentInformation5)
	c.OriginalPaymentInformationAndStatus = append(c.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}
