package model

// Provides information about the entitlement.
type Entitlement1 struct {

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Idenfitication of the account.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Provides information about the securities distribution.
	SecuritiesDistributionDetails []*SecuritiesEntitlement1 `xml:"SctiesDstrbtnDtls,omitempty"`

	// Provides information about the cash distribution.
	CashDistributionDetails []*CashEntitlement1 `xml:"CshDstrbtnDtls,omitempty"`
}

func (e *Entitlement1) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	e.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return e.AccountOwnerIdentification
}

func (e *Entitlement1) SetAccountIdentification(value string) {
	e.AccountIdentification = (*Max35Text)(&value)
}

func (e *Entitlement1) AddSecuritiesDistributionDetails() *SecuritiesEntitlement1 {
	newValue := new(SecuritiesEntitlement1)
	e.SecuritiesDistributionDetails = append(e.SecuritiesDistributionDetails, newValue)
	return newValue
}

func (e *Entitlement1) AddCashDistributionDetails() *CashEntitlement1 {
	newValue := new(CashEntitlement1)
	e.CashDistributionDetails = append(e.CashDistributionDetails, newValue)
	return newValue
}
