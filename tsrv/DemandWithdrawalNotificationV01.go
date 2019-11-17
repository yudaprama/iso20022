package tsrv

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.017.001.01 Document"`
	Message *DemandWithdrawalNotificationV01 `xml:"DmndWdrwlNtfctn"`
}

func (d *Document01700101) AddMessage() *DemandWithdrawalNotificationV01 {
	d.Message = new(DemandWithdrawalNotificationV01)
	return d.Message
}

// The DemandWithdrawalNotification message is sent by the beneficiary to the party that issued the undertaking, either directly or via a presenting or nominated party, to inform the issuer or nominated party that it has elected to withdraw its demand under the demand guarantee or standby letter of credit.
type DemandWithdrawalNotificationV01 struct {

	// Details of the demand withdrawal notification.
	DemandWithdrawalNotificationDetails *model.UndertakingDemandWithdrawal1 `xml:"DmndWdrwlNtfctnDtls"`

	// Digital signature of the notification.
	DigitalSignature *model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (d *DemandWithdrawalNotificationV01) AddDemandWithdrawalNotificationDetails() *model.UndertakingDemandWithdrawal1 {
	d.DemandWithdrawalNotificationDetails = new(model.UndertakingDemandWithdrawal1)
	return d.DemandWithdrawalNotificationDetails
}

func (d *DemandWithdrawalNotificationV01) AddDigitalSignature() *model.PartyAndSignature2 {
	d.DigitalSignature = new(model.PartyAndSignature2)
	return d.DigitalSignature
}
