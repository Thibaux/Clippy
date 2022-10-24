import { reactive } from "vue";
import { useItemsStore } from "@/stores/items";

export const useCreateItem = () => {
    const { suggestedTags, sendNewItem } = useItemsStore();
    const newItemErrors = reactive({ contentError: true });
    const newItemState = reactive({
        title: "",
        content: "",
        tags: [{ id: "", name: "" }],
        // eslint-disable-next-line no-undef
    }) as Item;

    const validateNewItemValues = (newItemState: Item): boolean => {
        if (newItemState.content === "") {
            newItemErrors.contentError = false;
            return false;
        }
        return newItemState.title !== "";
    };

    const submitNewItem = async (newItemState: Item) => {
        if (validateNewItemValues(newItemState)) return;

        await sendNewItem(newItemState);
    };

    return {
        newItemErrors,
        newItemState,
        suggestedTags,
        submitNewItem,
    };
};
