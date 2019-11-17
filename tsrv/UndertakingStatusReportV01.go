package tsrv

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01900101 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.019.001.01 Document"`
	Message *UndertakingStatusReportV01 `xml:"UdrtkgStsRpt"`
}

func (d *Document01900101) AddMessage() *UndertakingStatusReportV01 {
	d.Message = new(UndertakingStatusReportV01)
	return d.Message
}

// The UndertakingStatusReport message is exchanged between parties that have an interest in the referenced undertaking transaction. It notifies the recipient of the status of the transaction, such as acceptance or rejection, withdrawal, or non-conformation. The sender may add additional information, as appropriate.
type UndertakingStatusReportV01 struct {

	// Details of the undertaking status report
	UndertakingStatusReportDetails *model.UndertakingStatusAdvice1 `xml:"UdrtkgStsRptDtls"`

	// Digital signature of the report.
	DigitalSignature *model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingStatusReportV01) AddUndertakingStatusReportDetails() *model.UndertakingStatusAdvice1 {
	u.UndertakingStatusReportDetails = new(model.UndertakingStatusAdvice1)
	return u.UndertakingStatusReportDetails
}

func (u *UndertakingStatusReportV01) AddDigitalSignature() *model.PartyAndSignature2 {
	u.DigitalSignature = new(model.PartyAndSignature2)
	return u.DigitalSignature
}
