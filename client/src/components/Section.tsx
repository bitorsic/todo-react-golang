import React from 'react'

interface Props {
	children: React.ReactNode
}

const Section: React.FC<Props> = ({ children }) => {
	return (
		<section className="flex flex-col items-center justify-center px-6 py-8 mx-auto min-h-screen lg:py-0 bg-gray-50 dark:bg-gray-900">
			{children}
		</section>
	)
}

export default Section