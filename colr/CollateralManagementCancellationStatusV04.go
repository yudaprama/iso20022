package colr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600104 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.006.001.04 Document"`
	Message *CollateralManagementCancellationStatusV04 `xml:"CollMgmtCxlSts"`
}

func (d *Document00600104) AddMessage() *CollateralManagementCancellationStatusV04 {
	d.Message = new(CollateralManagementCancellationStatusV04)
	return d.Message
}

// Scope
// The CollateralManagementCancellationStatus message is sent by:
// - the collateral taker or its collateral manager to the collateral giver or its collateral manager, or
// - the collateral giver or its collateral manager to the collateral taker or its collateral manager.
// This message is used to provide the status of accepted or rejected for the CollateralManagementCancellationRequest message previously received.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The CollateralManagementCancellationStatus message can be sent to provide a status for a previously received CollateralManagementCancellationRequest message.
type CollateralManagementCancellationStatusV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Provides reference to the previous received message.
	Reference *model.Reference16 `xml:"Ref"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *model.Obligation4 `xml:"Oblgtn"`

	// Provides status details of the collateral cancellation message.
	CancellationStatus *model.CollateralCancellationStatus1 `xml:"CxlSts"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CollateralManagementCancellationStatusV04) SetTransactionIdentification(value string) {
	c.TransactionIdentification = (*model.Max35Text)(&value)
}

func (c *CollateralManagementCancellationStatusV04) AddReference() *model.Reference16 {
	c.Reference = new(model.Reference16)
	return c.Reference
}

func (c *CollateralManagementCancellationStatusV04) AddObligation() *model.Obligation4 {
	c.Obligation = new(model.Obligation4)
	return c.Obligation
}

func (c *CollateralManagementCancellationStatusV04) AddCancellationStatus() *model.CollateralCancellationStatus1 {
	c.CancellationStatus = new(model.CollateralCancellationStatus1)
	return c.CancellationStatus
}

func (c *CollateralManagementCancellationStatusV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
