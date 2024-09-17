import { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";
import apiClient from "../config/axiosConfig";

export const useAxios = () => {
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
			});

			return response;
		} catch (error: unknown) {
			const axiosError = error as AxiosError<{ error: string }>

			if (axiosError.response) {
				if (
					axiosError.response.data.error == "missing auth token" ||
					axiosError.response.data.error == "auth token expired"
				) {
					// TODO: send refresh request to backend
					alert("refresh required")
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

