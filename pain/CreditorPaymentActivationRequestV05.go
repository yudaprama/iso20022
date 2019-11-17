package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300105 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.05 Document"`
	Message *CreditorPaymentActivationRequestV05 `xml:"CdtrPmtActvtnReq"`
}

func (d *Document01300105) AddMessage() *CreditorPaymentActivationRequestV05 {
	d.Message = new(CreditorPaymentActivationRequestV05)
	return d.Message
}

// The CreditorPaymentActivationRequest message is sent by the Creditor sending party to the Debtor receiving party, directly or through agents. It is used by a Creditor to request movement of funds from the debtor account to a creditor.
type CreditorPaymentActivationRequestV05 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader45 `xml:"GrpHdr"`

	// Set of characteristics that applies to the debit side of the payment transactions included in the creditor payment initiation.
	PaymentInformation []*model.PaymentInstruction19 `xml:"PmtInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CreditorPaymentActivationRequestV05) AddGroupHeader() *model.GroupHeader45 {
	c.GroupHeader = new(model.GroupHeader45)
	return c.GroupHeader
}

func (c *CreditorPaymentActivationRequestV05) AddPaymentInformation() *model.PaymentInstruction19 {
	newValue := new(model.PaymentInstruction19)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}

func (c *CreditorPaymentActivationRequestV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
