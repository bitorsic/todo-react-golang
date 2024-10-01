import React, { useState } from "react"
import Box from "../../components/Box";
import AuthForm from "../../components/AuthForm";
import FormInput from "../../components/FormInput";
import { Link } from "react-router-dom";
import { useAuth } from "../../hooks/useAuth";
import { AuthUserType } from "../../contexts/AuthContext";
import Section from "../../components/Section";
import { useAxios } from "../../hooks/useAxios";
import PasswordResetModal from "../../components/PasswordResetModal";

interface FormData {
	email: string,
	password: string,
}

const Login: React.FC = () => {
	const [formData, setFormData] = useState<FormData>({
		email: "",
		password: "",
	});
	const { setAuthUser } = useAuth();
	const { apiReq } = useAxios();

	const [isLoading, setLoading] = useState<boolean>(false);
	const [showModal, setShowModal] = useState<boolean>(false);

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = e.target;
		setFormData({
			...formData,
			[name]: value,
		});
	};

	const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault();

		setLoading(true);

		const response = await apiReq<AuthUserType, FormData>("post", "/api/login", formData)

		if (response) {
			const obj = response.data;

			setAuthUser(obj);
			localStorage.setItem("authUser", JSON.stringify(obj));
		}

		setLoading(false);
	};

	return (
		<Section>
			<Box title="Login">
				<AuthForm
					submitHandler={handleSubmit}
					isLoading={isLoading}
					buttonText="Log in">
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
				<div className="grid grid-flow-col">
					<p className="text-sm font-light text-gray-500 dark:text-gray-400 pr-2">
						Don't have an account?{" "}
						<Link to={"/register"}>
							<button className="font-medium text-primary-600 hover:underline dark:text-primary-500">
								Register here
							</button>
						</Link>
					</p>
					<button 
					type="button"
					onClick={() => setShowModal(true)}
					className="text-blue-700 hover:text-white border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-3 py-2 text-center me-2 mb-2 dark:border-blue-500 dark:text-blue-500 dark:hover:text-white dark:hover:bg-blue-500 dark:focus:ring-blue-800">
						Forgot Password?
					</button>

					<PasswordResetModal isVisible={showModal} onClose={() => setShowModal(false)} />
				</div>
			</Box>
		</Section>
	)
}

export default Login