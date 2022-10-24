<script setup lang="ts">
import { CreateItem } from "@/components/Item";
</script>

<template>
    <div class="flex flex-col w-full">
        <div class="flex justify-center items-center pb-8">
            <h1 class="text-center font-medium text-5xl">Clippy</h1>
        </div>
        <div class="sidebar-create-item">
            <CreateItem></CreateItem>
        </div>
    </div>
</template>

<style scoped></style>
