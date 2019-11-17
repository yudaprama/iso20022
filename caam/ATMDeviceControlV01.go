package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200101 struct {
	XMLName xml.Name             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.01 Document"`
	Message *ATMDeviceControlV01 `xml:"ATMDvcCtrl"`
}

func (d *Document00200101) AddMessage() *ATMDeviceControlV01 {
	d.Message = new(ATMDeviceControlV01)
	return d.Message
}

// The ATMDeviceControl message is sent by a maintenance host to an ATM in response to an ATMDeviceReport message. The message contains a sequence of maintenance commands the ATM must perform.
type ATMDeviceControlV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDeviceControl *model.ContentInformationType10 `xml:"PrtctdATMDvcCtrl,omitempty"`

	// Information related to the control of an ATM device.
	ATMDeviceControl *model.ATMDeviceControl1 `xml:"ATMDvcCtrl,omitempty"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType13 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDeviceControlV01) AddHeader() *model.Header20 {
	a.Header = new(model.Header20)
	return a.Header
}

func (a *ATMDeviceControlV01) AddProtectedATMDeviceControl() *model.ContentInformationType10 {
	a.ProtectedATMDeviceControl = new(model.ContentInformationType10)
	return a.ProtectedATMDeviceControl
}

func (a *ATMDeviceControlV01) AddATMDeviceControl() *model.ATMDeviceControl1 {
	a.ATMDeviceControl = new(model.ATMDeviceControl1)
	return a.ATMDeviceControl
}

func (a *ATMDeviceControlV01) AddSecurityTrailer() *model.ContentInformationType13 {
	a.SecurityTrailer = new(model.ContentInformationType13)
	return a.SecurityTrailer
}
