<script setup lang="ts">
import { reactive } from "vue";
import { useCreateItem } from "@/components/Item/CreateItem/useCreateItem";

const createItemState = reactive({ contentFieldHasNoError: true });
const { newItemState, suggestedTags } = useCreateItem();
</script>

<template>
    <div class="create-item-wrapper">
        <div
            class="md:w-full p-4 w-full max-w-sm bg-white rounded-lg border border-gray-200 shadow-md dark:bg-gray-800 dark:border-gray-700"
        >
            <form class="space-y-6" action="#">
                <h5 class="text-xl font-medium text-gray-900 dark:text-white">
                    New item
                </h5>
                <div>
                    <label
                        for="title"
                        class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300 px-2"
                        >Title</label
                    >
                    <input
                        v-model="newItemState.title"
                        class="text-sm rounded-md w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                        placeholder="useful react resource..."
                    />
                </div>
                <div>
                    <label
                        for="content"
                        class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300 px-2"
                        >Content</label
                    >
                    <input
                        v-if="createItemState.contentFieldHasNoError"
                        class="text-sm rounded-md w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                        v-model="newItemState.content"
                        placeholder="https://beta.reactjs.org/learn"
                    />
                    <div v-else>
                        <input
                            type="text"
                            id="username-error"
                            class="bg-red-50 border border-red-500 text-red-900 placeholder-red-700 text-sm rounded-lg focus:ring-red-500 focus:border-red-500 block w-full p-2.5 dark:bg-red-100 dark:border-red-400"
                            placeholder=""
                        />
                        <p class="mt-2 text-sm text-red-600 dark:text-red-500">
                            <span class="font-medium">Oops!</span> Please fill
                            in the content field taken!
                        </p>
                    </div>
                </div>
                <div>
                    <label
                        for="tags"
                        class="mb-2 text-sm font-medium text-gray-900 dark:text-gray-300 px-2"
                        >Tags</label
                    >
                    <p
                        id="floating-helper-text"
                        class="flex justify-end pr-2.5 text-xs text-gray-500 dark:text-gray-400"
                    >
                        Hit enter to add tag
                    </p>
                    <input
                        class="mb-4 text-sm rounded-md w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                        name="content"
                        id="content"
                        placeholder="react"
                    />
                    <label
                        for="remember"
                        class="mt-10 text-sm font-medium text-gray-900 dark:text-gray-300 px-2"
                        >Suggested tags</label
                    >
                    <div class="flex flex-wrap m:justify-center">
                        <div v-for="(tag, index) in suggestedTags" :key="index">
                            <div class="pr-4">
                                {{ tag }}
                            </div>
                        </div>
                    </div>
                </div>
                <div class="flex justify-end w-full">
                    <button
                        type="submit"
                        class="w-auto text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                    >
                        Save item
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>

<style scoped></style>
