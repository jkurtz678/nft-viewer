<template>
  <Dialog
    header="Header"
    :visible="display_id != null"
    :modal="true"
    @update:visible="$emit('update:display_id', null)"
    style="min-width: 500px;"
  >
    <template #header>New Display</template>
    <div class="p-d-flex">
      <span
        class="p-float-label"
        style="margin-top: 20px;"
      >
        <InputText
          id="name"
          v-model="display.entity.name"
          type="text"
          style="margin: 0 auto;"
        />
        <label for="name">Name</label>
      </span>
    </div>
    <template #footer>
      <Button
        label="Save"
        @click="handleSave"
      ></Button>
      <Button
        class="p-button-text"
        label="Cancel"
        @click="$emit('update:display_id', null)"
      ></Button>
    </template>
  </Dialog>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Display, FirestoreDocument } from "../../types/types";
import { ref, watch} from "vue";
import { getDisplayByDisplayID, updateDisplay } from "../../api/display";

export default defineComponent({
  props: {
    createDisplay: {
      type: Function,
      required: true,
    },
    display_id: {
      type: String,
      required: false,
    },
  },
  setup(props, { emit }) {
    const display = ref<FirestoreDocument<Display>>({ id: "", entity: { name: "" } as Display });

    const handleSave = async () => {
      //props.createDisplay(props.account_id, name.value);
      await updateDisplay(display.value)
      emit("update:display_id", null);
    };
    watch(props, () => {
      if (props.display_id) {
        getDisplayByDisplayID(props.display_id).then((ret_display) => {
          display.value = ret_display;
        });
      } else {
        display.value = { id: "", entity: { name: "" } as Display };
      }
    }, {immediate: true});

    return {
      handleSave,
      display,
    };
  },
});
</script>

<style>
</style>