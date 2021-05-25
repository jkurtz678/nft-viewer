<template>
  <Card class="controller-card">
    <template #header>
      <div class="p-d-flex p-ai-center p-px-3">
        <h2>Displays</h2>
        <div style="flex-grow: 1"></div>
        <Button
          label="New display"
          icon="pi pi-plus"
          @click="openDialog"
          :disabled="!account_id"
          style="height: 0%;"
        />
      </div>
    </template>
    <template #content>
      <div v-if="!displays_loaded">Connect an account to load your displays</div>
      <div v-if="displays_loaded" class="p-d-flex p-flex-wrap">
        <DisplayItem
          v-for="display in displays"
          :key="display.id"
          :name="display.entity.name"
          :code="display.entity.code"
        ></DisplayItem>
      </div>
      <div v-if="displays_loaded && displays.length == 0">No displays found</div>
      <EditDisplayDialog
        v-if="account_id"
        v-model="show_dialog"
        :createDisplay="createDisplay"
        :account_id="account_id"
      ></EditDisplayDialog>
    </template>
  </Card>
</template>

<script lang="ts">
import { defineComponent, watch, ref } from "vue";
import displayManagement from "@/composables/displayManagement";
import EditDisplayDialog from "@/components/Controller/EditDisplayDialog.vue";
import DisplayItem from "@/components/Controller/DisplayItem.vue";

export default defineComponent({
  props: {
    account_id: String,
  },
  components: { EditDisplayDialog, DisplayItem },
  /* if (account.value) {
        await loadDisplays(account.value.id);
      } else {
        alert("Tried to load displays but had no account");
      } */

  setup(props) {
    const show_dialog = ref(false);

    const {
      displays,
      displays_loaded,
      loadDisplays,
      createDisplay,
    } = displayManagement();

    watch(() => props.account_id, async (account_id) => {
      if (account_id) {
        await loadDisplays(account_id);
      } else {
        alert("Tried to load displays but had no account");
      }
    });

    const openDialog = () => {
      show_dialog.value = true;
    };
    return {
      createDisplay,
      displays,
      displays_loaded,
      openDialog,
      show_dialog,
    };
  },
});
</script>

<style>
</style>