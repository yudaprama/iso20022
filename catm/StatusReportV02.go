package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100102 struct {
	XMLName xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Document"`
	Message *StatusReportV02 `xml:"StsRpt"`
}

func (d *Document00100102) AddMessage() *StatusReportV02 {
	d.Message = new(StatusReportV02)
	return d.Message
}

// Informs the master terminal manager (MTM) or the terminal manager (TM) about the status of the acceptor system including the identification of the POI, its components and their installed versions.
type StatusReportV02 struct {

	// Set of characteristics related to the transfer of the status report.
	Header *model.Header4 `xml:"Hdr"`

	// Status of the point of interaction (POI), its components and their installed versions.
	StatusReport *model.StatusReport2 `xml:"StsRpt"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType4 `xml:"SctyTrlr"`
}

func (s *StatusReportV02) AddHeader() *model.Header4 {
	s.Header = new(model.Header4)
	return s.Header
}

func (s *StatusReportV02) AddStatusReport() *model.StatusReport2 {
	s.StatusReport = new(model.StatusReport2)
	return s.StatusReport
}

func (s *StatusReportV02) AddSecurityTrailer() *model.ContentInformationType4 {
	s.SecurityTrailer = new(model.ContentInformationType4)
	return s.SecurityTrailer
}
