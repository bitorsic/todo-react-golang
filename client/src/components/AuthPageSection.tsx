import React from 'react'

interface Props {
	children: React.ReactNode
}

const AuthPageSection: React.FC<Props> = ({children}) => {
	return (
		<section className="bg-gray-50 dark:bg-gray-900">
			<div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
				<a
					href="#"
					className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white"
				>
					Task-inator 3000
				</a>
				{children}
			</div>
		</section>
	)
}

export default AuthPageSection