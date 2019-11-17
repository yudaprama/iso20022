package model

// Choice between a document identification provided either by the account owner or by the account servicer.
type DocumentIdentification3Choice struct {

	// Identification of the document assigned by the account servicer.
	AccountServicerDocumentIdentification *Max35Text `xml:"AcctSvcrDocId"`

	// Identification of the document assigned by the account owner.
	AccountOwnerDocumentIdentification *Max35Text `xml:"AcctOwnrDocId"`
}

func (d *DocumentIdentification3Choice) SetAccountServicerDocumentIdentification(value string) {
	d.AccountServicerDocumentIdentification = (*Max35Text)(&value)
}

func (d *DocumentIdentification3Choice) SetAccountOwnerDocumentIdentification(value string) {
	d.AccountOwnerDocumentIdentification = (*Max35Text)(&value)
}
