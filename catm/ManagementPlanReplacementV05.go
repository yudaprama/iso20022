package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200105 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Document"`
	Message *ManagementPlanReplacementV05 `xml:"MgmtPlanRplcmnt"`
}

func (d *Document00200105) AddMessage() *ManagementPlanReplacementV05 {
	d.Message = new(ManagementPlanReplacementV05)
	return d.Message
}

// Terminal maintenance actions to be performed by a point of interaction (POI).
type ManagementPlanReplacementV05 struct {

	// Set of characteristics related to the transfer of the management plan.
	Header *model.Header27 `xml:"Hdr"`

	// Sequence of terminal maintenance actions to be performed by a point of interaction (POI).
	ManagementPlan *model.ManagementPlan5 `xml:"MgmtPlan"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType12 `xml:"SctyTrlr,omitempty"`
}

func (m *ManagementPlanReplacementV05) AddHeader() *model.Header27 {
	m.Header = new(model.Header27)
	return m.Header
}

func (m *ManagementPlanReplacementV05) AddManagementPlan() *model.ManagementPlan5 {
	m.ManagementPlan = new(model.ManagementPlan5)
	return m.ManagementPlan
}

func (m *ManagementPlanReplacementV05) AddSecurityTrailer() *model.ContentInformationType12 {
	m.SecurityTrailer = new(model.ContentInformationType12)
	return m.SecurityTrailer
}
