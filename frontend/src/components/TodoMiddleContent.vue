<template>
  <div class="title">
    <h3>Todo</h3>
  </div>

  <div>
    <ul>
      <li
        v-for="todo in uncompleted"
        :key="todo.id"
        :class="{ done: todo.completed }"
      >
        <input
          type="checkbox"
          v-model="todo.completed"
          @change="toggleCompleted(todo)"
        />
        {{ todo.title }}
      </li>
    </ul>
  </div>

  <div class="title">
    <p>Completed</p>
  </div>
  <div>
    <ul>
      <li
        v-for="todo in completed"
        v-bind:key="todo.id"
        v-bind:class="{ done: todo.completed }"
      >
        <input
          type="checkbox"
          v-model="todo.completed"
          @change="toggleCompleted(todo)"
        />
        {{ todo.title }}
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import type { Todo } from "../types/todo";

// TODO koppla denna till backend istället och sköt den där, uppdatera efter varje
const todos = ref<Todo[]>([
  { id: 1, title: "köp mjölk", completed: false },
  { id: 2, title: "Träna", completed: false },
  { id: 3, title: "Plugga vue", completed: true },
]);

// computer ser till att när jag ändra något den är beroende av så ändras även denna?
const completed = computed(() => {
  return todos.value.filter((todo) => todo.completed);
});
const uncompleted = computed(() => {
  return todos.value.filter((todo) => !todo.completed);
});

const toggleCompleted = async (todo: Todo) => {
  // TODO uppdatera backend med patch requests
};
</script>

<style scoped>
.title {
  display: flex;
  justify-content: center;
  text-decoration: underline;
}

ul {
  list-style: none;
}

.done {
  text-decoration: line-through;
  color: gray;
}
</style>
