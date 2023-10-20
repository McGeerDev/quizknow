export type QuizQuestion = {
	category: string;
	correctAnswer: string;
	difficulty: string;
	id: string;
	incorrectAnswers: string[];
	isNiche: boolean;
	question: {
		text: string;
	};
	regions: string[];
	tags: string[];
};
