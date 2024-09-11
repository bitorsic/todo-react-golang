import { TaskListType } from "../pages/Dashboard"

interface Props {
	obj: TaskListType,
}

const TaskListCard: React.FC<Props> = ({ obj }) => {
	return (
		<div className="max-w-md p-4 m-4 bg-white border border-gray-200 rounded-lg shadow sm:p-8 dark:bg-gray-800 dark:border-gray-700"
			id={obj.id}>
			<div className="flex items-center justify-between mb-4">
				<h5 className="text-2xl font-bold leading-none text-gray-900 dark:text-white">
					{obj.title}
				</h5>
			</div>
			<div className="flow-root">
				<ul role="list" className="divide-y divide-gray-200 dark:divide-gray-700">
					{obj.tasks && obj.tasks.map((task) => (
						<li className="py-3 sm:py-4 flex items-center">
							<p className="flex-1 min-w-0 ms-4 text-md font-medium text-gray-900 dark:text-white">
								{task}
							</p>
						</li>
					))}
				</ul>
			</div>
		</div>
	)
}

export default TaskListCard