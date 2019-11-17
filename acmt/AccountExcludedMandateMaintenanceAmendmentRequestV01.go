package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600101 struct {
	XMLName xml.Name                                              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.016.001.01 Document"`
	Message *AccountExcludedMandateMaintenanceAmendmentRequestV01 `xml:"AcctExcldMndtMntncAmdmntReq"`
}

func (d *Document01600101) AddMessage() *AccountExcludedMandateMaintenanceAmendmentRequestV01 {
	d.Message = new(AccountExcludedMandateMaintenanceAmendmentRequestV01)
	return d.Message
}

// Scope
// The AccountExcludedMandateMaintenanceAmendmentRequest message is sent from an organisation to a financial institution as part of the account maintenance process. It is sent in response to a request from the financial institution to send additional information.
// Usage
// This update is about account details excluding any mandate information. The organisation will specify under the Account and Organisation tags the complete information as it should be in the financial institutions records after processing the update request.
type AccountExcludedMandateMaintenanceAmendmentRequestV01 struct {

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

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddReferences() *model.References4 {
	a.References = new(model.References4)
	return a.References
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddContractDates() *model.AccountContract2 {
	a.ContractDates = new(model.AccountContract2)
	return a.ContractDates
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddUnderlyingMasterAgreement() *model.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(model.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddAccount() *model.CustomerAccount1 {
	a.Account = new(model.CustomerAccount1)
	return a.Account
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddOrganisation() *model.Organisation6 {
	newValue := new(model.Organisation6)
	a.Organisation = append(a.Organisation, newValue)
	return newValue
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddDigitalSignature() *model.PartyAndSignature1 {
	newValue := new(model.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}
