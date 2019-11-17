package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100101 struct {
	XMLName xml.Name            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.001.001.01 Document"`
	Message *ATMDeviceReportV01 `xml:"ATMDvcRpt"`
}

func (d *Document00100101) AddMessage() *ATMDeviceReportV01 {
	d.Message = new(ATMDeviceReportV01)
	return d.Message
}

// The ATMDeviceReport message is sent to an acquirer by an ATM, or forwarded by an agent, to report:
// - The result of maintenance commands performed by the ATM,
// - The components of the ATM,
// - The status of the ATM components
type ATMDeviceReportV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDeviceReport *model.ContentInformationType10 `xml:"PrtctdATMDvcRpt,omitempty"`

	// Information related to the status report from an ATM device.
	ATMDeviceReport *model.ATMDeviceReport1 `xml:"ATMDvcRpt,omitempty"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType13 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDeviceReportV01) AddHeader() *model.Header20 {
	a.Header = new(model.Header20)
	return a.Header
}

func (a *ATMDeviceReportV01) AddProtectedATMDeviceReport() *model.ContentInformationType10 {
	a.ProtectedATMDeviceReport = new(model.ContentInformationType10)
	return a.ProtectedATMDeviceReport
}

func (a *ATMDeviceReportV01) AddATMDeviceReport() *model.ATMDeviceReport1 {
	a.ATMDeviceReport = new(model.ATMDeviceReport1)
	return a.ATMDeviceReport
}

func (a *ATMDeviceReportV01) AddSecurityTrailer() *model.ContentInformationType13 {
	a.SecurityTrailer = new(model.ContentInformationType13)
	return a.SecurityTrailer
}
