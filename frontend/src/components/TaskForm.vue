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
	"task-updated": [
		taskUpdated: { id: number; status: number; name: string; desc: string },
	];
	"task-added": [
		taskAdded: {
			id: number;
			status: number;
			name: string;
			desc: string;
			date: string;
		},
	];
	"task-action": [];
}>();

type Task = {
	name: string;
	desc: string;
};

var task = ref<Task>({
	name: "",
	desc: "",
});

function closeForm() {
	emit("task-action");
}

async function addTask() {
	try {
		const result = await AddTask({
			name: task.value.name,
			desc: task.value.desc,
		});
		emit("task-added", result);
	} catch (error) {
		console.error("Error adding task:", error);
	}
}

function updateTaskStatus() {
	UpdateTask({
		id: props.task?.id || 0,
		name: task.value.name,
		desc: task.value.desc,
		status: props.task?.status || 0,
	})
		.then(result => {
			emit("task-updated", {
				id: props.task?.id || 0,
				status: props.task?.status || 0,
				name: result.name,
				desc: task.value.desc,
			});
			emit("task-action");
		})
		.catch(error => {
			console.error("Error updating task:", error);
		});
}
const isEditing = ref(false);
function toggleEdit() {
	if (props.task === undefined || props.task.id === 0) {
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
		<div class="form-group-button">
			<button
				class="form-button"
				type="submit"
				@click.prevent="() => (isEditing ? updateTaskStatus() : addTask())"
			>
				{{ isEditing ? "Edit" : "Add" }}
			</button>
			<button class="form-button" type="button" @click="$emit('task-action')">
				Cancel
			</button>
		</div>
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

.form-group-button {
	display: flex;
	flex-direction: row;
	align-items: center;
	gap: 1rem;
}
</style>
