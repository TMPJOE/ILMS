<script lang="ts" setup>
import { GetTasks } from "../../wailsjs/go/services/TaskService";
import { ref } from "vue";
import { onMounted } from "vue";
import Task from "./Task.vue";
import ValidatedForm from "./validatedForm.vue";

interface Task {
  id: number;
  status: number;
  name: string;
  desc: string;
  date: string;
}

interface TaskUpdate {
  id: number;
  status: number;
  name: string;
  desc: string;
}

var tasks = ref([] as Task[]);
var indexes = ref([0] as number[]);
var hasMore = ref(false);

async function managePage(way: boolean) {
  if (way == false) {
    indexes.value.pop()
    indexes.value.pop()
    await fetchTasks();
  } else {
    await fetchTasks();
  }
}

async function fetchTasks() {
  try {
    // Get the last ID in the stack, or 0 if empty
    const currentCursor = indexes.value.at(-1) ?? 0;

    const result = await GetTasks(currentCursor);

    if (!result || !result.Tasks) {
      console.error("Invalid response from GetTasks:", result);
      tasks.value = [];
      hasMore.value = false;
      return;
    }

    tasks.value = result.Tasks;
    hasMore.value = result.Tasks.length > 10;

    if (hasMore.value) {
      tasks.value.length = 10;
    }

    if (result.Tasks.length > 0) {
      indexes.value.push(result.LastId);
    }

    console.log("Fetched tasks:", tasks.value.length, "hasMore:", hasMore.value, "lastId:", result.LastId, "indexes length", indexes.value.length);
  } catch (error) {
    console.error("Failed to fetch tasks:", error);
    tasks.value = [];
    hasMore.value = false;
  }
}

const selectedTask = ref<TaskUpdate | undefined>(undefined);
const openForm = ref<boolean>(false);
function openFormFunc(taskToEdit?: TaskUpdate) {
  selectedTask.value = taskToEdit;
  openForm.value = true;
}
function openFormAdd() {
  openForm.value = !openForm.value;
}

function closeForm() {
  openForm.value = false;
  selectedTask.value = undefined;
}

function taskAction() {
  indexes.value.pop();
  fetchTasks();
}

function taskupdated(updatedTask: TaskUpdate) {
  // Update the task in the tasks array
  var index = -1;
  if (updatedTask !== undefined) {
    index = tasks.value.findIndex(task => task.id === updatedTask.id);
  }
  if (index !== -1) {
    tasks.value[index].name = updatedTask.name;
    tasks.value[index].desc = updatedTask.desc;
  }
}

onMounted(() => {
  fetchTasks();
});
</script>
<template>
  <button @click="openFormAdd()">Add New Task</button>
  <div class="list-of-tasks">
    <div class="list-of-tasks__header">
      <h2>List of Tasks</h2>
    </div>
    <div class="list-of-tasks__content">
      <div v-if="tasks.length === 0">No tasks available.</div>

      <div v-for="task in tasks" :key="task.id" class="task">
        <Task :task="task" @open-form="openFormFunc" @delete-task="taskAction()" />
      </div>

      <div v-if="openForm" class="modal-backdrop">
        <ValidatedForm :task="selectedTask" @task-added="taskAction()" @task-updated="taskupdated"
          @task-action="closeForm" class="form" />
      </div>
    </div>
  </div>
  <div class="nav-group-button">
    <button @click="managePage(false)" v-if="indexes.length > 1">Previous</button>
    <button @click="managePage(true)" v-if="hasMore">Next</button>
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

.nav-group-button {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 1rem;

}

.modal-backdrop {
  /* 1. Cover the entire screen */
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;

  /* 2. Create the blur effect */
  background-color: rgba(0, 0, 0, 0.5);
  /* Semi-transparent dark overlay */
  backdrop-filter: blur(8px);
  /* The magic blur line */

  /* 3. Center the Form */
  display: flex;
  justify-content: center;
  align-items: center;

  /* 4. Ensure it's on top of everything else */
  z-index: 999;
}

.form {
  background: rgba(27, 38, 54, 1);
  padding: 2rem;
  border-radius: 12px;
  border: 1px solid #008e7f;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.5);
  width: 90%;
  max-width: 500px;
}
</style>
