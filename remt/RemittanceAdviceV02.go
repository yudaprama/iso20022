package remt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100102 struct {
	XMLName xml.Name             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.02 Document"`
	Message *RemittanceAdviceV02 `xml:"RmtAdvc"`
}

func (d *Document00100102) AddMessage() *RemittanceAdviceV02 {
	d.Message = new(RemittanceAdviceV02)
	return d.Message
}

// The RemittanceAdvice message allows the originator to provide remittance details that can be associated with a payment.
type RemittanceAdviceV02 struct {

	// Set of characteristics shared by all remittance information included in the message.
	GroupHeader *model.GroupHeader62 `xml:"GrpHdr"`

	// Provides information to enable the matching of an entry with the items that the associated payment is intended to settle, such as commercial invoices in an accounts' receivable system, tax obligations, or garnishment orders.
	RemittanceInformation []*model.RemittanceInformation12 `xml:"RmtInf"`

	// Additional information that cannot be  captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RemittanceAdviceV02) AddGroupHeader() *model.GroupHeader62 {
	r.GroupHeader = new(model.GroupHeader62)
	return r.GroupHeader
}

func (r *RemittanceAdviceV02) AddRemittanceInformation() *model.RemittanceInformation12 {
	newValue := new(model.RemittanceInformation12)
	r.RemittanceInformation = append(r.RemittanceInformation, newValue)
	return newValue
}

func (r *RemittanceAdviceV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}
