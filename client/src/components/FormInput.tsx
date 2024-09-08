import React from 'react'

interface Props {
	id: string,
	label: string,
	type: string,
	value: string,
	changeHandler: (e: React.ChangeEvent<HTMLInputElement>) => void,
	isRequired?: boolean
}

const FormInput: React.FC<Props> = ({
	id,
	label,
	type,
	value,
	changeHandler,
	isRequired
}) => {
	return (
		<div>
			<label
				htmlFor={id}
				className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
			>
				{label}
			</label>
			<input
				type={type}
				name={id}
				id={id}
				value={value}
				onChange={changeHandler}
				className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
				required={isRequired}
				/>
		</div>
	)
}

export default FormInput