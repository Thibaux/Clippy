<script setup lang="ts">
import { ItemCard } from "@/components/Item";
</script>

<template>
  <div class="items-overview-wrapper">
    <ItemCard></ItemCard>
  </div>
</template>

<style scoped>
.items-overview-wrapper {
  display: flex;
  flex-direction: row;
}
</style>
