import { BrowserRouter, Route, Routes } from "react-router-dom"
import Register from "./pages/auth/Register"
import Login from "./pages/auth/Login"
import ProtectedRoute from "./components/ProtectedRoute"
import Dashboard from "./pages/Dashboard"

function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/register" element={<Register />} />
				<Route path="/login" element={<Login />} />
				<Route path="/" element={
					<ProtectedRoute><Dashboard /></ProtectedRoute>
				} />
			</Routes>
		</BrowserRouter>
	)
}

export default App
