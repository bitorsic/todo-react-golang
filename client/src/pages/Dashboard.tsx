import { useEffect, useState } from "react"
import Section from "../components/Section"
import TaskListCard, { TaskListType } from "../components/TaskListCard"
import { useAuth } from "../hooks/useAuth"
import apiClient from "../config/axiosConfig"

const Dashboard: React.FC = () => {
	const [taskLists, setTaskLists] = useState<TaskListType[]>([])
	const { authUser } = useAuth()

	useEffect(() => {
		(async () => {
			try {
				const response = await apiClient.get("/api/tasks")
	
				setTaskLists(response.data.task_lists)
			} catch (error: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
				alert(error.response.data.message)
			}
		})();
	}, [])

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
						<TaskListCard key={taskList.id} obj={taskList}/>
					))}

				</div>
			</div>
		</Section>
	)
}

export default Dashboard