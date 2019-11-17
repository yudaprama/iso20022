package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200101 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.012.001.01 Document"`
	Message *ATMDepositRequestV01 `xml:"ATMDpstReq"`
}

func (d *Document01200101) AddMessage() *ATMDepositRequestV01 {
	d.Message = new(ATMDepositRequestV01)
	return d.Message
}

// The ATMDepositRequest message is sent by an ATM to an acquirer or its agent to request the approval of a deposit transaction at an ATM, before or after deposit media inside the ATM.
type ATMDepositRequestV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDepositRequest *model.ContentInformationType10 `xml:"PrtctdATMDpstReq,omitempty"`

	// Information related to the request of a deposit transaction from an ATM.
	ATMDepositRequest *model.ATMDepositRequest1 `xml:"ATMDpstReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDepositRequestV01) AddHeader() *model.Header31 {
	a.Header = new(model.Header31)
	return a.Header
}

func (a *ATMDepositRequestV01) AddProtectedATMDepositRequest() *model.ContentInformationType10 {
	a.ProtectedATMDepositRequest = new(model.ContentInformationType10)
	return a.ProtectedATMDepositRequest
}

func (a *ATMDepositRequestV01) AddATMDepositRequest() *model.ATMDepositRequest1 {
	a.ATMDepositRequest = new(model.ATMDepositRequest1)
	return a.ATMDepositRequest
}

func (a *ATMDepositRequestV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
