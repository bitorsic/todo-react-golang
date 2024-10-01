import React, { useState } from "react"
import AuthForm from "./AuthForm";
import FormInput from "./FormInput";
import Box from "./Box";
import { useAxios } from "../hooks/useAxios";

interface FormData {
	email: string,
	new_password: string,
	otp: string,
}

interface Props {
	isVisible: boolean,
	onClose: () => void,
}

const PasswordResetModal: React.FC<Props> = ({ isVisible, onClose }) => {
	const [formData, setFormData] = useState<FormData>({
		email: "",
		new_password: "",
		otp: ""
	});
	const [isLoading, setLoading] = useState<boolean>(false);
	const [isOTPSent, setOTPSent] = useState<boolean>(false);
	const { apiReq } = useAxios();

	const handleClose = (e: React.MouseEvent<HTMLDivElement, MouseEvent>) => {
		if ((e.target as HTMLElement).id === "wrapper") {
			onClose();

			// could have setOTPSent(false), but avoiding it in case user misclicks outside
		}
	};

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

		if (!isOTPSent) { // first request for sending otp,
			const response = await apiReq<unknown, FormData>("post", "/api/reset-password", formData)

			if (response) {
				alert("OTP has been sent to your email");
				setOTPSent(true);
			}
		} else { // then using otp to change password
			const response = await apiReq<unknown, FormData>("put", "/api/reset-password", formData)

			if (response) {
				alert("Password has been successfully reset\nPlease log in again");

				// clear the form
				setFormData({
					email: "",
					otp: "",
					new_password: "",
				})

				// close modal
				onClose();
			}
		}

		setLoading(false);
	};

	if (!isVisible) return null;

	return (
		<div
			id="wrapper"
			className="fixed inset-0 bg-black bg-opacity-25 backdrop-blur-sm flex justify-center items-center"
			onClick={handleClose}>
			<Box title="Password Reset">
				<AuthForm
					submitHandler={handleSubmit}
					isLoading={isLoading}
					buttonText={isOTPSent ? "Change Password" : "Send OTP"}>
					<FormInput
						id="email"
						label="Your email"
						type="email"
						value={formData.email}
						changeHandler={handleChange}
						isRequired />

					{isOTPSent && (<>
						<FormInput
							id="otp"
							label="OTP"
							type="text"
							value={formData.otp}
							changeHandler={handleChange}
							isRequired />
						<FormInput
							id="new_password"
							label="New Password"
							type="password"
							value={formData.new_password}
							changeHandler={handleChange}
							isRequired />
					</>)}
				</AuthForm>
			</Box>
		</div>
	)
}

export default PasswordResetModal