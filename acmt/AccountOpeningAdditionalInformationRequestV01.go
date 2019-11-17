package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900101 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.009.001.01 Document"`
	Message *AccountOpeningAdditionalInformationRequestV01 `xml:"AcctOpngAddtlInfReq"`
}

func (d *Document00900101) AddMessage() *AccountOpeningAdditionalInformationRequestV01 {
	d.Message = new(AccountOpeningAdditionalInformationRequestV01)
	return d.Message
}

// Scope
// The AccountOpeningAdditionalInformationRequest message is sent from a financial institution to an organisation as part of the account opening process. This message is sent in response to an opening request message from the organisation, if the business content is valid, but additional information is required.
// Usage
// This message should only be sent if additional information is required as part of the account opening process.
type AccountOpeningAdditionalInformationRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References3 `xml:"Refs"`

	// Identification of the organisation.
	OrganisationIdentification []*model.OrganisationIdentification6 `xml:"OrgId"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *model.CustomerAccount1 `xml:"Acct"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Account contract established between the organisation or the Group to which the organisation belongs, and the account Servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *model.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountOpeningAdditionalInformationRequestV01) AddReferences() *model.References3 {
	a.References = new(model.References3)
	return a.References
}

func (a *AccountOpeningAdditionalInformationRequestV01) AddOrganisationIdentification() *model.OrganisationIdentification6 {
	newValue := new(model.OrganisationIdentification6)
	a.OrganisationIdentification = append(a.OrganisationIdentification, newValue)
	return newValue
}

func (a *AccountOpeningAdditionalInformationRequestV01) AddAccount() *model.CustomerAccount1 {
	a.Account = new(model.CustomerAccount1)
	return a.Account
}

func (a *AccountOpeningAdditionalInformationRequestV01) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountOpeningAdditionalInformationRequestV01) AddUnderlyingMasterAgreement() *model.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(model.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountOpeningAdditionalInformationRequestV01) AddDigitalSignature() *model.PartyAndSignature1 {
	newValue := new(model.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}
