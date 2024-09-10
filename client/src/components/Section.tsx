import React from 'react'

interface Props {
	children: React.ReactNode
}

const Section: React.FC<Props> = ({ children }) => {
	return (
		<section className="bg-gray-50 dark:bg-gray-900">
			{children}
		</section>
	)
}

export default Section