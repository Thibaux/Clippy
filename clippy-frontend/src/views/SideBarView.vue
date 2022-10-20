<script setup lang="ts">
import { CreateItem } from "@/components/Item";
</script>

<template>
  <div class="sidebar-wrapper">
    <div class="sidebar-title">
      <h1>Clippy</h1>
    </div>
    <div class="sidebar-create-item">
      <CreateItem></CreateItem>
    </div>
  </div>
</template>

<style scoped>
.sidebar-wrapper {
  display: flex;
  flex-direction: column;
  height: 20em;
  width: 30em;
}

.sidebar-title {
  display: flex;
  justify-content: center;
  align-items: center;
}

h1 {
  text-align: center;
  font-weight: 500;
  font-size: 2.6rem;
  top: -10px;
}

@media (min-width: 1024px) {
  h1 {
    text-align: left;
  }
}
</style>
