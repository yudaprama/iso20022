package colr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200104 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.012.001.04 Document"`
	Message *CollateralSubstitutionConfirmationV04 `xml:"CollSbstitnConf"`
}

func (d *Document01200104) AddMessage() *CollateralSubstitutionConfirmationV04 {
	d.Message = new(CollateralSubstitutionConfirmationV04)
	return d.Message
}

// Scope
// The CollateralSubstitutionConfirmation message is sent by:
// - the collateral taker or its collateral manager to the collateral giver or its collateral manager, or
// - the collateral giver or its collateral manager to the collateral taker or its collateral manager.
// This message confirms the collateral delivery. The collateral taker will only release the return of collateral when the new piece of collateral is received. The collateral giver sends the collateral taker the notification that the collateral substitution (that is new piece(s) of collateral) have been released. In the event that multiple pieces of collateral are being delivered in place of the collateral due to be returned by the giver, this message should only be generated once all collateral pieces have been agreed between both parties. Then the taker confirms that the collateral substitution (that is all pieces have been received) and acknowledges return of collateral.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// This message confirms the collateral delivery. The collateral taker will only release the return of collateral when the new piece of collateral is received. The collateral giver sends the collateral taker the notification that the collateral substitution (that is new piece(s) of collateral) have been released.
type CollateralSubstitutionConfirmationV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *model.Obligation4 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *model.Agreement4 `xml:"Agrmt,omitempty"`

	// Provides the status about the collateral substitution.
	SubstitutionConfirmation *model.CollateralConfirmation1 `xml:"SbstitnConf"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CollateralSubstitutionConfirmationV04) SetTransactionIdentification(value string) {
	c.TransactionIdentification = (*model.Max35Text)(&value)
}

func (c *CollateralSubstitutionConfirmationV04) AddObligation() *model.Obligation4 {
	c.Obligation = new(model.Obligation4)
	return c.Obligation
}

func (c *CollateralSubstitutionConfirmationV04) AddAgreement() *model.Agreement4 {
	c.Agreement = new(model.Agreement4)
	return c.Agreement
}

func (c *CollateralSubstitutionConfirmationV04) AddSubstitutionConfirmation() *model.CollateralConfirmation1 {
	c.SubstitutionConfirmation = new(model.CollateralConfirmation1)
	return c.SubstitutionConfirmation
}

func (c *CollateralSubstitutionConfirmationV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
