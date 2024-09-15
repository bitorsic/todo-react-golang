import { useEffect, useState } from "react"
import Section from "../components/Section"
import TaskListCard, { TaskListType } from "../components/TaskListCard"
import { useAuth } from "../hooks/useAuth"
import apiClient from "../config/axiosConfig"

const Dashboard: React.FC = () => {
	const [taskLists, setTaskLists] = useState<TaskListType[]>([])
	const [taskListInput, setTaskListInput] = useState<string>("")
	const { authUser } = useAuth()

	useEffect(() => {
		(async () => {
			try {
				const response = await apiClient.get("/api/tasks", {
					headers: { 'Authorization': `Bearer ${authUser?.authToken}` }
				})

				setTaskLists(response.data.task_lists)
			} catch (error: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
				alert(error.response.data.message)
			}
		})();
	}, []) // eslint-disable-line react-hooks/exhaustive-deps

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { value } = e.target;
		setTaskListInput(value);
	};

	const addNewTaskList = async () => {
		let newTaskList: TaskListType = {
			id: "",
			title: taskListInput,
			tasks: [],
		}

		try {
			const response = await apiClient.post("/api/tasks", newTaskList, {
				headers: { 'Authorization': `Bearer ${authUser?.authToken}` }
			})

			newTaskList = response.data.task_list
			setTaskLists((prevTaskLists) => [...prevTaskLists, newTaskList]);

			// Clear the input after adding the tasklist
			setTaskListInput("");
		} catch (error: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
			console.log(error)
			alert(error.response.data.message)
		}
	}

	return (
		<Section>
			<div className="flex flex-col items-center justify-center px-6 py-8 mx-auto min-h-screen md:h-full lg:py-0">
				<p
					className="flex items-center mb-6 mt-5 text-2xl font-semibold text-gray-900 dark:text-white"
				>
					Hi, {authUser?.first_name}
				</p>
				<div className="grid w-11/12 cols md:grid-cols-2 lg:grid-cols-3">
					{taskLists.map((taskList) => (
						<TaskListCard key={taskList.id} obj={taskList} />
					))}
					<div
						className="max-w-md p-4 m-4 bg-white border border-gray-200 rounded-lg shadow sm:p-8 dark:bg-gray-800 dark:border-gray-700">
						<div className="flex items-center justify-between mb-4">
							<input
								type="text"
								value={taskListInput}
								onChange={handleChange}
								placeholder="Add New List"
								className="block w-full mr-2 p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 text-base focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" />
							<button
								type="button"
								onClick={addNewTaskList}
								className="text-white bg-blue-700 hover:bg-blue-800 font-medium rounded-lg text-sm p-2.5 text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700">
								<svg className="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
									<path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M5 12h14m-7 7V5" />
								</svg>
								<span className="sr-only">Add TaskList</span>
							</button>
						</div>
					</div>
				</div>
			</div>
		</Section>
	)
}

export default Dashboard