package model

// Choice between a document identification provided either by the account owner or by the account servicer.
type DocumentIdentification4Choice struct {

	// Identification of the document assigned by the account servicer.
	AccountServicerDocumentIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrDocId"`

	// Identification of the document assigned by the account owner.
	AccountOwnerDocumentIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrDocId"`
}

func (d *DocumentIdentification4Choice) SetAccountServicerDocumentIdentification(value string) {
	d.AccountServicerDocumentIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (d *DocumentIdentification4Choice) SetAccountOwnerDocumentIdentification(value string) {
	d.AccountOwnerDocumentIdentification = (*RestrictedFINXMax16Text)(&value)
}
