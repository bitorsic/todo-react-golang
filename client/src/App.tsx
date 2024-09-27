import { BrowserRouter, Route, Routes } from "react-router-dom"
import Register from "./pages/auth/Register"
import Login from "./pages/auth/Login"
import Dashboard from "./pages/Dashboard"
import Navbar from "./components/Navbar"
import { useAuth } from "./hooks/useAuth"
import FullPageSpinner from "./components/FullPageSpinner"

function App() {
	const { authUser, isLoading } = useAuth()

	if (isLoading) {
		return <FullPageSpinner />
	}

	return (
		<BrowserRouter>
			<Navbar />
			<Routes>
				<Route path="/register"
					element={authUser ? <Dashboard /> : <Register />} />
				<Route path="/login"
					element={authUser ? <Dashboard /> : <Login />} />
				<Route path="/"
					element={authUser ? <Dashboard /> : <Login />} />
			</Routes>
		</BrowserRouter>
	)
}

export default App
