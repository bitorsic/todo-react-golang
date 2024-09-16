import React, { useState } from "react"
import apiClient from "../../config/axiosConfig";
import Box from "../../components/Box";
import AuthForm from "../../components/AuthForm";
import FormInput from "../../components/FormInput";
import { Link } from "react-router-dom";
import { useAuth } from "../../hooks/useAuth";
import { AuthUserType } from "../../contexts/AuthContext";
import Section from "../../components/Section";

interface FormData {
	email: string,
	password: string,
}

const Login: React.FC = () => {
	const [formData, setFormData] = useState<FormData>({
		email: "",
		password: "",
	});
	const { setAuthUser } = useAuth()

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = e.target;
		setFormData({
			...formData,
			[name]: value,
		});
	};

	const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault();

		try {
			const response = await apiClient.post("/api/login", formData)
			const obj: AuthUserType = {
				authToken: response.data.authToken,
				first_name: response.data.first_name,
			}

			setAuthUser(obj)
			localStorage.setItem("authUser", JSON.stringify(obj))
		} catch (error: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
			alert(error.response.data.error)
		}
	};

	return (
		<Section>
			<Box title="Login">
				<AuthForm submitHandler={handleSubmit} buttonText="Log in">
					<FormInput
						id="email"
						label="Your email"
						type="email"
						value={formData.email}
						changeHandler={handleChange}
						isRequired />
					<FormInput
						id="password"
						label="Password"
						type="password"
						value={formData.password}
						changeHandler={handleChange}
						isRequired />
				</AuthForm>
				<p className="text-sm font-light text-gray-500 dark:text-gray-400">
					Don't have an account?{" "}
					<Link to={"/register"}>
						<button className="font-medium text-primary-600 hover:underline dark:text-primary-500">
							Register here
						</button>
					</Link>
				</p>
			</Box>
		</Section>
	)
}

export default Login