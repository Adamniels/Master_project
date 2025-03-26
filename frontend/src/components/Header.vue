<template>
  <header class="header">
    <div class="left">
      <h1 class="title">Ascendify</h1>
    </div>
    <div class="right">
      <span class="username">{{ username }}</span>
      <button @click="logout">Logga ut</button>
    </div>
  </header>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref } from "vue";

const username = ref(localStorage.getItem("username"));

const router = useRouter();

const logout = () => {
  localStorage.removeItem("token");
  router.push("/login");
};
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #1e1e2f;
  color: white;
  padding: 1rem 2rem;
  margin: 0;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
  font-family: "Segoe UI", sans-serif;
}

.title {
  margin: 0;
  font-size: 1.5rem;
}

.right {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.username {
  font-weight: 500;
  opacity: 0.9;
}

button {
  background-color: #ff4b5c;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #e63e4c;
}
</style>
