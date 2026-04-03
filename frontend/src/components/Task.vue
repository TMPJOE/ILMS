<script lang="ts" setup>
import { ref } from "vue";
import { UpdateTask } from "../../wailsjs/go/services/TaskService";
import { DeleteTask } from "../../wailsjs/go/services/TaskService";
import { onClickOutside } from "@vueuse/core";

const props = defineProps<{
  task: Task;
}>();

interface Task {
  id: number;
  status: number;
  name: string;
  desc: string;
  date: string;
}

interface TaskOptions {
  danger?: boolean;
  label: string;
}

const TaskOptions: TaskOptions[] = [
  { label: "Edit", danger: false },
  { label: "Delete", danger: true },
];

interface TaskUpdate {
  id: number;
  status: number;
  name: string;
  desc: string;
}

var TaskStatus: Record<number, string> = {
  0: "To Do",
  1: "In Progress",
  2: "Pending Review",
  3: "Done",
};

var isVisible = ref<boolean>(true);

function updateTaskStatus(
  taskId: number,
  newStatus: number,
  newDesc: string,
  newName: string,
) {
  var taskToUpdate: TaskUpdate = {
    id: taskId,
    status: newStatus,
    desc: newDesc,
    name: newName,
  };
  UpdateTask(taskToUpdate);
}

const menuRef = ref<HTMLDivElement | null>(null);

onClickOutside(
  () => menuRef.value,
  () => {
    openIndex.value = false;
  },
);

const openIndex = ref<boolean>(false);

function toggleMenu() {
  openIndex.value = !openIndex.value;
}

const getStatusColor = (status: number) => {
  const colors: Record<number, string> = {
    0: "#008e7f",
    1: "coral",
    2: "red",
    3: "lightgreen",
  };
  return colors[status] || "#008e7f";
};

const emit = defineEmits<{
  "open-form": [taskSeleceted: TaskUpdate];
  "delete-task": [];
}>();

function handleAction(task: number, action: string) {
  //edit
  if (action === "Edit") {
    emit("open-form", props.task);
  }
  //delete
  if (action === "Delete") {
    DeleteTask(task);
    isVisible.value = false;
    emit("delete-task")
  }
}
</script>

<template>
  <div class="task" v-if="isVisible" :style="{ borderColor: getStatusColor(task.status) }">
    <div class="task-info">
      <h3>{{ task.name }}</h3>
      <p>{{ task.desc }}</p>
      <p>{{ task.date }}</p>
      <select v-model.number="task.status" @change="updateTaskStatus(task.id, task.status, task.desc, task.name)">
        <option v-for="(status, index) in TaskStatus" :key="index" :value="Number(index)">
          {{ status }}
        </option>
      </select>
    </div>

    <div class="dotmenu" ref="menuRef">
      <button class="dotmenu-trigger" @click.stop="toggleMenu()">
        <span class="dot"></span>
        <span class="dot"></span>
        <span class="dot"></span>
      </button>

      <Transition name="menu-pop">
        <div v-if="openIndex" class="dotmenu-panel">
          <button v-for="action in TaskOptions" :key="action.label" class="dotmenu-action"
            :class="{ 'dotmenu-action--danger': action.danger }" @click="handleAction(task.id, action.label)">
            {{ action.label }}
          </button>
        </div>
      </Transition>
    </div>
  </div>
</template>

<style lang="css" scoped>
.task {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 1rem;
  margin: 10px;
  border: 2px solid #008e7f;
  position: relative;
  border-radius: 7px;
  text-align: left;
}

/* --- 3-dot trigger --- */
.dotmenu {
  position: relative;
  margin-left: auto;
}

.dotmenu-trigger {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 3px;
  padding: 6px;
  background: none;
  border: rgba(27, 38, 54, 1) solid 1px;
  border-radius: 10px;
  cursor: pointer;
  transition: background var(--transition-fast);
}

.dotmenu-trigger:hover {
  border: #008e7f solid 1px;
}

.dot {
  width: 4px;
  height: 4px;
  border-radius: 50%;
  background: white;
}

/* --- Panel --- */
.dotmenu-panel {
  position: absolute;
  top: 100%;
  right: 0;
  min-width: 170px;
  background: rgba(27, 38, 54, 1);
  border: 1px solid #008e7f;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  z-index: 50;
}

.dotmenu-action {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
  border-radius: 7px;
  padding: 8px 14px;
  font-size: 0.84rem;
  font-family: inherit;
  color: white;
  background: none;
  border: none;
  cursor: pointer;
  text-align: left;
  transition: background var(--transition-fast);
}

.dotmenu-action:hover {
  background: darkgray;
}

.dotmenu-action--danger {
  color: lightcoral;
}

.dotmenu-action--danger:hover {
  background: lightcoral;
  color: white;
}

.menu-pop-enter-active,
.menu-pop-leave-active {
  transition:
    opacity 120ms ease,
    transform 120ms ease;
}

.menu-pop-enter-from,
.menu-pop-leave-to {
  opacity: 0;
  transform: scale(0.95) translateY(-4px);
}
</style>
