package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700101 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.007.001.01 Document"`
	Message *AccountOpeningRequestV01 `xml:"AcctOpngReq"`
}

func (d *Document00700101) AddMessage() *AccountOpeningRequestV01 {
	d.Message = new(AccountOpeningRequestV01)
	return d.Message
}

// Scope
// The AccountOpeningRequest message is sent from an organisation to a financial institution as part of the account opening process. It is the initial request to open an account.
// Usage
// This message should only be used for initiating the account opening process.
// It could be sent together with other related documents.
type AccountOpeningRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References4 `xml:"Refs"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *model.CustomerAccount1 `xml:"Acct"`

	// Specifies target dates.
	ContractDates *model.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the Group to which the organisation belongs, and the account Servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *model.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation []*model.Organisation7 `xml:"Org"`

	// Information specifying the account mandate.
	Mandate []*model.OperationMandate1 `xml:"Mndt,omitempty"`

	// Unique and unambiguous identification of the account used as a reference for the opening of another account.
	ReferenceAccount *model.CashAccount16 `xml:"RefAcct,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountOpeningRequestV01) AddReferences() *model.References4 {
	a.References = new(model.References4)
	return a.References
}

func (a *AccountOpeningRequestV01) AddAccount() *model.CustomerAccount1 {
	a.Account = new(model.CustomerAccount1)
	return a.Account
}

func (a *AccountOpeningRequestV01) AddContractDates() *model.AccountContract2 {
	a.ContractDates = new(model.AccountContract2)
	return a.ContractDates
}

func (a *AccountOpeningRequestV01) AddUnderlyingMasterAgreement() *model.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(model.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountOpeningRequestV01) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountOpeningRequestV01) AddOrganisation() *model.Organisation7 {
	newValue := new(model.Organisation7)
	a.Organisation = append(a.Organisation, newValue)
	return newValue
}

func (a *AccountOpeningRequestV01) AddMandate() *model.OperationMandate1 {
	newValue := new(model.OperationMandate1)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountOpeningRequestV01) AddReferenceAccount() *model.CashAccount16 {
	a.ReferenceAccount = new(model.CashAccount16)
	return a.ReferenceAccount
}

func (a *AccountOpeningRequestV01) AddDigitalSignature() *model.PartyAndSignature1 {
	newValue := new(model.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}
