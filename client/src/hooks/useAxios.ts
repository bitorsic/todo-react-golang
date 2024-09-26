import { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";
import apiClient from "../config/axiosConfig";
import { useAuth } from "./useAuth";

export const useAxios = () => {
	const { authUser, setAuthUser } = useAuth()

	const apiReq = async <TResponse, TData>(
		method: AxiosRequestConfig["method"],
		url: string,
		data?: TData, // Data for POST/PUT requests
		authToken?: string,
	): Promise<AxiosResponse<TResponse> | undefined> => {
		try {
			const response = await apiClient<TResponse>({
				method,
				url,
				// Use empty string if authToken is undefined
				headers: { 'Authorization': `Bearer ${authToken || ''}` },
				data,
				withCredentials: true,
			});

			return response;
		} catch (error: unknown) {
			const axiosError = error as AxiosError<{ error: string }>

			if (axiosError.response) {
				// handling refresh token
				if (
					axiosError.response.data.error == "missing auth token" ||
					axiosError.response.data.error == "auth token expired"
				) {
					const fetchedAuthToken = await refreshReq()

					if (fetchedAuthToken && authUser) {
						const obj = {
							...authUser,
							authToken: fetchedAuthToken,
						}

						setAuthUser(obj)
						localStorage.setItem("authUser", JSON.stringify(obj))

						return apiReq(method, url, data, fetchedAuthToken)
					}
				} else {
					alert(axiosError.response.data.error || "An error occurred");
				}

			} else {
				console.error("An unknown error occurred", axiosError);
				alert("An unknown error occurred.");
			}
			return undefined; // Return undefined in case of an error
		}
	}

	const refreshReq = async () => {
		try {
			const { data } = await apiClient<{ authToken: string }>({
				method: "get",
				url: "/api/refresh",
				withCredentials: true,
			})

			return data.authToken
		} catch (error: unknown) {
			const axiosError = error as AxiosError<{ error: string }>

			if (axiosError.response) {
				if (
					axiosError.response.data.error == "missing refresh token" ||
					axiosError.response.data.error == "refresh token expired" ||
					axiosError.response.data.error == "refresh token blacklisted"
				) {
					// logout
					localStorage.clear()
					setAuthUser(null)
					alert("Session expired. Please log in again")
				} else {
					alert(axiosError.response.data.error || "An error occurred");
				}

			} else {
				console.error("An unknown error occurred", axiosError);
				alert("An unknown error occurred.");
			}
			return undefined; // Return undefined in case of an error
		}
	}

	return { apiReq }
}

