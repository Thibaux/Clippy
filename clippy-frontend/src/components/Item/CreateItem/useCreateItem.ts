import { reactive } from "vue";

export const useCreateItem = () => {
    const newItemState = reactive({
        title: "",
        content: "",
        tags: [{ id: "", name: "" }],
        // eslint-disable-next-line no-undef
    }) as NewItemState;

    // const validateField;

    setInterval(() => {
        console.log(newItemState);
    }, 2000);

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

    return {
        newItemState,
        suggestedTags,
    };
};
