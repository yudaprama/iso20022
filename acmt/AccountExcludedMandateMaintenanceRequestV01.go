package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500101 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Document"`
	Message *AccountExcludedMandateMaintenanceRequestV01 `xml:"AcctExcldMndtMntncReq"`
}

func (d *Document01500101) AddMessage() *AccountExcludedMandateMaintenanceRequestV01 {
	d.Message = new(AccountExcludedMandateMaintenanceRequestV01)
	return d.Message
}

// Scope
// This AccountExcludedMandateMaintenanceRequest message is sent from an organisation to a financial institution as part of the account maintenance process. It is the initial request message to update an account.
// Usage
// This update is about account details excluding any mandate information. The organisation will specify under the Account and Organisation tags the complete information as it should be in the financial institutions records after processing the update request.
type AccountExcludedMandateMaintenanceRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References4 `xml:"Refs"`

	// Specifies target dates.
	ContractDates *model.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the Group to which the organisation belongs, and the account Servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *model.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *model.CustomerAccount1 `xml:"Acct"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation []*model.Organisation6 `xml:"Org"`

	// Contains additional information related to the message.
	AdditionalMessageInformation *model.AdditionalInformation5 `xml:"AddtlMsgInf,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountExcludedMandateMaintenanceRequestV01) AddReferences() *model.References4 {
	a.References = new(model.References4)
	return a.References
}

func (a *AccountExcludedMandateMaintenanceRequestV01) AddContractDates() *model.AccountContract2 {
	a.ContractDates = new(model.AccountContract2)
	return a.ContractDates
}

func (a *AccountExcludedMandateMaintenanceRequestV01) AddUnderlyingMasterAgreement() *model.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(model.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountExcludedMandateMaintenanceRequestV01) AddAccount() *model.CustomerAccount1 {
	a.Account = new(model.CustomerAccount1)
	return a.Account
}

func (a *AccountExcludedMandateMaintenanceRequestV01) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountExcludedMandateMaintenanceRequestV01) AddOrganisation() *model.Organisation6 {
	newValue := new(model.Organisation6)
	a.Organisation = append(a.Organisation, newValue)
	return newValue
}

func (a *AccountExcludedMandateMaintenanceRequestV01) AddAdditionalMessageInformation() *model.AdditionalInformation5 {
	a.AdditionalMessageInformation = new(model.AdditionalInformation5)
	return a.AdditionalMessageInformation
}

func (a *AccountExcludedMandateMaintenanceRequestV01) AddDigitalSignature() *model.PartyAndSignature1 {
	newValue := new(model.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}
