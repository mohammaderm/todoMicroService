package types

type (
	CreateTodoReq struct {
		CategoryId  uint64 `json:"categoryId" example:"2"`
		Title       string `json:"title" example:"do homeWork"`
		Description string `json:"description" example:"solve all problems in chapter 2"`
	}
	CreateCategoryReq struct {
		Title string `json:"title" example:"Work"`
	}
	UpdatePriority struct {
		Priority int `json:"priority" example:"5"`
	}
	UpdateDueDate struct {
		DueDate string `json:"due_date" example:"2022-09-12 12:50:05"`
	}
)
