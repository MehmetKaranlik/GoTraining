package patterns

import "mhmtkrnlk.com/goroutines_channels/models"

// in general callbacks involves two functions

// 1- long running function that takes second function as a parameter

// 2- second function  typicaly designed to receive the result of the first function

// callback pattern

var PurchaseOrder = new(models.PurchaseOrder)
var Callback = make(chan *models.PurchaseOrder)

func SavePo() {
	go func() {
		go models.SavePO(PurchaseOrder, Callback)
		po := <-Callback
		println(po.Number)
	}()
}
