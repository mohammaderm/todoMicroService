package repository

var (

	// todo query
	limit          = 3
	createTodo     = `INSERT INTO todo (title, description, categoryid, accountid) VALUES (?, ?, ?, ?)`
	deleteTodo     = `delete from todo where id = ? and accountid = ?;`
	getAllTodo     = `select * from todo where accountid = ? order by created_at desc limit ? offset ?;`
	updateStatus   = `update todo set status = true where id = ? and accountid = ?;`
	updatePriority = `update todo set priority = ? where id = ? and accountid = ?;`
	updateDueDate  = `update todo set due_date = ? where id = ? and accountid = ?;`

	// category query
	createCategory = `insert into category (title, accountid) values (?, ?);`
	deleteCategory = `delete from category where id = ? and accountid = ?;`
	getAllCategory = `select * from category where accountid = ? order by created_at`
)
