<template>
  <div class="p-6 space-y-6">
    <h1 class="text-2xl font-bold">Proxy Admin Dashboard</h1>
    <StatsChart :data="statsData" />
    <AttackAlerts :alerts="alerts" />
    <ConnectionList :connections="connections" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useProxyApi } from '~/composables/useProxyApi'
import StatsChart from '~/components/StatsChart.vue'
import AttackAlerts from '~/components/AttackAlerts.vue'
import ConnectionList from '~/components/ConnectionList.vue'

const { getConnections, getAlerts, getStats } = useProxyApi()
const connections = ref([])
const alerts = ref([])
const statsData = ref({})

onMounted(async () => {
  connections.value = await getConnections()
  alerts.value = await getAlerts()
  statsData.value = await getStats()
})
</script>
