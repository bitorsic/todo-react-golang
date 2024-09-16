import React, { createContext, useEffect, useState } from "react";

export interface AuthUserType {
	authToken: string,
	first_name: string,
}

interface AuthContextType {
  authUser: AuthUserType | null,
  setAuthUser: (user: AuthUserType | null) => void,
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined)

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
	const [authUser, setAuthUser] = useState<AuthUserType | null>(null)

	const value = {
		authUser,
		setAuthUser,
	}

	useEffect(() => {
		const localUser = localStorage.getItem('authUser')
		if (localUser) {
			const userObj: AuthUserType = JSON.parse(localUser)
			setAuthUser(userObj)
		}
	}, [])

	return (
		<AuthContext.Provider value={value}>
			{children}
		</AuthContext.Provider>
	)
}

