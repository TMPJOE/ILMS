<script lang="ts" setup>
import { GetTasks } from "../../wailsjs/go/services/TaskService";
import { ref } from "vue";
import { onMounted } from "vue";
import Task from "./Task.vue";

interface Task {
	id: number;
	status: number;
	name: string;
	desc: string;
	date: string;
}

var tasks = ref([] as Task[]);

function fetchTasks() {
	GetTasks().then(result => {
		tasks.value = result;
	});
}

onMounted(() => {
	fetchTasks();
});
</script>
<template>
	<div class="list-of-tasks">
		<div class="list-of-tasks__header">
			<h2>List of Tasks</h2>
		</div>
		<div class="list-of-tasks__content">
			<div v-if="tasks.length === 0">No tasks available.</div>

			<div v-for="task in tasks" :key="task.id" class="task">
				<Task :task="task" />
			</div>
		</div>
	</div>
</template>

<style lang="css" scoped>
.list-of-tasks {
	display: flex;
	flex-direction: column;
	gap: 1rem;
	margin: 30px;
	border-radius: 10px;
	border: 2px solid #008e7f;
}
.list-of-tasks__header {
	padding: 1rem;
	border-bottom: #008e7f solid 2px;
}
.list-of-tasks__content {
	padding: 1rem;
}
</style>
