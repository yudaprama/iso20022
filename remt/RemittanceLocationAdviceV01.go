package remt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200101 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.01 Document"`
	Message *RemittanceLocationAdviceV01 `xml:"RmtLctnAdvc"`
}

func (d *Document00200101) AddMessage() *RemittanceLocationAdviceV01 {
	d.Message = new(RemittanceLocationAdviceV01)
	return d.Message
}

// The RemittanceLocationAdvice message allows the originator of the message to identify where the remittance advice is located for a related payment.
type RemittanceLocationAdviceV01 struct {

	// Set of characteristics shared by all remittance location information included in the message.
	GroupHeader *model.GroupHeader62 `xml:"GrpHdr"`

	// Provides information related to location and/or delivery of the remittance information.  This information is used to enable the matching of an entry with the items that the associated payment is intended to settle.
	RemittanceLocation []*model.RemittanceLocation3 `xml:"RmtLctn"`

	// Additional information that cannot be  captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RemittanceLocationAdviceV01) AddGroupHeader() *model.GroupHeader62 {
	r.GroupHeader = new(model.GroupHeader62)
	return r.GroupHeader
}

func (r *RemittanceLocationAdviceV01) AddRemittanceLocation() *model.RemittanceLocation3 {
	newValue := new(model.RemittanceLocation3)
	r.RemittanceLocation = append(r.RemittanceLocation, newValue)
	return newValue
}

func (r *RemittanceLocationAdviceV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}
