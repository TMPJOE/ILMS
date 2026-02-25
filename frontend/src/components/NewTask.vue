<script setup lang="ts">
import { AddTask } from "../../wailsjs/go/main/App";
import { ref } from "vue";

const task = ref("");
function addTask() {
	// Logic to add a new task goes here
	AddTask(task.value).then(result => {
		console.log(result);
	});
}
</script>

<template>
	<div>
		<form>
			<div>
				<label for="task" class="form-label">Task</label>
				<input
					type="text"
					class="form-control"
					id="task"
					v-model="task"
					aria-describedby="taskHelp"
				/>
				<div id="taskHelp" class="form-text">Add a new task to your list.</div>
				<button type="submit" @click="addTask">Add</button>
			</div>
		</form>
	</div>
</template>
<style lang="css" scoped>
.button {
	margin-top: 10px;
}

.input {
	width: 100%;
	padding: 10px;
	border: 1px solid #ccc;
	border-radius: 4px;
	box-sizing: border-box;
}

#taskHelp {
	font-size: 0.8em;
	color: #6c757d;
}
</style>
