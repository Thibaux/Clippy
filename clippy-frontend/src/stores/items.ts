import { defineStore } from "pinia";

export const useItemsStore = defineStore("items", () => {
    const items: Items = [];
    const suggestedTags = [
        "tag1",
        "tag2",
        "tag1",
        "tag2",
        "tag1",
        "tag1",
        "tag2",
        "tag1",
        "tag2",
        "tag2",
        "tag1",
        "tag2",
    ];

    const sendNewItem = (newItemState: Item) => {
        console.log(newItemState.content);
    };

    return { items, suggestedTags, sendNewItem };
});
