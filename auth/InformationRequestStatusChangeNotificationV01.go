package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300101 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.003.001.01 Document"`
	Message *InformationRequestStatusChangeNotificationV01 `xml:"InfReqStsChngNtfctn"`
}

func (d *Document00300101) AddMessage() *InformationRequestStatusChangeNotificationV01 {
	d.Message = new(InformationRequestStatusChangeNotificationV01)
	return d.Message
}

// This message is sent by the authorities (police, customs, tax authorities, enforcement authorities) to a financial institution to inform the financial institution that the confidentiality status of the investigation has changed.
type InformationRequestStatusChangeNotificationV01 struct {

	// Reference of the information request opening message that this message is an update of.
	OriginalBusinessQuery *model.Max35Text `xml:"OrgnlBizQry"`

	// Specifies the confidentiality status of the investigation.
	ConfidentialityStatus *model.YesNoIndicator `xml:"CnfdtltySts"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InformationRequestStatusChangeNotificationV01) SetOriginalBusinessQuery(value string) {
	i.OriginalBusinessQuery = (*model.Max35Text)(&value)
}

func (i *InformationRequestStatusChangeNotificationV01) SetConfidentialityStatus(value string) {
	i.ConfidentialityStatus = (*model.YesNoIndicator)(&value)
}

func (i *InformationRequestStatusChangeNotificationV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
