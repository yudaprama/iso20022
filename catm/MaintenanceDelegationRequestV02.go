package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500102 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.005.001.02 Document"`
	Message *MaintenanceDelegationRequestV02 `xml:"MntncDlgtnReq"`
}

func (d *Document00500102) AddMessage() *MaintenanceDelegationRequestV02 {
	d.Message = new(MaintenanceDelegationRequestV02)
	return d.Message
}

// A terminal manager requests to the master terminal manager the delegation of maintenance functions or maintenance operation on the terminal estate managed by the master terminal manager.
type MaintenanceDelegationRequestV02 struct {

	// Information related to the protocol management.
	Header *model.Header29 `xml:"Hdr,omitempty"`

	// Information related to the request of maintenance delegations.
	MaintenanceDelegationRequest *model.MaintenanceDelegationRequest2 `xml:"MntncDlgtnReq"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType12 `xml:"SctyTrlr"`
}

func (m *MaintenanceDelegationRequestV02) AddHeader() *model.Header29 {
	m.Header = new(model.Header29)
	return m.Header
}

func (m *MaintenanceDelegationRequestV02) AddMaintenanceDelegationRequest() *model.MaintenanceDelegationRequest2 {
	m.MaintenanceDelegationRequest = new(model.MaintenanceDelegationRequest2)
	return m.MaintenanceDelegationRequest
}

func (m *MaintenanceDelegationRequestV02) AddSecurityTrailer() *model.ContentInformationType12 {
	m.SecurityTrailer = new(model.ContentInformationType12)
	return m.SecurityTrailer
}
