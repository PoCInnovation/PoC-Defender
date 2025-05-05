<template>
  <div class="bg-white p-4 rounded-2xl shadow">
    <h2 class="text-xl mb-2">Active Connections</h2>
    <table class="min-w-full text-left">
      <thead>
        <tr><th>IP</th><th>Endpoint</th><th>Duration</th></tr>
      </thead>
      <tbody>
        <tr v-for="conn in connections" :key="conn.id">
          <td>{{ conn.ip }}</td>
          <td>{{ conn.endpoint }}</td>
          <td>{{ conn.duration }}s</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
    import { defineProps } from 'vue'
const props = defineProps<{ connections: any[] }>()
</script>
