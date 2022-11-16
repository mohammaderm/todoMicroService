package repository

var (

	// todo query
	limit          = 3
	createTodo     = `INSERT INTO todo (title, description, categoryid, accountid) VALUES ($1, $2, $3, $4)`
	deleteTodo     = `delete from todo where id = $1 and accountid = $2;`
	getAllTodo     = `select * from todo where accountid = $1 order by created_at desc limit $2 offset $3;`
	updateStatus   = `update todo set status = true where id = $1 and accountid = $2;`
	updatePriority = `update todo set priority = $1 where id = $2 and accountid = $3;`
	updateDueDate  = `update todo set due_date = $1 where id = $2 and accountid = $3;`

	// category query
	createCategory = `insert into category (title, accountid) values ($1, $2);`
	deleteCategory = `delete from category where id = $1 and accountid = $2;`
	getAllCategory = `select * from category where accountid = $1 order by created_at`
)
