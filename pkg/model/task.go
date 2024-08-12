package model

type Task struct {
	Id      string `json:"id,omitempty"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat"`
}

const (
	TimeFormat = "20060102"
)

type TasksResp struct {
	Tasks []Task `json:"tasks"`
}
