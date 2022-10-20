<script setup lang="ts">
import { ItemsOverviewView, SideBarView } from "./views";
</script>

<template>
  <div class="app">
    <SideBarView></SideBarView>
    <ItemsOverviewView></ItemsOverviewView>
  </div>
</template>

<style scoped>
.app {
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
  width: 75em;
}

@media (min-width: 1024px) {
  .app {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
