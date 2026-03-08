<script lang="ts" setup>
import ListOfTasks from "./components/ListOfTasks.vue";
</script>

<template>
	<h1>ILMS - Task Manager</h1>
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
