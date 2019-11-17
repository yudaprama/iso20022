package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200103 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.03 Document"`
	Message *ManagementPlanReplacementV03 `xml:"MgmtPlanRplcmnt"`
}

func (d *Document00200103) AddMessage() *ManagementPlanReplacementV03 {
	d.Message = new(ManagementPlanReplacementV03)
	return d.Message
}

// Terminal maintenance actions to be performed by a point of interaction (POI).
type ManagementPlanReplacementV03 struct {

	// Set of characteristics related to the transfer of the management plan.
	Header *model.Header4 `xml:"Hdr"`

	// Sequence of terminal maintenance actions to be performed by a point of interaction (POI).
	ManagementPlan *model.ManagementPlan3 `xml:"MgmtPlan"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType9 `xml:"SctyTrlr"`
}

func (m *ManagementPlanReplacementV03) AddHeader() *model.Header4 {
	m.Header = new(model.Header4)
	return m.Header
}

func (m *ManagementPlanReplacementV03) AddManagementPlan() *model.ManagementPlan3 {
	m.ManagementPlan = new(model.ManagementPlan3)
	return m.ManagementPlan
}

func (m *ManagementPlanReplacementV03) AddSecurityTrailer() *model.ContentInformationType9 {
	m.SecurityTrailer = new(model.ContentInformationType9)
	return m.SecurityTrailer
}
