package asynchronous

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var baseUrl string = "https://jsonplaceholder.typicode.com/todos/"

// with my internet connection using go at the start of getTodo function is took around 500 ms on repeated tries (for all ids)
// without go keyword it took around 1.5-2 seconds on repeated tries so its a huge difference around 3-4 x faster (for all ids)
// with GOPROCSMAX set to 4 it took around 490 ms on repeated tries since my test are error prone this is basically have no difference

func GetTodos() {
	start := time.Now()
	ids := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	for _, id := range ids {

		go getTodoById(id)

	}

	elapsed := time.Since(start)

	defer fmt.Printf("Total time elapsed: %v\n ", elapsed)
}

// fetching single todo with http.Get

func getTodoById(id string) {
	start := time.Now()
	url := baseUrl + id

	resp, err := http.Get(url)

	if err != nil {
		println("Error occurred: %s", err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	todo := new(Todo)
	json.Unmarshal(body, &todo)

	elapsed := time.Since(start)
	fmt.Printf("Url : %s, took %v\n\n", url, elapsed)
	print(todo.String())
}
