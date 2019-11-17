package model

// Specifies the details relative to the submission of a data set.
type RequiredSubmission2 struct {

	// Specifies with party(ies) is authorised to submit the data set as part of the transaction.
	Submitter []*BICIdentification1 `xml:"Submitr"`
}

func (r *RequiredSubmission2) AddSubmitter() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.Submitter = append(r.Submitter, newValue)
	return newValue
}
