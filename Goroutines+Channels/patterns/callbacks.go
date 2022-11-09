package patterns

// in general callbacks involves two functions

// 1- long running function that takes second function as a parameter

// 2- second function  typicaly designed to receive the result of the first function

// callback pattern

/*
	purchaseOrder := new(models.PurchaseOrder)
	callback := make(chan *models.PurchaseOrder)
	go models.SavePO(purchaseOrder, callback)
	po := <-callback
	println(po.Number)
*/
