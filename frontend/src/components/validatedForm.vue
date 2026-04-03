<script lang="ts" setup>
import { Field, Form, useFieldArray } from "vee-validate";
import { AddTask } from "../../wailsjs/go/services/TaskService";
import { UpdateTask } from "../../wailsjs/go/services/TaskService";
import { computed } from "vue";
import { object, string } from "yup";
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
  "task-added": [];
  "task-action": [];
}>();

const schema = object({
  name: string()
    .required("Task name is required")
    .min(4, "Task name must be at least 4 characters")
    .max(50, "Task name must be at most 50 characters"),
  desc: string()
    .required("Description is required")
    .min(4, "Task description must be at least 4 characters")
    .max(200, "Description must be at most 200 characters"),
});

async function onSubmit(values: any) {
  if (isEditing.value) {
    try {
      const result = await UpdateTask({
        ...values,
        id: props.task?.id || 0,
        status: props.task?.status || 0,
      });

      emit("task-updated", {
        id: props.task?.id || 0,
        status: props.task?.status || 0,
        name: result.name,
        desc: result.desc,
      });
    } catch (error) {
      console.error("Error updating task:", error);
    }
  } else {
    try {
      await AddTask({
        name: values.name,
        desc: values.desc,
      });
      emit("task-added");
    } catch (error) {
      console.error("Error adding task:", error);
    }
  }
  emit("task-action");
}

const isEditing = computed(() => !!(props.task && props.task.id !== 0));
</script>

<template>
  <Form @submit="onSubmit" :validation-schema="schema" :initial-values="props.task" v-slot="{ errors }">
    <div class="form-element">
      <label for="name" class="form-label">Task</label>
      <Field type="text" class="form-input" id="name" name="name" />
      <p v-if="errors.name" class="error-message">{{ errors.name }}</p>
    </div>
    <div class="form-element">
      <label for="desc" class="form-label">Description</label>
      <Field type="text" class="form-input" id="desc" name="desc" />
      <p v-if="errors.desc" class="error-message">{{ errors.desc }}</p>
    </div>
    <div class="form-group-button">
      <button class="form-button" type="submit">
        {{ isEditing ? "Edit" : "Add" }}
      </button>
      <button class="form-button" type="button" @click="$emit('task-action')">
        Cancel
      </button>
    </div>
  </Form>
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

.error-message {
  display: block;
  padding: 0.5rem;
  border: 2px solid rgb(255, 127, 80);
  border-radius: 5px;
  background-color: rgba(255, 127, 80, 0.301);
  margin: 2px;
  color: coral;
}
</style>
