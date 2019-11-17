package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02000101 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.020.001.01 Document"`
	Message *AccountClosingAmendmentRequestV01 `xml:"AcctClsgAmdmntReq"`
}

func (d *Document02000101) AddMessage() *AccountClosingAmendmentRequestV01 {
	d.Message = new(AccountClosingAmendmentRequestV01)
	return d.Message
}

// Scope
// The AccountClosingAmendmentRequest message is sent from an organisation to a financial institution as part of the account closing process. It is sent in response to a request from the financial institution to send additional information.
// Usage
// This message may only be sent in response to a request from the financial institution to send additional information.
// It could be sent together with other related documents.
type AccountClosingAmendmentRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References4 `xml:"Refs"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification *model.AccountForAction1 `xml:"AcctId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Identification of the organisation requesting the change.
	OrganisationIdentification []*model.OrganisationIdentification6 `xml:"OrgId"`

	// Specifies target dates.
	ContractDates *model.AccountContract4 `xml:"CtrctDts,omitempty"`

	// Identification of the account to which the remaining positive balance of the account to be closed must be transferred or account from which funds can be moved to the account to be closed and which balance is negative. This account must be held in the same financial institution as the account to be closed if the transfer account is used to compensate a negative balance. For a positive balance to be transferred, an account in another financial institution might be used. In that case the account servicer is mandatory.
	BalanceTransferAccount *model.AccountForAction1 `xml:"BalTrfAcct,omitempty"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme, that is the servicer of the transfer account.
	TransferAccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification4 `xml:"TrfAcctSvcrId,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountClosingAmendmentRequestV01) AddReferences() *model.References4 {
	a.References = new(model.References4)
	return a.References
}

func (a *AccountClosingAmendmentRequestV01) AddAccountIdentification() *model.AccountForAction1 {
	a.AccountIdentification = new(model.AccountForAction1)
	return a.AccountIdentification
}

func (a *AccountClosingAmendmentRequestV01) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountClosingAmendmentRequestV01) AddOrganisationIdentification() *model.OrganisationIdentification6 {
	newValue := new(model.OrganisationIdentification6)
	a.OrganisationIdentification = append(a.OrganisationIdentification, newValue)
	return newValue
}

func (a *AccountClosingAmendmentRequestV01) AddContractDates() *model.AccountContract4 {
	a.ContractDates = new(model.AccountContract4)
	return a.ContractDates
}

func (a *AccountClosingAmendmentRequestV01) AddBalanceTransferAccount() *model.AccountForAction1 {
	a.BalanceTransferAccount = new(model.AccountForAction1)
	return a.BalanceTransferAccount
}

func (a *AccountClosingAmendmentRequestV01) AddTransferAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification4 {
	a.TransferAccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification4)
	return a.TransferAccountServicerIdentification
}

func (a *AccountClosingAmendmentRequestV01) AddDigitalSignature() *model.PartyAndSignature1 {
	newValue := new(model.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}
