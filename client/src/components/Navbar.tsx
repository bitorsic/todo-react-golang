import React from 'react'
import { Link } from 'react-router-dom'
import { useAuth } from '../hooks/useAuth'
import { useAxios } from '../hooks/useAxios'

const Navbar: React.FC = () => {
	const { authUser, setAuthUser } = useAuth()
	const { apiReq } = useAxios()

	const handleLogout = async () => {
		const response = await apiReq<undefined, undefined>("delete", "/api/logout")

		if (response) {
			localStorage.clear()
			setAuthUser(null)
		}
	}

	return (
		<nav className="bg-white border-gray-200 dark:bg-gray-900">
			<div className="flex flex-wrap items-center justify-between mx-auto p-4">
				<Link
					to="/"
					className="flex items-center space-x-3 rtl:space-x-reverse"
				>
					<span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">
						Task-inator 3000
					</span>
				</Link>
				{!authUser && (<div className="flex mt-4 sm:mt-0 space-x-3 md:space-x-4 rtl:space-x-reverse">
					<Link
						to={"/login"}
						className="text-white bg-blue-700 hover:bg-blue-800 font-medium rounded-lg text-sm px-4 py-2 text-center dark:bg-blue-600 dark:hover:bg-blue-700">
						Log in
					</Link>
					<Link
						to={"/register"}
						className="text-white bg-blue-700 hover:bg-blue-800 font-medium rounded-lg text-sm px-4 py-2 text-center dark:bg-blue-600 dark:hover:bg-blue-700">
						Register
					</Link>
				</div>)}

				{authUser && (<div className="flex mt-4 sm:mt-0 space-x-3 md:space-x-4 rtl:space-x-reverse">
					<p className="text-black dark:text-white px-4 py-2">Hi, {authUser.first_name}</p>
					<button
						onClick={handleLogout}
						className="text-white bg-blue-700 hover:bg-blue-800 font-medium rounded-lg text-sm px-4 py-2 text-center dark:bg-blue-600 dark:hover:bg-blue-700">
						Log out
					</button>
				</div>)}
			</div>
		</nav>

	)
}

export default Navbar