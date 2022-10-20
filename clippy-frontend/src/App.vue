<script setup lang="ts">
import { ItemsOverviewView } from "./views";
</script>

<template>
  <div class="app">
    <ItemsOverviewView></ItemsOverviewView>
  </div>
</template>

<style scoped>
.app {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-content: center;
}

@media (min-width: 1024px) {
  .app {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
