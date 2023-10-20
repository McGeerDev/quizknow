<script lang="ts">
	import { onMount } from "svelte";
	import type { QuizQuestion } from "../types/types";
	let questions: QuizQuestion[] = [];
	let index = 0;
	const nextQuestion = async () => {
		index++;
		if (index >= questions.length) {
			index = 0;
		}
	};

	onMount(async () => {
		try {
			const res = await fetch("http://localhost:8080/questions");
			const data = await res.json();

			questions = data.map((item: QuizQuestion) => ({
				category: item.category,
				correctAnswer: item.correctAnswer,
				difficulty: item.difficulty,
				id: item.id,
				incorrectAnswers: item.incorrectAnswers,
				isNiche: item.isNiche,
				question: { text: item.question.text },
				regions: item.regions,
				tags: item.tags,
			}));
			console.log("questions", questions);
		} catch (e) {
			console.error(e);
		}
	});
</script>

<!-- Wait for data to be returned before displaying a div -->
<button on:click={nextQuestion}> Next Question </button>

{#if questions.length > 0}
	<div class="card">
		<h1>Category {index + 1}: {questions[index]?.category}</h1>
		<div class="answers">
			<!-- use a loop to display the answers together -->
			<button>{questions[index]?.correctAnswer}</button>
			{#each questions[index]?.incorrectAnswers as answer}
				<button>{answer}</button>
			{/each}
		</div>
	</div>

	<div class="card">
		<h3>Question: {questions[index]?.question.text}</h3>
		<button>True</button>
		<button>False</button>
	</div>
{/if}

<style>
	.card {
		background-color: lightgray;
		width: 100%;
		height: 100%;
	}
	.answers {
		display: flex;
		flex-direction: row;
		justify-content: space-evenly;
	}
</style>
