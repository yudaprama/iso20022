package model

// Choice of status.
type Status24Choice struct {

	// Status report details of all the individual orders conveyed in a bulk or multiple order message. Can be used if all the individual orders conveyed in a bulk or multiple order message have the same status.
	OrderDetailsReport *OrderStatusAndReason10 `xml:"OrdrDtlsRpt"`

	// Status report details of an individual order.
	IndividualOrderDetailsReport []*IndividualOrderStatusAndReason7 `xml:"IndvOrdrDtlsRpt"`

	// Status report details of a switch order.
	SwitchOrderDetailsReport []*SwitchOrderStatusAndReason2 `xml:"SwtchOrdrDtlsRpt"`
}

func (s *Status24Choice) AddOrderDetailsReport() *OrderStatusAndReason10 {
	s.OrderDetailsReport = new(OrderStatusAndReason10)
	return s.OrderDetailsReport
}

func (s *Status24Choice) AddIndividualOrderDetailsReport() *IndividualOrderStatusAndReason7 {
	newValue := new(IndividualOrderStatusAndReason7)
	s.IndividualOrderDetailsReport = append(s.IndividualOrderDetailsReport, newValue)
	return newValue
}

func (s *Status24Choice) AddSwitchOrderDetailsReport() *SwitchOrderStatusAndReason2 {
	newValue := new(SwitchOrderStatusAndReason2)
	s.SwitchOrderDetailsReport = append(s.SwitchOrderDetailsReport, newValue)
	return newValue
}
