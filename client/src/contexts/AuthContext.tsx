import React, { createContext, useState } from "react";

interface AuthUserType {
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

	return (
		<AuthContext.Provider value={value}>
			{children}
		</AuthContext.Provider>
	)
}

