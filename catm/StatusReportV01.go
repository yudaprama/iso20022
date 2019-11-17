package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100101 struct {
	XMLName xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.01 Document"`
	Message *StatusReportV01 `xml:"StsRpt"`
}

func (d *Document00100101) AddMessage() *StatusReportV01 {
	d.Message = new(StatusReportV01)
	return d.Message
}

// Scope
// The StatusReport message is sent by the card acceptor to the master terminal manager or delegated terminal manager to inform these entities about the status of the acceptor system.
// Usage
// The StatusReport message may embed the information required by the master terminal manager or delegated terminal manager for preparing the needed TMS actions in the management plan to be replaced and the acceptor parameters to be updated.
type StatusReportV01 struct {

	// Set of characteristics related to the transfer of the status report.
	Header *model.Header4 `xml:"Hdr"`

	// Status of the point of interaction (POI), its components and their installed versions.
	StatusReport *model.StatusReport1 `xml:"StsRpt"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType1 `xml:"SctyTrlr"`
}

func (s *StatusReportV01) AddHeader() *model.Header4 {
	s.Header = new(model.Header4)
	return s.Header
}

func (s *StatusReportV01) AddStatusReport() *model.StatusReport1 {
	s.StatusReport = new(model.StatusReport1)
	return s.StatusReport
}

func (s *StatusReportV01) AddSecurityTrailer() *model.ContentInformationType1 {
	s.SecurityTrailer = new(model.ContentInformationType1)
	return s.SecurityTrailer
}
