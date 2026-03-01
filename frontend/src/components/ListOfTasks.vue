<script lang="ts" setup>
import { GetTasks } from "../../wailsjs/go/services/TaskService";
import { ref } from "vue";
import { onMounted } from "vue";

type Task = {
	id: number;
	status: number;
	name: string;
	desc: string;
	date: string;
};
var tasks = ref([] as Task[]);

function fetchTasks() {
	GetTasks().then(result => {
		tasks.value = result;
		console.log(result);
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
			<div v-if="tasks.length > 0" v-for="task in tasks" :key="task.id">
				<h3>{{ task.name }}</h3>
				<p>{{ task.desc }}</p>
				<p>{{ task.date }}</p>
			</div>
		</div>
	</div>
</template>

<style lang="css" scoped>
.list-of-tasks {
	display: flex;
	flex-direction: column;
	gap: 1rem;
}
.list-of-tasks__header {
	padding: 1rem;
	border-radius: 4px;
}
.list-of-tasks__content {
	background-color: #2b2b43;
	padding: 1rem;
	border-radius: 4px;
	box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
</style>
