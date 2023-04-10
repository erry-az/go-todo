package todo

type Response[T any] struct {
	Data T `json:"data,omitempty"`
	OK   bool
}

type GetRequest struct {
	ID uint `param:"id"`
}

type CreateRequest struct {
	TaskName    string `form:"task_name"`
	Description string `form:"description"`
	IsDone      bool   `form:"is_done"`
}
