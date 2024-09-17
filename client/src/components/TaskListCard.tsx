import { useState } from "react"
import { useAuth } from "../hooks/useAuth"
import { useAxios } from "../hooks/useAxios"

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
	const { authUser } = useAuth()
	const { apiReq } = useAxios()

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setTaskInput(e.target.value);
	};

	const addNewTask = async () => {
		const newTask: TaskType = {
			id: "",
			content: taskInput,
		}

		const response = await apiReq<{ taskID: string }, TaskType>(
			"post",
			"/api/tasks/" + obj.id,
			newTask,
			authUser?.authToken,
		)

		if (response) {
			newTask.id = response.data.taskID

			// handling empty array
			if (tasks) {
				setTasks((prevTasks) => [...prevTasks, newTask]);
			} else {
				setTasks([newTask])
			}

			// Clear the input after adding the task
			setTaskInput("");
			setClickedAdd(false);
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
							<path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M5 12h14m-7 7V5" />
						</svg>
					) : ( // minus symbol when clicked
						<svg className="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
							<path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M5 12h14" />
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
						<li className="pt-3 sm:pt-4 px-1.5 sm:px-2 items-center">
							<input
								type="text"
								name="test"
								id="test"
								placeholder="Add New Task"
								value={taskInput}
								onChange={handleChange}
								className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
							/>
							<div className="flex justify-end">
								<button
									onClick={addNewTask}
									className="text-white bg-primary-600 hover:bg-primary-700 font-medium rounded-lg text-sm mt-4 mx-2 px-5 py-2 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
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