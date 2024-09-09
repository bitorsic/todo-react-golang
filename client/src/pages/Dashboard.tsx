import { useAuth } from "../hooks/useAuth"

const Dashboard = () => {
	const {authUser} = useAuth()

	return (
		<div>Hi, {authUser?.first_name}</div>
	)
}

export default Dashboard