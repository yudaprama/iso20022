package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.007.001.02 Document"`
	Message *AccountOpeningRequestV02 `xml:"AcctOpngReq"`
}

func (d *Document00700102) AddMessage() *AccountOpeningRequestV02 {
	d.Message = new(AccountOpeningRequestV02)
	return d.Message
}

// The AccountOpeningRequest message is sent from an organisation to a financial institution as part of the account opening process. It is the initial request message to open an account.
type AccountOpeningRequestV02 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References4 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution.
	From *model.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *model.CustomerAccount4 `xml:"Acct"`

	// Specifies target dates.
	ContractDates *model.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *model.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *model.Organisation12 `xml:"Org"`

	// Information specifying the account mandate.
	Mandate []*model.OperationMandate2 `xml:"Mndt,omitempty"`

	// Definition of a group of parties.
	Group []*model.Group1 `xml:"Grp,omitempty"`

	// Unique and unambiguous identification of the account used as a reference for the opening of another account.
	ReferenceAccount *model.CashAccount24 `xml:"RefAcct,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountOpeningRequestV02) AddReferences() *model.References4 {
	a.References = new(model.References4)
	return a.References
}

func (a *AccountOpeningRequestV02) AddFrom() *model.OrganisationIdentification8 {
	a.From = new(model.OrganisationIdentification8)
	return a.From
}

func (a *AccountOpeningRequestV02) AddAccount() *model.CustomerAccount4 {
	a.Account = new(model.CustomerAccount4)
	return a.Account
}

func (a *AccountOpeningRequestV02) AddContractDates() *model.AccountContract2 {
	a.ContractDates = new(model.AccountContract2)
	return a.ContractDates
}

func (a *AccountOpeningRequestV02) AddUnderlyingMasterAgreement() *model.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(model.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountOpeningRequestV02) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountOpeningRequestV02) AddOrganisation() *model.Organisation12 {
	a.Organisation = new(model.Organisation12)
	return a.Organisation
}

func (a *AccountOpeningRequestV02) AddMandate() *model.OperationMandate2 {
	newValue := new(model.OperationMandate2)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountOpeningRequestV02) AddGroup() *model.Group1 {
	newValue := new(model.Group1)
	a.Group = append(a.Group, newValue)
	return newValue
}

func (a *AccountOpeningRequestV02) AddReferenceAccount() *model.CashAccount24 {
	a.ReferenceAccount = new(model.CashAccount24)
	return a.ReferenceAccount
}

func (a *AccountOpeningRequestV02) AddDigitalSignature() *model.PartyAndSignature2 {
	newValue := new(model.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountOpeningRequestV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
