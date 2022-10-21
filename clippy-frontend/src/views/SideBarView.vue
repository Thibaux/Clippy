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
    width: 30em;
}

.sidebar-title {
    display: flex;
    justify-content: center;
    align-items: center;
    padding-bottom: 2em;
}

h1 {
    text-align: center;
    font-weight: 500;
    font-size: 2.6rem;
    top: -10px;
}
</style>
