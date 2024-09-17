import React, { useState } from "react"
import Box from "../../components/Box";
import AuthForm from "../../components/AuthForm";
import FormInput from "../../components/FormInput";
import { Link } from "react-router-dom";
import { useAuth } from "../../hooks/useAuth";
import { AuthUserType } from "../../contexts/AuthContext";
import Section from "../../components/Section";
import { useAxios } from "../../hooks/useAxios";

interface FormData {
	email: string,
	password: string,
	device_id: string,
}

const Login: React.FC = () => {
	const [formData, setFormData] = useState<FormData>({
		email: "",
		password: "",
		device_id: "",
	});
	const { setAuthUser } = useAuth()
	const { apiReq } = useAxios()

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = e.target;
		setFormData({
			...formData,
			[name]: value,
		});
	};

	const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault();

		const response = await apiReq<AuthUserType, FormData>("post", "/api/login", formData)

		if (response) {
			const obj = response.data

			setAuthUser(obj)
			localStorage.setItem("authUser", JSON.stringify(obj))
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