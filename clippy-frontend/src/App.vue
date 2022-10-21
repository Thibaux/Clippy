<script setup lang="ts">
import { ItemsOverviewView, SideBarView } from "./views";
</script>

<template>
    <div class="app">
        <div class="side-bar">
            <SideBarView></SideBarView>
        </div>
        <ItemsOverviewView></ItemsOverviewView>
    </div>
</template>

<style scoped>
.app {
    display: flex;
    flex-direction: row;
    justify-content: space-evenly;
    align-content: center;
    width: 75em;
    height: 100%;
}

.side-bar {
    display: flex;
    margin-top: -10em;
}

@media (max-width: 1224px) {
    .app {
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: center;
        width: 100%;
    }

    .side-bar {
        margin: 0;
        padding-bottom: 4em;
    }
}
</style>
