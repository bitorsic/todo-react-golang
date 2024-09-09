import React, { createContext, useState } from "react";

interface AuthUserType {
	email: string,
	first_name: string,
}

interface AuthContextType {
  authUser: AuthUserType | null,
  setAuthUser: (user: AuthUserType | null) => void,
  isLoggedIn: boolean,
  setIsLoggedIn: (loggedIn: boolean) => void,
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined)

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
	const [authUser, setAuthUser] = useState<AuthUserType | null>(null)
	const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false)

	const value = {
		authUser,
		setAuthUser,
		isLoggedIn,
		setIsLoggedIn
	}

	return (
		<AuthContext.Provider value={value}>
			{children}
		</AuthContext.Provider>
	)
}

