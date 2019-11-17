package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900102 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.009.001.02 Document"`
	Message *AccountOpeningAdditionalInformationRequestV02 `xml:"AcctOpngAddtlInfReq"`
}

func (d *Document00900102) AddMessage() *AccountOpeningAdditionalInformationRequestV02 {
	d.Message = new(AccountOpeningAdditionalInformationRequestV02)
	return d.Message
}

// The AccountOpeningAdditionalInformationRequest message is sent from a financial institution to an organisation as part of the account opening process. This message is sent in response to an opening request message from the organisation, if the business content is valid, but additional information is required.
type AccountOpeningAdditionalInformationRequestV02 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References3 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution. OrganisationIdentification6
	From *model.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Identification of the organisation.
	OrganisationIdentification *model.OrganisationIdentification8 `xml:"OrgId"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *model.CustomerAccount4 `xml:"Acct"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *model.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddReferences() *model.References3 {
	a.References = new(model.References3)
	return a.References
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddFrom() *model.OrganisationIdentification8 {
	a.From = new(model.OrganisationIdentification8)
	return a.From
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddOrganisationIdentification() *model.OrganisationIdentification8 {
	a.OrganisationIdentification = new(model.OrganisationIdentification8)
	return a.OrganisationIdentification
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddAccount() *model.CustomerAccount4 {
	a.Account = new(model.CustomerAccount4)
	return a.Account
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddUnderlyingMasterAgreement() *model.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(model.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddDigitalSignature() *model.PartyAndSignature2 {
	newValue := new(model.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
