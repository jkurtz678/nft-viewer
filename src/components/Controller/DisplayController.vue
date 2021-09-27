<template>
  <Card class="controller-card">
    <template #header>
      <div class="p-d-flex p-ai-center p-px-3">
        <h2>Displays</h2>
        <div style="flex-grow: 1"></div>
        <Button
          v-if="scan_qr"
          label="Cancel scanning"
          icon="pi pi-ban"
          :disabled="!account_id"
          style="height: 0%;"
          @click="scan_qr = false"
        />
        <Button
          v-else
          label="Scan display"
          icon="pi pi-camera"
          :disabled="!account_id"
          style="height: 0%;"
          @click="scan_qr = true"
        />
      </div>
    </template>
    <template #content>
      <div v-if="!displays_loaded">Connect an account to load your displays</div>
      <qrcode-stream v-if="scan_qr" @decode="onDecode"></qrcode-stream>
      <div
        v-else-if="displays_loaded"
        class="p-d-flex p-flex-wrap"
      >
        <DisplayItem
          v-for="display in displays"
          :key="display.id"
          :name="display.entity.name"
          :token_id="display.entity.token_id"
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
import { QrcodeStream } from "vue-qrcode-reader";
import {addAccountToDisplay} from "@/api/display"

export default defineComponent({
  props: {
    account_id: String,
  },
  components: { EditDisplayDialog, DisplayItem, QrcodeStream },
  /* if (account.value) {
        await loadDisplays(account.value.id);
      } else {
        alert("Tried to load displays but had no account");
      } */

  setup(props) {
    const edit_display_id = ref<String | null>(null);
    const scan_qr = ref(false);
    
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
      console.log("EDIT DISPLAY", display_id);
      //show_dialog.value = true;
      edit_display_id.value = display_id;
    };

    const onDecode = async (qr_code_link: string) => {
      const display_id_index = qr_code_link.indexOf("display_id=")
      if(display_id_index > -1) {
        const display_id = qr_code_link.slice(display_id_index + 11);
        await addAccountToDisplay(display_id, props.account_id || "")
      }
      scan_qr.value = false;
    }

    return {
      createDisplay,
      displays,
      displays_loaded,
      editDisplay,
      edit_display_id,
      scan_qr,
      onDecode
    };
  },
});
</script>

<style scoped>
.p-card::v-deep .p-card-body {
  padding: 0.8rem;
}
</style>