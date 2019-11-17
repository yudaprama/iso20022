package model

// Message used to report to a corporate on the actual set up up of the account, related services and mandates.
type AccountReport1 struct {

	// Characteristics of the account.
	Account *CustomerAccount1 `xml:"Acct"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Specifies target and actual dates.
	ContractDates *AccountContract3 `xml:"CtrctDts,omitempty"`

	// Information specifying the account mandate.
	Mandate []*OperationMandate1 `xml:"Mndt,omitempty"`

	// Unique and unambiguous identification of the account used as a reference for the opening of another account.
	ReferenceAccount *CashAccount16 `xml:"RefAcct,omitempty"`

	// Unique and unambiguous identification of the account where to transfer the balance.
	BalanceTransferAccount *AccountForAction1 `xml:"BalTrfAcct,omitempty"`

	// Identification of  the transfer account servicer.
	TransferAccountServicerIdentification *BranchAndFinancialInstitutionIdentification4 `xml:"TrfAcctSvcrId,omitempty"`
}

func (a *AccountReport1) AddAccount() *CustomerAccount1 {
	a.Account = new(CustomerAccount1)
	return a.Account
}

func (a *AccountReport1) AddUnderlyingMasterAgreement() *ContractDocument1 {
	a.UnderlyingMasterAgreement = new(ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountReport1) AddContractDates() *AccountContract3 {
	a.ContractDates = new(AccountContract3)
	return a.ContractDates
}

func (a *AccountReport1) AddMandate() *OperationMandate1 {
	newValue := new(OperationMandate1)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountReport1) AddReferenceAccount() *CashAccount16 {
	a.ReferenceAccount = new(CashAccount16)
	return a.ReferenceAccount
}

func (a *AccountReport1) AddBalanceTransferAccount() *AccountForAction1 {
	a.BalanceTransferAccount = new(AccountForAction1)
	return a.BalanceTransferAccount
}

func (a *AccountReport1) AddTransferAccountServicerIdentification() *BranchAndFinancialInstitutionIdentification4 {
	a.TransferAccountServicerIdentification = new(BranchAndFinancialInstitutionIdentification4)
	return a.TransferAccountServicerIdentification
}
