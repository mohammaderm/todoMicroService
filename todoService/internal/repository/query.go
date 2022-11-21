package repository

var (

	// todo query
	limit       = 3
	createTodo  = `INSERT INTO todo (title, description, categoryid, accountid) VALUES ($1, $2, $3, $4) RETURNING id`
	deleteTodo  = `delete from todo where id = $1 and accountid = $2;`
	getAllTodo  = `select * from todo where accountid = $1 order by created_at desc limit $2 offset $3;`
	updateTodo  = `update todo set title = $1, categoryid = $2,description = $3,status = $4, due_date = $5, priority = $6 where id = $7 and accountid = $8;`
	getTodoById = `select * from todo where id = $1`

	// category query
	createCategory  = `insert into category (title, accountid) values ($1, $2) RETURNING id`
	deleteCategory  = `delete from category where id = $1 and accountid = $2;`
	getAllCategory  = `select * from category where accountid = $1 order by created_at`
	getCategoryById = `select * from category where id = $1`
)
