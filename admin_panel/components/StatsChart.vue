<template>
  <div class="bg-white p-4 rounded-2xl shadow">
    <h2 class="text-xl mb-2">Traffic Stats</h2>
    <LineChart :chart-data="chartData" />
  </div>
</template>

<script setup lang="ts">
import { computed, defineProps } from 'vue'
import { Line } from 'vue-chartjs'
import { Chart, Title, Tooltip, Legend, LineElement, PointElement, CategoryScale, LinearScale } from 'chart.js'
Chart.register(Title, Tooltip, Legend, LineElement, PointElement, CategoryScale, LinearScale)

const props = defineProps<{ data: any }>()

const chartData = computed(() => ({
  labels: props.data.timestamps || [],
  datasets: [{
    label: 'Requests per Minute',
    data: props.data.requests || [],
    tension: 0.4
  }]
}))
</script>
