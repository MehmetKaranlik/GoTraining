package main

func main() {
	// Event pattern is a very common pattern in user interface frameworks.
	/*
		btn := models.MakeButton()

		handler1 := make(chan string)
		handler2 := make(chan string)

		btn.AddEventListener("click", handler1)
		btn.AddEventListener("example", handler2)
		go func() {
			for {
				msg := <-handler1
				msg2 := <-handler2
				fmt.Printf("Handler 1: %s\n", msg)
				fmt.Printf("Handler 2: %s\n", msg2)
			}
		}()

		btn.TriggerEvent("click", "Hello World!")
		btn.TriggerEvent("example", "Example Output")
		fmt.Scanln()
	*/

	// callback pattern

	/*
		purchaseOrder := new(models.PurchaseOrder)
		callback := make(chan *models.PurchaseOrder)
		go models.SavePO(purchaseOrder, callback)
		po := <-callback
		println(po.Number)
	*/
}
