import React from 'react'
import { useAuth } from '../hooks/useAuth'
import { Navigate } from 'react-router-dom'

interface Props {
	children: React.ReactNode
}

const ProtectedRoute: React.FC<Props> = ({ children }) => {
	const { authUser } = useAuth()

	if (!authUser)
		return <Navigate to="/login" />

	return (<>{ children }</>)
}

export default ProtectedRoute