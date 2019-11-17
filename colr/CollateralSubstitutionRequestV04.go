package colr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000104 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.010.001.04 Document"`
	Message *CollateralSubstitutionRequestV04 `xml:"CollSbstitnReq"`
}

func (d *Document01000104) AddMessage() *CollateralSubstitutionRequestV04 {
	d.Message = new(CollateralSubstitutionRequestV04)
	return d.Message
}

// Scope
// The CollateralSubstitutionRequest message is sent by either the collateral giver or its collateral manager to the collateral taker or its collateral manager. It is used to request a substitution of collateral by specifying the collateral to be returned and proposing the new type(s) of collateral to be delivered. Note: There are cases where the collateral taker can initiate the CollateralSubstitutionRequest message, for example in case of breach in the concentration limit.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The CollateralSubstitutionRequest message can be sent by either the collateral giver or collateral taker in order to substitute collateral that is already held by the other party.
type CollateralSubstitutionRequestV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *model.Obligation4 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *model.Agreement4 `xml:"Agrmt,omitempty"`

	// Provides details about the collateral that will be returned.
	CollateralSubstitutionReturn *model.CollateralSubstitution5 `xml:"CollSbstitnRtr"`

	// Provides details about the collateral that will be delivered.
	CollateralSubstitutionDeliver *model.CollateralSubstitution4 `xml:"CollSbstitnDlvr,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CollateralSubstitutionRequestV04) SetTransactionIdentification(value string) {
	c.TransactionIdentification = (*model.Max35Text)(&value)
}

func (c *CollateralSubstitutionRequestV04) AddObligation() *model.Obligation4 {
	c.Obligation = new(model.Obligation4)
	return c.Obligation
}

func (c *CollateralSubstitutionRequestV04) AddAgreement() *model.Agreement4 {
	c.Agreement = new(model.Agreement4)
	return c.Agreement
}

func (c *CollateralSubstitutionRequestV04) AddCollateralSubstitutionReturn() *model.CollateralSubstitution5 {
	c.CollateralSubstitutionReturn = new(model.CollateralSubstitution5)
	return c.CollateralSubstitutionReturn
}

func (c *CollateralSubstitutionRequestV04) AddCollateralSubstitutionDeliver() *model.CollateralSubstitution4 {
	c.CollateralSubstitutionDeliver = new(model.CollateralSubstitution4)
	return c.CollateralSubstitutionDeliver
}

func (c *CollateralSubstitutionRequestV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
