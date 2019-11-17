package model

// Choice of status.
type Status26Choice struct {

	// Status report details of a bulk or multiple or switch order cancellation message.
	CancellationStatusReport *OrderStatusAndReason9 `xml:"CxlStsRpt"`

	// Status report details of one or more individual orders from a bulk or multiple or switch order cancellation request.
	IndividualCancellationStatusReport []*IndividualOrderStatusAndReason8 `xml:"IndvCxlStsRpt"`
}

func (s *Status26Choice) AddCancellationStatusReport() *OrderStatusAndReason9 {
	s.CancellationStatusReport = new(OrderStatusAndReason9)
	return s.CancellationStatusReport
}

func (s *Status26Choice) AddIndividualCancellationStatusReport() *IndividualOrderStatusAndReason8 {
	newValue := new(IndividualOrderStatusAndReason8)
	s.IndividualCancellationStatusReport = append(s.IndividualCancellationStatusReport, newValue)
	return newValue
}
