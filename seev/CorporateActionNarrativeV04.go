package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03800104 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.04 Document"`
	Message *CorporateActionNarrativeV04 `xml:"CorpActnNrrtv"`
}

func (d *Document03800104) AddMessage() *CorporateActionNarrativeV04 {
	d.Message = new(CorporateActionNarrativeV04)
	return d.Message
}

// Scope
// The CorporateActionNarrative message is sent between an account servicer and an account owner or its designated agent to cater for tax reclaims, restrictions, documentation requirements. This message is bi-directional.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionNarrativeV04 struct {

	// General information about the safekeeping account and the account owner.
	AccountDetails *model.AccountIdentification33Choice `xml:"AcctDtls,omitempty"`

	// Provides information about the securitised right for entitlement.
	UnderlyingSecurity *model.SecurityIdentification19 `xml:"UndrlygScty,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation92 `xml:"CorpActnGnlInf"`

	// Provides additional information.
	AdditionalInformation *model.UpdatedAdditionalInformation8 `xml:"AddtlInf"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionNarrativeV04) AddAccountDetails() *model.AccountIdentification33Choice {
	c.AccountDetails = new(model.AccountIdentification33Choice)
	return c.AccountDetails
}

func (c *CorporateActionNarrativeV04) AddUnderlyingSecurity() *model.SecurityIdentification19 {
	c.UnderlyingSecurity = new(model.SecurityIdentification19)
	return c.UnderlyingSecurity
}

func (c *CorporateActionNarrativeV04) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation92 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation92)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionNarrativeV04) AddAdditionalInformation() *model.UpdatedAdditionalInformation8 {
	c.AdditionalInformation = new(model.UpdatedAdditionalInformation8)
	return c.AdditionalInformation
}

func (c *CorporateActionNarrativeV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
