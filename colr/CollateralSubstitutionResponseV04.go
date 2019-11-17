package colr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100104 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.011.001.04 Document"`
	Message *CollateralSubstitutionResponseV04 `xml:"CollSbstitnRspn"`
}

func (d *Document01100104) AddMessage() *CollateralSubstitutionResponseV04 {
	d.Message = new(CollateralSubstitutionResponseV04)
	return d.Message
}

// Scope
// The CollateralSubstitutionResponse message is sent by either the collateral taker or its collateral manager to the collateral giver or its collateral manager. This is a response to the CollateralSubstitutionRequest message and the collateral proposed in the substitution request can be accepted or rejected. If the collateral proposed is rejected then a new CollateralSubstitutionRequest should be sent. Note: There are cases where the collateral giver will send this message when the collateral taker has initiated the CollateralSubstitutionRequest message, for example in case of breach in the concentration limit.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// This is a response to the CollateralSubstitutionRequest message and the collateral proposed in the substitution request can be accepted or rejected.
type CollateralSubstitutionResponseV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *model.Obligation4 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *model.Agreement4 `xml:"Agrmt,omitempty"`

	// Provides details about the collateral substitution response.
	SubstitutionResponse *model.SubstitutionResponse1 `xml:"SbstitnRspn"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CollateralSubstitutionResponseV04) SetTransactionIdentification(value string) {
	c.TransactionIdentification = (*model.Max35Text)(&value)
}

func (c *CollateralSubstitutionResponseV04) AddObligation() *model.Obligation4 {
	c.Obligation = new(model.Obligation4)
	return c.Obligation
}

func (c *CollateralSubstitutionResponseV04) AddAgreement() *model.Agreement4 {
	c.Agreement = new(model.Agreement4)
	return c.Agreement
}

func (c *CollateralSubstitutionResponseV04) AddSubstitutionResponse() *model.SubstitutionResponse1 {
	c.SubstitutionResponse = new(model.SubstitutionResponse1)
	return c.SubstitutionResponse
}

func (c *CollateralSubstitutionResponseV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
