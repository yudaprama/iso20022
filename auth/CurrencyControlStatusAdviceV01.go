package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02700101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Document"`
	Message *CurrencyControlStatusAdviceV01 `xml:"CcyCtrlStsAdvc"`
}

func (d *Document02700101) AddMessage() *CurrencyControlStatusAdviceV01 {
	d.Message = new(CurrencyControlStatusAdviceV01)
	return d.Message
}

// The CurrencyControlStatusAdvice message is sent by either the reporting party (respectively the registration agent or the registration agent (respectively the reporting party) to provide a status advice on a previously sent currency control message.
//
// Usage:
// The message may be sent in response to requests on the registration of  the currency control contract, supporting document or on the payment regulatory information notification.
type CurrencyControlStatusAdviceV01 struct {

	// Characteristics shared by all individual items included in the message.
	GroupHeader *model.CurrencyControlHeader2 `xml:"GrpHdr"`

	// Provides the group status for the global message.
	GroupStatus []*model.CurrencyControlGroupStatus1 `xml:"GrpSts"`

	// Provides the status of the package in the message, which may contain the individual records.
	PackageStatus []*model.CurrencyControlPackageStatus1 `xml:"PackgSts,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CurrencyControlStatusAdviceV01) AddGroupHeader() *model.CurrencyControlHeader2 {
	c.GroupHeader = new(model.CurrencyControlHeader2)
	return c.GroupHeader
}

func (c *CurrencyControlStatusAdviceV01) AddGroupStatus() *model.CurrencyControlGroupStatus1 {
	newValue := new(model.CurrencyControlGroupStatus1)
	c.GroupStatus = append(c.GroupStatus, newValue)
	return newValue
}

func (c *CurrencyControlStatusAdviceV01) AddPackageStatus() *model.CurrencyControlPackageStatus1 {
	newValue := new(model.CurrencyControlPackageStatus1)
	c.PackageStatus = append(c.PackageStatus, newValue)
	return newValue
}

func (c *CurrencyControlStatusAdviceV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
