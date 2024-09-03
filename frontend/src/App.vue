<script lang="ts" setup>
import { AddTodo, GetTodos, RemoveTodo, UpdateTodo } from "@wails/go/main/App";
import { todo } from "@wails/go/models";
import { onMounted, ref } from "vue";
import Todo from "./components/Todo.vue";
import Button from "./components/shadcn/button/Button.vue";

const todos = ref<todo.Todo[]>([]);

async function getTodos() {
  try {
    todos.value = await GetTodos();
  } catch (err) {
    console.error(err);
  }
}

async function addTodo(e: Event) {
  try {
    const form = e.target as HTMLFormElement;
    const formData = new FormData(form);
    const title = formData.get("title")!;
    await AddTodo(title.toString());
    form.reset();
    getTodos();
  } catch (err) {
    console.error(err);
  }
}

async function removeTodo(id: string) {
  try {
    await RemoveTodo(id);
    getTodos();
  } catch (err) {
    console.error(err);
  }
}

async function toggleCheck(id: string, checked: boolean) {
  try {
    await UpdateTodo(new todo.TodoUpdate({ id, completed: checked }));
    getTodos();
  } catch (err) {
    console.error(err);
  }
}

function debounce<A extends any[], T extends (...a: A) => void | Promise<void>>(
  cb: T,
  ms = 300,
) {
  let timeoutId: number | undefined;
  return function (...a: A) {
    if (timeoutId) {
      clearTimeout(timeoutId);
    }

    timeoutId = window.setTimeout(() => cb(...a), ms);
  };
}

async function updateTitle(id: string, title: string) {
  try {
    await UpdateTodo(new todo.TodoUpdate({ id, title }));
    getTodos();
  } catch (err) {
    console.error(err);
  }
}

const debouncedUpdateTitle = debounce(updateTitle);

onMounted(async () => {
  getTodos();
});
</script>

<template>
  <main class="mx-auto mt-12 max-w-3xl">
    <form @submit.prevent="addTodo($event)" class="text-center">
      <input type="text" class="border" name="title" />
      <Button type="submit" class="ml-2">Add Todo</Button>
    </form>
    <ul class="flex flex-col items-center my-4">
      <li v-for="todo in todos" :key="todo.id">
        <Todo
          v-bind="todo"
          @delete="removeTodo"
          @toggle-check="toggleCheck"
          @update-title="debouncedUpdateTitle"
        />
      </li>
    </ul>
  </main>
</template>
