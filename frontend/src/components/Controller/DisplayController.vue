<template>
  <Card class="controller-card">
    <template #header>
      <div class="p-d-flex p-ai-center p-px-3">
        <h2>Displays</h2>
        <div style="flex-grow: 1"></div>
        <Button
          label="Scan display"
          icon="pi pi-camera"
          :disabled="!account_id"
          style="height: 0%;"
          @click="store.commit('setCameraScanMode', true)"
        />
      </div>
    </template>
    <template #content>
      <div v-if="!displays_loaded">Connect an account to load your displays</div>
      <div
        v-else
        class="p-d-flex p-flex-wrap"
      >
        <DisplayItem
          v-for="display in displays"
          :key="display.id"
          :name="display.entity.name"
          :token_id="display.entity.token_id"
          :asset_contract_address="display.entity.asset_contract_address"
          @editDisplay="editDisplay(display.id)"
        ></DisplayItem>
      </div>
      <div v-if="displays_loaded && displays.length == 0">No displays found</div>
      <EditDisplayDialog
        v-if="account_id"
        :createDisplay="createDisplay"
        :account_id="account_id"
        v-model:display_id="edit_display_id"
      ></EditDisplayDialog>
    </template>
  </Card>
</template>

<script lang="ts">
import { defineComponent, watch, ref } from "vue";
import displayManagement from "@/composables/displayManagement";
import EditDisplayDialog from "@/components/Controller/EditDisplayDialog.vue";
import DisplayItem from "@/components/Controller/DisplayItem.vue";
import { useStore } from "vuex";

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
    const store = useStore();
    const edit_display_id = ref<String | null>(null);
    
    const { displays, displays_loaded, loadDisplays, createDisplay } =
      displayManagement();

    watch(
      () => props.account_id,
      async (account_id) => {
        if (account_id) {
          await loadDisplays(account_id);
        } else {
          alert("Tried to load displays but had no account");
        }
      }
    );

    const editDisplay = (display_id: string) => {
      edit_display_id.value = display_id;
    };

    

    return {
      createDisplay,
      displays,
      displays_loaded,
      editDisplay,
      edit_display_id,
      store
    };
  },
});
</script>

<style scoped>
.p-card :deep(.p-card-body){
  padding: 0.8rem;
}
</style>