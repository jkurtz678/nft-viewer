<template>
  <Dialog
    header="Header"
    :visible="modelValue"
    :modal="true"
    @update:visible="$emit('update:modelValue', false)"
    style="min-width: 500px;"
  >
    <template #header>New Display</template>
    <span
      class="p-float-label"
      style="margin-top: 20px"
    >
      <InputText
        id="name"
        v-model="name"
        type="text"
      />
      <label for="name">Name</label>
    </span>

    <span
      class="p-float-label"
      style="margin-top: 30px"
    >
      <InputText
        id="code"
        v-model="code"
        type="number"
      ></InputText>
      <label for="code">Code</label>
    </span>
    <template #footer>
      <Button
        label="Save"
        @click="handleSave"
      ></Button>
      <Button
        class="p-button-text"
        label="Cancel"
        @click="$emit('update:modelValue', false)"
      ></Button>
    </template>
  </Dialog>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { ref } from "vue";
export default defineComponent({
  props: {
    modelValue: {
      type: Boolean,
      required: true,
    },
    account_id: {
      type: String,
      required: true,
    },
    createDisplay: {
      type: Function,
      required: true,
    },
  },
  setup(props, { emit }) {
    const name = ref("");
    const code = ref("");

    const handleSave = async () => {
      props.createDisplay(props.account_id, name.value, code.value);
      emit("update:modelValue", false);
    };

    return {
      handleSave,
      name,
      code,
    };
  },
});
</script>

<style>
</style>