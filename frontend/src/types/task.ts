export default interface Task {
	id: number;
	title: string;
	description: string;
	statue: boolean;
	due_date: Date;
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
	category: string;
	categories: string[];
	tasks: Task[];
	task: Task;
}
