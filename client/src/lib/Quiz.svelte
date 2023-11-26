<script lang="ts">
	import { onMount } from "svelte";
	import type {
		Quiz,
		Question,
		Answer,
		QuizApi,
	} from "../types/types";

	let quiz: Quiz = {
		id: "",
		questions: [],
	};
	let incorrectAnswers = [];
	let index = 0;

	const nextQuestion = async () => {
		index++;
		if (index >= quiz.questions.length) {
			index = 0;
		}
	};

	const newQuiz = async () => {
		const res = await fetch("http://localhost:8080/questions");
		const data = await res.json();

		let answers = [];
		data.map((item: QuizApi) => {
			console.log(item);
			answers.push(<Answer>{
				text: item.correctAnswer,
				isCorrect: true,
			});
			item.incorrectAnswers.map((answer) => {
				answers.push(<Answer>{ text: answer, isCorrect: false });
			});
		});
		answers = answers.sort(() => Math.random() - 0.5);

		quiz = <Quiz>{
			id: "",
			questions: data.map((item: QuizApi) => {
				let answers: Answer[] = [];
				answers.push(<Answer>{
					text: item.correctAnswer,
					isCorrect: true,
				});
				item.incorrectAnswers.map((answer) => {
					answers.push(<Answer>{ text: answer, isCorrect: false });
				});
				answers = answers.sort(() => Math.random() - 0.5);

				return <Question>{
					text: item.question.text,
					answers,
				};
			}),
		};

		index = 0;
	};

	const checkAnswer = (e: any) => {
		const answer = e.target.innerText;
		const correctAnswer = quiz.questions[index].answers.find(
			(answer) => answer.isCorrect
		)?.text;
		if (answer === correctAnswer) {
			alert("Correct!");
			nextQuestion();
		} else {
			alert("Incorrect!");
			incorrectAnswers = [
				...incorrectAnswers,
				{
					question: quiz.questions[index].text,
					givenAnswer: answer,
					correctAnswer,
				},
			];
		}
		console.log("incorrectAnswers", incorrectAnswers);
	};
</script>

<!-- Wait for data to be returned before displaying a div -->
<!-- Create big button to run newQuiz for a new set of 	questions -->

<div class="card">
	<button on:click={newQuiz}> New Quiz </button>
</div>

{#if quiz.questions.length > 0}
	<div class="card">
		<!-- <h1>Category {index + 1}: {quiz[index]?.category}</h1> -->
		<h3>Question: {quiz.questions[index].text}</h3>
		<div class="answers">
			{#each quiz.questions[index].answers as answer}
				<button on:click={checkAnswer}>{answer.text}</button>
			{/each}
		</div>
		<button class="next" on:click={nextQuestion}>
			Next Question
		</button>
	</div>
{/if}

{#if incorrectAnswers.length > 0}
	{#each incorrectAnswers as answers}
		<div class="card">
			<div>{answers.question}</div>
			<div>Givern answer: {answers.givenAnswer}</div>
			<div>Correct answer: {answers.correctAnswer}</div>
		</div>
	{/each}
{/if}

<style>
	.card {
		background-color: lightgray;
		width: 100%;
		height: 100%;
		padding: 1rem;
		margin: 1rem;
		border-radius: 25px;
	}
	.answers {
		display: flex;
		flex-direction: row;
		justify-content: space-evenly;
	}
	.next {
		margin: 1rem;
	}
</style>
