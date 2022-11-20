export default interface Task {
	id: number;
	title: string;
	description: string;
	status: boolean;
	due_date: Date;
}
