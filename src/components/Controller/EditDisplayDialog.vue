<template>
  <Dialog
    header="Header"
    :visible="display_id != null"
    :modal="true"
    @update:visible="$emit('update:display_id', null)"
    style="min-width: 500px;"
  >
    <template #header>New Display</template>
    <div class="p-grid">
      <div class="p-col">
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
        <Button
          class="p-button-sm p-mt-2"
          label="Open in browser"
          @click="openDisplayInBrowser"
          style="display: block;"
        ></Button>
        <Button
          class="p-button-danger p-button-sm p-mt-2"
          label="Forget display"
          @click="forgetDisplay"
          style="display: block;"
        ></Button>
      </div>
      <div class="p-col">
        <display-item name="test"></display-item>
      </div>
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
import DisplayItem from "../Controller/DisplayItem.vue";
import { ref, watch } from "vue";
import { getDisplayByDisplayID, updateDisplay } from "../../api/display";

export default defineComponent({
  components: { DisplayItem },
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
    const display = ref<FirestoreDocument<Display>>({
      id: "",
      entity: { name: "" } as Display,
    });

    const handleSave = async () => {
      //props.createDisplay(props.account_id, name.value);
      await updateDisplay(display.value);
      emit("update:display_id", null);
    };
    const forgetDisplay = async () => {
      display.value.entity.account_id = "";
      display.value.entity.media_url = "";
      await updateDisplay(display.value);
      emit("update:display_id", null);
    };
    const openDisplayInBrowser = async () => {
      window.open(`/#/display?display_id=${props.display_id}`);
    };
    watch(
      props,
      () => {
        if (props.display_id) {
          getDisplayByDisplayID(props.display_id).then((ret_display) => {
            display.value = ret_display;
          });
        } else {
          display.value = { id: "", entity: { name: "" } as Display };
        }
      },
      { immediate: true }
    );

    return {
      handleSave,
      forgetDisplay,
      display,
      openDisplayInBrowser,
    };
  },
});
</script>

<style>
</style>