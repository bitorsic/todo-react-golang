import { useState } from "react"
import apiClient from "../config/axiosConfig"

interface TaskType {
	id: string,
	content: string
}

export interface TaskListType {
	id: string,
	title: string,
	tasks: TaskType[]
}

interface Props {
	obj: TaskListType,
}

const TaskListCard: React.FC<Props> = ({ obj }) => {
	const [clickedAdd, setClickedAdd] = useState<boolean>(false)
	const [taskInput, setTaskInput] = useState<string>("")
	const [tasks, setTasks] = useState<TaskType[]>(obj.tasks); // Store tasks in state

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { value } = e.target;
		setTaskInput(value);
	};

	const addNewTask = async () => {
		const newTask: TaskType = {
			id: "",
			content: taskInput,
		}

		try {
			const response = await apiClient.post("/api/tasks/" + obj.id, newTask)

			newTask.id = response.data.taskID
			setTasks((prevTasks) => [...prevTasks, newTask]);
			
			// Clear the input after adding the task
			setTaskInput("");
			setClickedAdd(false);
		} catch (error: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
			console.log(error)
			alert(error.response.data.message)
		}
	}

	return (
		<div className="max-w-md p-4 m-4 bg-white border border-gray-200 rounded-lg shadow sm:p-8 dark:bg-gray-800 dark:border-gray-700"
			id={obj.id}>
			<div className="flex items-center justify-between mb-4">
				<h5 className="text-2xl font-bold leading-none text-gray-900 dark:text-white">
					{obj.title}
				</h5>
				<button
					type="button"
					onClick={() => { setClickedAdd(!clickedAdd) }} // toggle the boolean
					className="text-white bg-blue-700 hover:bg-blue-800 font-medium rounded-lg text-sm p-2.5 text-center inline-flex items-center me-2 dark:bg-blue-600 dark:hover:bg-blue-700">
					{!clickedAdd ? ( // plus symbol when not clicked
						<svg className="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
							<path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14m-7 7V5" />
						</svg>
					) : ( // minus symbol when clicked
						<svg className="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
							<path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14" />
						</svg>
					)}
					<span className="sr-only">Add Task</span>
				</button>
			</div>
			<div className="flow-root">
				<ul role="list" className="divide-y divide-gray-200 dark:divide-gray-700">
					{tasks && tasks.map((task) => (
						<li key={task.id} className="py-3 sm:py-4 flex items-center">
							<p className="flex-1 min-w-0 ms-4 text-md font-medium text-gray-900 dark:text-white">
								{task.content}
							</p>
						</li>
					))}
					{clickedAdd && (
						// Code for the input field for new task
						<li className="pt-3 sm:pt-4 px-3 sm:px-4 items-center">
							<input
								type="text"
								name="test"
								id="test"
								value={taskInput}
								onChange={handleChange}
								className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
							/>
							<div className="flex justify-end">
								<button
									onClick={addNewTask}
									className="text-white bg-primary-600 hover:bg-primary-700 font-medium rounded-lg text-sm mt-4 mx-4 px-5 py-2 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
								>
									Add
								</button>
							</div>
						</li>
					)}
				</ul>
			</div>
		</div>
	)
}

export default TaskListCard