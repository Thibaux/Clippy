<script setup lang="ts">
import { ItemCard } from "@/components/Item";

const items = [
    { id: 1, title: "first" },
    { id: 2, title: "sec" },
];
</script>

<template>
    <div class="flex flex-col md:pb-4">
        <div v-for="(item, index) in items" :key="index">
            <div class="pb-10">
                <ItemCard></ItemCard>
            </div>
        </div>
    </div>
</template>

<style scoped></style>
