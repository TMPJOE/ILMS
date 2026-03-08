<script lang="ts" setup>
import ListOfTasks from "./components/ListOfTasks.vue";
import NewTask from "./components/TaskForm.vue";
import { ref } from "vue";

const showForm = ref(false);
</script>

<template>
	<h1>ILMS - Task Manager</h1>
	<NewTask v-if="showForm" />
	<button @click="showForm = !showForm">
		{{ showForm ? "Cancel" : "Add New Task" }}
	</button>
	<ListOfTasks />
</template>

<style lang="css">
h1 {
	text-align: center;
}
button {
	display: block;
	margin: 10px auto;
	padding: 10px 20px;
	font-size: 16px;
	background-color: #008e7f;
	color: white;
	border: none;
	border-radius: 5px;
	cursor: pointer;
}
button:hover {
	background-color: white;
	color: #008e7f;
}
</style>
