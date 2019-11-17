package colr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800104 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.008.001.04 Document"`
	Message *CollateralProposalResponseV04 `xml:"CollPrpslRspn"`
}

func (d *Document00800104) AddMessage() *CollateralProposalResponseV04 {
	d.Message = new(CollateralProposalResponseV04)
	return d.Message
}

// Scope
// The CollateralProposalResponse message is sent by the collateral taker or its collateral manager to the collateral giver or its collateral manager to either accept or reject the collateral which has been proposed for the margin call. This message applies to both initial and counter proposals. If the collateral proposal is rejected then a new collateral proposal must be made.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The CollateralProposalResponse message can be sent in response to a previously received CollateralProposal message in order to accept or reject the collateral that has been proposed to cover the margin call.
type CollateralProposalResponseV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *model.Obligation4 `xml:"Oblgtn"`

	// Details the response to the collateral which has been proposed for the margin call. The proposed collateral can be accepted or rejected.
	//
	ProposalResponse *model.CollateralProposalResponse2Choice `xml:"PrpslRspn"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CollateralProposalResponseV04) SetTransactionIdentification(value string) {
	c.TransactionIdentification = (*model.Max35Text)(&value)
}

func (c *CollateralProposalResponseV04) AddObligation() *model.Obligation4 {
	c.Obligation = new(model.Obligation4)
	return c.Obligation
}

func (c *CollateralProposalResponseV04) AddProposalResponse() *model.CollateralProposalResponse2Choice {
	c.ProposalResponse = new(model.CollateralProposalResponse2Choice)
	return c.ProposalResponse
}

func (c *CollateralProposalResponseV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
