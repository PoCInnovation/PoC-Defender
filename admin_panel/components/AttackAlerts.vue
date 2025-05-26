<template>
  <div class="bg-red-50 p-4 rounded-2xl shadow">
    <h2 class="text-xl mb-2 text-red-600">Attack Alerts</h2>
    <ul>
      <li v-for="alert in alerts" :key="alert.id" class="py-2 border-b">
        <strong>{{ alert.type }}</strong> at {{ alert.timestamp }}
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
    import { defineProps } from 'vue'

const props = defineProps<{ alerts: any[] }>()
</script>
