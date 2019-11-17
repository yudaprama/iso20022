package model

// Specifies an identification of a document assigned by and relative to the issuing party (of the identification).
// Optionally, the component can contain a copy of the identified document and a URI/URL (Universal Resource Information/Location) facilitating retrieval of the document.
// The component may also contain a cryptographic hash of the referenced document.
// Financial items are identified by three parts:
// (1) the creator of the document,
// (2) an identification of a dossier, and
// (3) an identification of a financial item.
// The two latter identifiers are independent permitting to identify the same item in several lists.
// The element identification is of schema type ID, it can be referenced by IDREF typed elements (composite=false).
type QualifiedDocumentInformation1 struct {

	// Local identification to be used in IDREFs in this message.
	Identification *ID `xml:"Id"`

	// Party issuing the reference.
	Issuer *QualifiedPartyIdentification1 `xml:"Issr,omitempty"`

	// Unambiguous identifier relative to the issuing party of a list of items.
	ItemListIdentifier *Max35Text `xml:"ItmListIdr,omitempty"`

	// Unambiguous identifier relative to the issuing party of an item (independent of any list).
	ItemIdentifier *Max35Text `xml:"ItmIdr,omitempty"`

	// Date of document or element. This may be used as a control value to indicate a specific version.
	Date *ISODate `xml:"Dt,omitempty"`

	// Identification of the version of the document or element. This may be used as a control value to indicate a specific version.
	Version *Max6Text `xml:"Vrsn,omitempty"`

	// If true, document is in its original form, otherwise it is a scanned version.
	ElectronicOriginal *YesNoIndicator `xml:"ElctrncOrgnl"`

	// Cryptographic hash of the document.
	Digest []*AlgorithmAndDigest1 `xml:"Dgst,omitempty"`

	// Specifies the type of the document, for example commercial invoice.
	DocumentType *ExternalDocumentType1Code `xml:"DocTp,omitempty"`

	// URL (Uniform Resource Locator) where the document can be found.
	URL *Max2048Text `xml:"URL,omitempty"`

	// Attached file for this document. The file must be in a self-describing format.
	AttachedFile []*BinaryFile1 `xml:"AttchdFile,omitempty"`
}

func (q *QualifiedDocumentInformation1) SetIdentification(value string) {
	q.Identification = (*ID)(&value)
}

func (q *QualifiedDocumentInformation1) AddIssuer() *QualifiedPartyIdentification1 {
	q.Issuer = new(QualifiedPartyIdentification1)
	return q.Issuer
}

func (q *QualifiedDocumentInformation1) SetItemListIdentifier(value string) {
	q.ItemListIdentifier = (*Max35Text)(&value)
}

func (q *QualifiedDocumentInformation1) SetItemIdentifier(value string) {
	q.ItemIdentifier = (*Max35Text)(&value)
}

func (q *QualifiedDocumentInformation1) SetDate(value string) {
	q.Date = (*ISODate)(&value)
}

func (q *QualifiedDocumentInformation1) SetVersion(value string) {
	q.Version = (*Max6Text)(&value)
}

func (q *QualifiedDocumentInformation1) SetElectronicOriginal(value string) {
	q.ElectronicOriginal = (*YesNoIndicator)(&value)
}

func (q *QualifiedDocumentInformation1) AddDigest() *AlgorithmAndDigest1 {
	newValue := new(AlgorithmAndDigest1)
	q.Digest = append(q.Digest, newValue)
	return newValue
}

func (q *QualifiedDocumentInformation1) SetDocumentType(value string) {
	q.DocumentType = (*ExternalDocumentType1Code)(&value)
}

func (q *QualifiedDocumentInformation1) SetURL(value string) {
	q.URL = (*Max2048Text)(&value)
}

func (q *QualifiedDocumentInformation1) AddAttachedFile() *BinaryFile1 {
	newValue := new(BinaryFile1)
	q.AttachedFile = append(q.AttachedFile, newValue)
	return newValue
}
