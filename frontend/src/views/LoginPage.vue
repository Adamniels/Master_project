<script setup lang="ts">
import { ref } from "vue";
import { getBackendIP } from "../utils";
import { useRouter } from "vue-router";

const backendIP = getBackendIP();

const register = ref(false);
const username = ref("");
const password = ref("");
const errorMessage = ref("");
const router = useRouter();

const registerUser = async () => {};
const login = async () => {
  try {
    const response = await fetch(`${backendIP}/api/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        username: username.value,
        password: password.value,
      }),
    });

    if (!response.ok) {
      const errorText = await response.text();
      errorMessage.value = errorText || "Fel användarnamn eller lösenord";
      return;
    }

    const data = await response.json();
    localStorage.setItem("token", data.token);
    router.push("/homepage");
  } catch (error) {
    errorMessage.value = "Något gick fel med inloggningen";
    console.error(error);
  }
};
</script>

<template>
  <h1>Welcome to Ascendify</h1>

  <div v-if="register">
    <form @submit.prevent="registerUser">
      <input v-model="username" placeholder="Username" />
      <input v-model="password" placeholder="Password" />
      <button type="submit">Registrera</button>
    </form>
  </div>

  <div v-else>
    <form @submit.prevent="login">
      <input v-model="username" placeholder="Username" />
      <input v-model="password" placeholder="Password" />
      <button type="submit">Logga in</button>
    </form>
  </div>

  <!-- Växlingsknappen -->
  <button @click="register = !register" style="margin-top: 1rem">
    {{
      register ? "Redan medlem? Logga in här" : "Ny användare? Registrera dig"
    }}
  </button>

  <p v-if="errorMessage">{{ errorMessage }}</p>
</template>

<style scoped></style>
