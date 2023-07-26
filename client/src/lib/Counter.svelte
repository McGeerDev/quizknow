<script lang="ts">
	let count: number = 0;
	let questions = [];
	const increment = async () => {
		try {
			const res = await fetch("http://localhost:8080/questions");
			const data = await res.json();

			questions = data.map((item) => ({
				category: item.category,
				text: item.question.text,
			}));
			console.log("questions", questions);
			console.log("data", data);
		} catch (e) {
			console.error(e);
		}
	};
</script>

<button on:click={increment}> get questions </button>

<ul>
	{#each questions as question}
		<div>{question.text}</div>
	{/each}
</ul>
