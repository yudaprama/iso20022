package model

// Choice between a document identification provided either by the account owner or by the account servicer.
type DocumentIdentification1Choice struct {

	// Identification of the document asigned by the account servicer.
	AccountServicerDocumentIdentification *Max35Text `xml:"AcctSvcrDocId"`

	// Identification of the document asigned by the account owner.
	AccountOwnerDocumentIdentification *Max35Text `xml:"AcctOwnrDocId"`
}

func (d *DocumentIdentification1Choice) SetAccountServicerDocumentIdentification(value string) {
	d.AccountServicerDocumentIdentification = (*Max35Text)(&value)
}

func (d *DocumentIdentification1Choice) SetAccountOwnerDocumentIdentification(value string) {
	d.AccountOwnerDocumentIdentification = (*Max35Text)(&value)
}
