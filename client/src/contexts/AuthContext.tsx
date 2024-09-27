import React, { createContext, useEffect, useState } from "react";

export interface AuthUserType {
	authToken: string,
	first_name: string,
}

interface AuthContextType {
  authUser: AuthUserType | null,
  setAuthUser: (user: AuthUserType | null) => void,
	isLoading: boolean,
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined)

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
	const [authUser, setAuthUser] = useState<AuthUserType | null>(null);
	const [isLoading, setLoading] = useState<boolean>(true); // will set to false once authUser set

	const value: AuthContextType = {
		authUser,
		setAuthUser,
		isLoading,
	}

	useEffect(() => {
		const localUser = localStorage.getItem('authUser')
		if (localUser) {
			const userObj: AuthUserType = JSON.parse(localUser)
			setAuthUser(userObj)
		}

		setLoading(false);
	}, [])

	return (
		<AuthContext.Provider value={value}>
			{children}
		</AuthContext.Provider>
	)
}

