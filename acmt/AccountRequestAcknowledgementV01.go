package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.010.001.01 Document"`
	Message *AccountRequestAcknowledgementV01 `xml:"AcctReqAck"`
}

func (d *Document01000101) AddMessage() *AccountRequestAcknowledgementV01 {
	d.Message = new(AccountRequestAcknowledgementV01)
	return d.Message
}

// Scope
// The AccountRequestAcknowledgement message is sent from a financial institution to an organisation. This message is sent in response to a request message from the organisation. It is sent after the request has been validated from an authentication and authorization point of view. The business content has not yet been validated at this stage.
// Usage
// This message should only be sent after the request has been validated from an authentication and authorization point of view.
type AccountRequestAcknowledgementV01 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References5 `xml:"Refs"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification []*model.AccountForAction1 `xml:"AcctId,omitempty"`

	// Identifier for an organisation.
	OrganisationIdentification []*model.OrganisationIdentification6 `xml:"OrgId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountRequestAcknowledgementV01) AddReferences() *model.References5 {
	a.References = new(model.References5)
	return a.References
}

func (a *AccountRequestAcknowledgementV01) AddAccountIdentification() *model.AccountForAction1 {
	newValue := new(model.AccountForAction1)
	a.AccountIdentification = append(a.AccountIdentification, newValue)
	return newValue
}

func (a *AccountRequestAcknowledgementV01) AddOrganisationIdentification() *model.OrganisationIdentification6 {
	newValue := new(model.OrganisationIdentification6)
	a.OrganisationIdentification = append(a.OrganisationIdentification, newValue)
	return newValue
}

func (a *AccountRequestAcknowledgementV01) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountRequestAcknowledgementV01) AddDigitalSignature() *model.PartyAndSignature1 {
	newValue := new(model.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}
