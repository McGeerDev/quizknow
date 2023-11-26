export type Quiz = {
	id: string;
	questions: Question[];
};

export type Question = {
	text: string;
	answers: Answer[];
};

export type Answer = {
	text: string;
	isCorrect: boolean;
};

export type QuizApi = {
	category: string;
	correctAnswer: string;
	difficulty: string;
	id: string;
	question: {
		text: string;
	};
	incorrectAnswers: string[];
	isNiche: boolean;
	regions: string[];
	tags: string[];
};
