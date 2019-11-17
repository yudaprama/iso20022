package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200102 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.02 Document"`
	Message *ManagementPlanReplacementV02 `xml:"MgmtPlanRplcmnt"`
}

func (d *Document00200102) AddMessage() *ManagementPlanReplacementV02 {
	d.Message = new(ManagementPlanReplacementV02)
	return d.Message
}

// Terminal maintenance actions to be performed by a point of interaction (POI).
type ManagementPlanReplacementV02 struct {

	// Set of characteristics related to the transfer of the management plan.
	Header *model.Header4 `xml:"Hdr"`

	// Sequence of terminal maintenance actions to be performed by a point of interaction (POI).
	ManagementPlan *model.ManagementPlan2 `xml:"MgmtPlan"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType4 `xml:"SctyTrlr"`
}

func (m *ManagementPlanReplacementV02) AddHeader() *model.Header4 {
	m.Header = new(model.Header4)
	return m.Header
}

func (m *ManagementPlanReplacementV02) AddManagementPlan() *model.ManagementPlan2 {
	m.ManagementPlan = new(model.ManagementPlan2)
	return m.ManagementPlan
}

func (m *ManagementPlanReplacementV02) AddSecurityTrailer() *model.ContentInformationType4 {
	m.SecurityTrailer = new(model.ContentInformationType4)
	return m.SecurityTrailer
}
