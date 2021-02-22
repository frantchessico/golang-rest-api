package main

import("fmt")


type tasks struct {
	ID int  `json:ID`
	Name string `json:Name`
	Content string `json:Content`
}

type allTasks []tasks

var task = allTasks {
    {
		ID: 1,
		Name: "Task one",
		Content: "Some Content",
	},
}
func main() {
    fmt.Println(task)
}
 
 