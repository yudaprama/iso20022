package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100105 struct {
	XMLName xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.05 Document"`
	Message *StatusReportV05 `xml:"StsRpt"`
}

func (d *Document00100105) AddMessage() *StatusReportV05 {
	d.Message = new(StatusReportV05)
	return d.Message
}

// Informs the master terminal manager (MTM) or the terminal manager (TM) about the status of the acceptor system including the identification of the POI, its components and their installed versions.
type StatusReportV05 struct {

	// Set of characteristics related to the transfer of the status report.
	Header *model.Header27 `xml:"Hdr"`

	// Status of the point of interaction (POI), its components and their installed versions.
	StatusReport *model.StatusReport5 `xml:"StsRpt"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType12 `xml:"SctyTrlr,omitempty"`
}

func (s *StatusReportV05) AddHeader() *model.Header27 {
	s.Header = new(model.Header27)
	return s.Header
}

func (s *StatusReportV05) AddStatusReport() *model.StatusReport5 {
	s.StatusReport = new(model.StatusReport5)
	return s.StatusReport
}

func (s *StatusReportV05) AddSecurityTrailer() *model.ContentInformationType12 {
	s.SecurityTrailer = new(model.ContentInformationType12)
	return s.SecurityTrailer
}
