package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.01 Document"`
	Message *MaintenanceDelegationResponseV01 `xml:"MntncDlgtnRspn"`
}

func (d *Document00600101) AddMessage() *MaintenanceDelegationResponseV01 {
	d.Message = new(MaintenanceDelegationResponseV01)
	return d.Message
}

// The master terminal manager provides the outcome of a maintenance delegation request to a terminal manager.
type MaintenanceDelegationResponseV01 struct {

	// Maintenance delegation response message management information.
	Header *model.Header16 `xml:"Hdr"`

	// Information related to the request of maintenance delegations.
	MaintenanceDelegationResponse *model.MaintenanceDelegationResponse1 `xml:"MntncDlgtnRspn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType12 `xml:"SctyTrlr"`
}

func (m *MaintenanceDelegationResponseV01) AddHeader() *model.Header16 {
	m.Header = new(model.Header16)
	return m.Header
}

func (m *MaintenanceDelegationResponseV01) AddMaintenanceDelegationResponse() *model.MaintenanceDelegationResponse1 {
	m.MaintenanceDelegationResponse = new(model.MaintenanceDelegationResponse1)
	return m.MaintenanceDelegationResponse
}

func (m *MaintenanceDelegationResponseV01) AddSecurityTrailer() *model.ContentInformationType12 {
	m.SecurityTrailer = new(model.ContentInformationType12)
	return m.SecurityTrailer
}
