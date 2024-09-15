import React, { useState } from "react"
import apiClient from "../../config/axiosConfig";
import AuthPageSection from "../../components/AuthPageSection";
import FormInput from "../../components/FormInput";
import AuthForm from "../../components/AuthForm";
import Box from "../../components/Box";
import { Link, useNavigate } from "react-router-dom";

interface FormData {
	email: string,
	first_name: string,
	last_name: string,
	password: string,
}

const Register: React.FC = () => {
	const [formData, setFormData] = useState<FormData>({
		email: "",
		first_name: "",
		last_name: "",
		password: "",
	});
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
			const response = await apiClient.post("/api/register", formData)
			alert(response.data.message)
			
			navigate("/login")
		} catch (error: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
			alert(error.response.data.message)
		}
	};

	return (
		<AuthPageSection>
			<Box title="Create an account">
				<AuthForm submitHandler={handleSubmit} buttonText="Register">
					<FormInput
						id="email"
						label="Your email"
						type="email"
						value={formData.email}
						changeHandler={handleChange}
						isRequired />
					<FormInput
						id="first_name"
						label="First name"
						type="text"
						value={formData.first_name}
						changeHandler={handleChange}
						isRequired />
					<FormInput
						id="last_name"
						label="Last name"
						type="text"
						value={formData.last_name}
						changeHandler={handleChange} />
					<FormInput
						id="password"
						label="Password"
						type="password"
						value={formData.password}
						changeHandler={handleChange}
						isRequired />
				</AuthForm>
				<p className="text-sm font-light text-gray-500 dark:text-gray-400">
					Already have an account?{" "}
					<Link to={"/login"}>
						<button className="font-medium text-primary-600 hover:underline dark:text-primary-500">
							Login here
						</button>
					</Link>
				</p>
			</Box>
		</AuthPageSection>
	)
}

export default Register