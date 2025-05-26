<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-100 to-indigo-200">
    <div class="bg-white p-8 rounded-2xl shadow-md w-full max-w-md">
      <h1 class="text-2xl font-bold mb-6 text-center text-indigo-600">Proxy Admin Login</h1>
      <form @submit.prevent="login">
        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700 mb-1">Username</label>
          <input v-model="username" type="text" required class="w-full p-2 border rounded focus:outline-none focus:ring-2 focus:ring-indigo-400" />
        </div>
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-1">Password</label>
          <input v-model="password" type="password" required class="w-full p-2 border rounded focus:outline-none focus:ring-2 focus:ring-indigo-400" />
        </div>
        <button type="submit" class="w-full py-2 px-4 bg-indigo-500 text-white font-semibold rounded hover:bg-indigo-600 transition">
          Log In
        </button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { ref } from 'vue'

const username = ref('')
const password = ref('')
const router = useRouter()

function login() {
  if (username.value === 'admin' && password.value === 'proxy123') {
    router.push('/')
  } else {
    alert('Invalid credentials')
  }
}
</script>
