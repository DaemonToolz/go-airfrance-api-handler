package main

import "reflect"

const (
	PRICE_DETAIL_REQUEST = "price_detail"
	TAX_BREAKDOWN        = "tax_breakdown"
	TICKET_CONDITION     = "ticket_condition"
)

var RequestRedirect map[string]reflect.Type

func initMappingRedirect() {
	RequestRedirect = make(map[string]reflect.Type)
	RequestRedirect[PRICE_DETAIL_REQUEST] = reflect.TypeOf(PriceDetailResult{})
	RequestRedirect[TAX_BREAKDOWN] = reflect.TypeOf(TaxBreakdownDetailResult{})
	RequestRedirect[TICKET_CONDITION] = reflect.TypeOf(TicketConditionsResponse{})
}
