import taskType from "./task";

export default interface Category {
	id: number;
	title: string;
	created_at: Date;
	tasks: taskType[];
}

export interface taskAction {
	method:
		| "addCategories"
		| "replaceCategories"
		| "deleteCategory"
		| "addTasks"
		| "replaceTasks"
		| "updateTask"
		| "deleteTask";
	category: Category;
	categories: Category[];
	tasks: taskType[];
	task: taskType;
}
