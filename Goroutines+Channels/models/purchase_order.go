package models

type PurchaseOrder struct {
	Number int
	Value  float64
}

func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.Number = 1234
	// Save the purchase order to the database
	// and then send the purchase order back to the channel
	callback <- po
}