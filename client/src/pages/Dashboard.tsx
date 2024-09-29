import { useEffect, useState } from "react"
import Section from "../components/Section"
import TaskListCard, { TaskListType } from "../components/TaskListCard"
import { useAuth } from "../hooks/useAuth"
import { useAxios } from "../hooks/useAxios"
import FullPageSpinner from "../components/FullPageSpinner"

const Dashboard: React.FC = () => {
	const [taskLists, setTaskLists] = useState<TaskListType[]>([]);
	const [taskListInput, setTaskListInput] = useState<string>("");
	const { authUser } = useAuth()
	const { apiReq } = useAxios()

	// for the entire page
	const [isPageLoading, setPageLoading] = useState<boolean>(true);
	// for new TaskList
	const [isTaskListLoading, setTaskListLoading] = useState<boolean>(false);
	// for deleting TaskList
	const [isDeleteLoading, setDeleteLoading] = useState<boolean>(false);

	useEffect(() => {
		(async () => {
			if (authUser) { // so only execute this code if authUser has been assigned
				const response = await apiReq<TaskListType[], undefined>(
					"get",
					"/api/tasks",
					undefined,
					authUser?.authToken,
				)

				if (response) {
					// handling null
					if (response.data) {
						setTaskLists(response.data);
					} else {
						setTaskLists([])
					}
				}
			}

			setPageLoading(false);
		})();
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [authUser])

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { value } = e.target;
		setTaskListInput(value);
	};

	const addNewTaskList = async () => {
		setTaskListLoading(true);

		const newTaskList: TaskListType = {
			id: "",
			title: taskListInput,
			tasks: [],
		}

		const response = await apiReq<{ taskListID: string }, TaskListType>(
			"post",
			"/api/tasks",
			newTaskList,
			authUser?.authToken
		)
		if (response) {
			newTaskList.id = response.data.taskListID
			setTaskLists((prevTaskLists) => [...prevTaskLists, newTaskList]);

			// Clear the input after adding the tasklist
			setTaskListInput("");
		}

		setTaskListLoading(false);
	}

	const deleteTaskList = async (taskListID: string) => {
		setDeleteLoading(true);

		const response = await apiReq<unknown, undefined>(
			"delete",
			"/api/tasks/list/" + taskListID,
			undefined,
			authUser?.authToken,
		);

		if (response) {
			// remove from tasks array
			const updatedTaskLists = taskLists.filter(taskList => taskList.id !== taskListID);
			setTaskLists(updatedTaskLists);
		}

		setDeleteLoading(false);
	}

	if (isPageLoading) {
		return <FullPageSpinner />
	}

	return (
		<Section>
			<div className="grid w-11/12 cols md:grid-cols-2 lg:grid-cols-3">
				{taskLists.map((taskList) => (
					<TaskListCard
						key={taskList.id}
						obj={taskList}
						deleteHandler={deleteTaskList}
						deleteLoadingState={isDeleteLoading} />
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
							{isTaskListLoading ? ( // spinner when loading
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
							) : ( // plus icon when not loading
								<svg className="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
									<path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M5 12h14m-7 7V5" />
								</svg>
							)
							}
							<span className="sr-only">Add TaskList</span>
						</button>
					</div>
				</div>
			</div>
		</Section>
	)
}

export default Dashboard