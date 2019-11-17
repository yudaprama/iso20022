package model

// Reports to a corporate on the actual set up up of the account, related services and mandates.
type AccountReport15 struct {

	// Characteristics of the account.
	Account *CustomerAccount5 `xml:"Acct"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Specifies target and actual dates.
	ContractDates *AccountContract3 `xml:"CtrctDts,omitempty"`

	// Information specifying the account mandate.
	Mandate []*OperationMandate2 `xml:"Mndt,omitempty"`

	// Definition of a group of parties.
	Group []*Group1 `xml:"Grp,omitempty"`

	// Unique and unambiguous identification of the account used as a reference for the opening of another account.
	ReferenceAccount *CashAccount24 `xml:"RefAcct,omitempty"`

	// Unique and unambiguous identification of the account where to transfer the balance.
	BalanceTransferAccount *AccountForAction1 `xml:"BalTrfAcct,omitempty"`

	// Identification of  the transfer account servicer.
	TransferAccountServicerIdentification *BranchAndFinancialInstitutionIdentification5 `xml:"TrfAcctSvcrId,omitempty"`
}

func (a *AccountReport15) AddAccount() *CustomerAccount5 {
	a.Account = new(CustomerAccount5)
	return a.Account
}

func (a *AccountReport15) AddUnderlyingMasterAgreement() *ContractDocument1 {
	a.UnderlyingMasterAgreement = new(ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountReport15) AddContractDates() *AccountContract3 {
	a.ContractDates = new(AccountContract3)
	return a.ContractDates
}

func (a *AccountReport15) AddMandate() *OperationMandate2 {
	newValue := new(OperationMandate2)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountReport15) AddGroup() *Group1 {
	newValue := new(Group1)
	a.Group = append(a.Group, newValue)
	return newValue
}

func (a *AccountReport15) AddReferenceAccount() *CashAccount24 {
	a.ReferenceAccount = new(CashAccount24)
	return a.ReferenceAccount
}

func (a *AccountReport15) AddBalanceTransferAccount() *AccountForAction1 {
	a.BalanceTransferAccount = new(AccountForAction1)
	return a.BalanceTransferAccount
}

func (a *AccountReport15) AddTransferAccountServicerIdentification() *BranchAndFinancialInstitutionIdentification5 {
	a.TransferAccountServicerIdentification = new(BranchAndFinancialInstitutionIdentification5)
	return a.TransferAccountServicerIdentification
}
