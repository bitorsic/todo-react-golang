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
	deleteHandler: (taskListID: string) => Promise<void>,
	deleteLoadingState: boolean,
}

const TaskListCard: React.FC<Props> = ({ obj, deleteHandler, deleteLoadingState }) => {
	const [clickedAdd, setClickedAdd] = useState<boolean>(false);
	const [taskInput, setTaskInput] = useState<string>("");
	const [tasks, setTasks] = useState<TaskType[]>(obj.tasks); // Store tasks in state
	const { authUser } = useAuth();
	const { apiReq } = useAxios();

	const [isAddLoading, setAddLoading] = useState<boolean>(false);
	const [isDeleteLoading, setDeleteLoading] = useState<boolean>(false);

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setTaskInput(e.target.value);
	};

	const addNewTask = async () => {
		setAddLoading(true);

		const newTask: TaskType = {
			id: "",
			content: taskInput,
		}

		const response = await apiReq<{ taskID: string }, TaskType>(
			"post",
			"/api/tasks/" + obj.id,
			newTask,
			authUser?.authToken,
		);

		if (response) {
			newTask.id = response.data.taskID;

			// handling empty array
			if (tasks) {
				setTasks((prevTasks) => [...prevTasks, newTask]);
			} else {
				setTasks([newTask]);
			}

			// Clear the input after adding the task
			setTaskInput("");
			setClickedAdd(false);
		}

		setAddLoading(false);
	}

	const deleteTask = async (taskID: string) => {
		setDeleteLoading(true);

		const response = await apiReq<unknown, undefined>(
			"delete",
			"/api/tasks/" + taskID,
			undefined,
			authUser?.authToken,
		);

		if (response) {
			// remove from tasks array
			const updatedTasks = tasks.filter(task => task.id !== taskID);
			setTasks(updatedTasks);
		}

		setDeleteLoading(false);
	}

	return (
		<div className="max-w-md p-4 m-4 bg-white border border-gray-200 rounded-lg shadow sm:p-8 dark:bg-gray-800 dark:border-gray-700"
			id={obj.id}>
			<div className="flex items-center justify-between mb-4">
				<h5 className="text-2xl font-bold leading-none text-gray-900 dark:text-white">
					{obj.title}
				</h5>
				<div>
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
					<button
						type="button"
						disabled={deleteLoadingState}
						onClick={() => deleteHandler(obj.id)} // trigger the function passed in as prop
						className="text-white bg-red-700 hover:bg-red-800 font-medium rounded-lg text-sm p-2.5 text-center inline-flex items-center dark:bg-red-600 dark:hover:bg-red-700">
						{deleteLoadingState ? ( // spinner when loading
							<svg
								aria-hidden="true"
								className="inline w-6 h-6 text-gray-200 animate-spin dark:text-gray-600 fill-gray-600 dark:fill-gray-300"
								viewBox="0 0 100 101"
								fill="none"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
									fill="currentColor"
								/>
								<path
									d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
									fill="currentFill"
								/>
							</svg>
						) : ( // trash bin when not loading
							<svg className="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
								<path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5" d="M5 7h14m-9 3v8m4-8v8M10 3h4a1 1 0 0 1 1 1v3H9V4a1 1 0 0 1 1-1ZM6 7h12v13a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7Z" />
							</svg>
						)}

						<span className="sr-only">Delete Tasklist</span>
					</button>
				</div>
			</div>
			<div className="flow-root">
				<ul role="list" className="divide-y divide-gray-200 dark:divide-gray-700">
					{tasks && tasks.map((task) => (
						<li key={task.id} className="py-3 sm:py-4 flex items-center">
							<p className="flex-1 min-w-0 ms-4 text-md font-medium text-gray-900 dark:text-white">
								{task.content}
							</p>
							<button
								className="ml-4"
								disabled={isDeleteLoading}
								onClick={() => deleteTask(task.id)}>
								<svg className="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
									<path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5" d="M5 7h14m-9 3v8m4-8v8M10 3h4a1 1 0 0 1 1 1v3H9V4a1 1 0 0 1 1-1ZM6 7h12v13a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7Z" />
								</svg>
							</button>
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
									disabled={isAddLoading} // disable when api being hit
									className="text-white bg-primary-600 hover:bg-primary-700 font-medium rounded-lg text-sm mt-4 mx-2 px-5 py-2 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
								>
									{isAddLoading ? "Adding..." : "Add"}
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