package asynchronous

import "fmt"

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t *Todo) String() string {
	return fmt.Sprintf(
		"TODO : { \n %v,\n %v,\n %v,\n %v,\n}\n\n", t.UserId, t.Id, t.Title, t.Completed,
	)
}
