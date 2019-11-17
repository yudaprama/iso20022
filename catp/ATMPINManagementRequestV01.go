package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000101 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.01 Document"`
	Message *ATMPINManagementRequestV01 `xml:"ATMPINMgmtReq"`
}

func (d *Document01000101) AddMessage() *ATMPINManagementRequestV01 {
	d.Message = new(ATMPINManagementRequestV01)
	return d.Message
}

// The ATMPINManagementRequest message is sent by an ATM to an ATM manager to request an operation on the cardholder PIN.
type ATMPINManagementRequestV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMPINManagementRequest *model.ContentInformationType10 `xml:"PrtctdATMPINMgmtReq,omitempty"`

	// Information related to the request of a PIN management from an ATM.
	ATMPINManagementRequest *model.ATMPINManagementRequest1 `xml:"ATMPINMgmtReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMPINManagementRequestV01) AddHeader() *model.Header20 {
	a.Header = new(model.Header20)
	return a.Header
}

func (a *ATMPINManagementRequestV01) AddProtectedATMPINManagementRequest() *model.ContentInformationType10 {
	a.ProtectedATMPINManagementRequest = new(model.ContentInformationType10)
	return a.ProtectedATMPINManagementRequest
}

func (a *ATMPINManagementRequestV01) AddATMPINManagementRequest() *model.ATMPINManagementRequest1 {
	a.ATMPINManagementRequest = new(model.ATMPINManagementRequest1)
	return a.ATMPINManagementRequest
}

func (a *ATMPINManagementRequestV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
