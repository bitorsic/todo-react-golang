import React, { useState } from "react"
import apiClient from "../../config/axiosConfig";
import AuthPageSection from "../../components/AuthPageSection";
import Box from "../../components/Box";
import AuthForm from "../../components/AuthForm";
import FormInput from "../../components/FormInput";
import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "../../hooks/useAuth";

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
	const navigate = useNavigate()

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

			setAuthUser({
				authToken: response.data.authToken,
				first_name: response.data.first_name,
			})

			alert(response.data.message)
			navigate("/")
		} catch (error: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
			console.log(error)
			alert(error.response.data.message)
		}
	};

	return (
		<AuthPageSection>
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
		</AuthPageSection>
	)
}

export default Login