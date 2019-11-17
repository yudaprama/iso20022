package model

// Specifies possible agreement clauses related to invoice financing.
type AgreementClauses1 struct {

	// Description of agreement clauses, given in a textual form.
	Description *Max256Text `xml:"Desc,omitempty"`

	// External reference to the document, containing agreement clauses, where it is stored.
	DocumentURL *Max350Text `xml:"DocURL,omitempty"`
}

func (a *AgreementClauses1) SetDescription(value string) {
	a.Description = (*Max256Text)(&value)
}

func (a *AgreementClauses1) SetDocumentURL(value string) {
	a.DocumentURL = (*Max350Text)(&value)
}
