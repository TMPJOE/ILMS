<script lang="ts" setup>
import { AddTask } from "../../wailsjs/go/services/TaskService";
import { UpdateTask } from "../../wailsjs/go/services/TaskService";
import { onMounted, ref } from "vue";

const props = defineProps<{
	task?: {
		id: number;
		status: number;
		name: string;
		desc: string;
	};
}>();

const emit = defineEmits<{
	taskadded: [];
	taskupdated: [taskUpdated: Task];
}>();

type Task = {
	name: string;
	desc: string;
};

var task = ref<Task>({
	name: "",
	desc: "",
});

function addTask() {
	// Logic to add a new task goes here
	AddTask(task.value).then(result => {
		console.log(result);
	});
}

function updateTaskStatus() {
	UpdateTask({
		id: props.task?.id || 0,
		name: task.value.name,
		desc: task.value.desc,
		status: props.task?.status || 0,
	})
		.then(result => {
			console.log(result);
			emit("taskupdated", {
				name: result.name,
				desc: task.value.desc,
			});
			emit("taskadded");
		})
		.catch(error => {
			console.error("Error updating task:", error);
		});
}
const isEditing = ref(false);
function toggleEdit() {
	if (props.task?.id === 0) {
		isEditing.value = false;
		return;
	}
	task.value.name = props.task?.name || "";
	task.value.desc = props.task?.desc || "";
	isEditing.value = true;
}

onMounted(() => {
	toggleEdit();
});
</script>

<template>
	<form class="form">
		<div class="form-element">
			<label for="name" class="form-label">Task</label>
			<input
				type="text"
				class="form-input"
				id="name"
				v-model="task.name"
				aria-describedby="taskHelp"
			/>
		</div>
		<div class="form-element">
			<label for="desc" class="form-label">Description</label>
			<input
				type="text"
				class="form-input"
				id="desc"
				v-model="task.desc"
				aria-describedby="taskHelp"
			/>
		</div>
		<div id="taskHelp" class="form-text">Add a new task to your list.</div>
		<button
			class="form-button"
			type="submit"
			@click.prevent="() => (isEditing ? updateTaskStatus() : addTask())"
		>
			{{ isEditing ? "Edit" : "Add" }}
		</button>
	</form>
</template>

<style lang="css" scoped>
.form {
	display: flex;
	justify-content: space-between;
	align-items: flex-start;
	flex-direction: column;
	padding: 1rem;
	margin: 10px;
	border: 2px solid #008e7f;
	position: relative;
	border-radius: 7px;
	text-align: left;
}

.form-element {
	display: flex;
	flex-direction: column;
	gap: 0.5rem;
	margin-top: 1rem;
	width: 100%;
}

.form-input {
	padding: 0.5rem;
	border: 2px solid #008e7f;
	border-radius: 5px;
	background-color: transparent;
	margin: 2px;
	color: white;
}

.form-button {
	padding: 0.5rem 1rem;
	border: none;
	border-radius: 5px;
	background-color: #008e7f;
	color: white;
	cursor: pointer;
	margin-top: 1rem;
}

.form-button:hover {
	background-color: white;
	color: #008e7f;
}
</style>
