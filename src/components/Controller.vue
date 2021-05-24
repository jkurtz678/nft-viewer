<template>
  <div>
    <Card class="controller-card">
      <template #title>Account</template>
      <template #content>
        <div :style="{color: web3_modal ? '#4CAF50' : '#F44336'}">Web3</div>
        <div :style="{color: signature ? '#4CAF50' : '#F44336'}">Metamask</div>
        <div :style="{color: account ? '#4CAF50' : '#F44336'}">Firebase</div>
      </template>
      <template #footer>
        <Button
          label="Connect"
          icon="pi pi-check"
          @click="connectAccount"
        />
      </template>
    </Card>
    <Card class="controller-card">
      <template #header>
        <div class="p-d-flex p-ai-center p-px-3">
          <h2>Displays</h2>
          <div style="flex-grow: 1"></div>
          <Button
            label="New display"
            icon="pi pi-plus"
            @click="openDialog"
            :disabled="!account"
            style="height: 0%;"
          />
        </div>
      </template>
      <template #content>
        <div :style="{color: displays ? '#4CAF50' : '#F44336'}">{{displays ? displays.length : ''}} Displays</div>
        <CreateDisplay
          v-if="account"
          v-model="show_dialog"
          :createDisplay="createDisplay"
          :account_id="account.id"
        ></CreateDisplay>
      </template>
    </Card>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { ref } from "vue";
import web3Interface from "@/composables/web3Interface";
import accountManagement from "@/composables/accountManagement";
import displayManagement from "@/composables/displayManagement";
import CreateDisplay from "@/components/CreateDisplay.vue";
export default defineComponent({
  components: { CreateDisplay },
  setup() {
    // set refs
    const show_dialog = ref(false);
    // load composables
    const {
      connectWallet,
      setupWeb3Modal,
      signature,
      address,
      web3_modal,
    } = web3Interface();
    const { account, loadAccount } = accountManagement();
    const { displays, loadDisplays, createDisplay } = displayManagement();

    onMounted(setupWeb3Modal);

    // connect to web3 wallet, and setup firebase account details
    const connectAccount = async () => {
      await connectWallet();
      // signature and account should now be set properly

      await loadAccount(address.value, signature.value);
      if (account.value) {
        await loadDisplays(account.value.id);
      } else {
        alert("Tried to load displays but had no account");
      }
    };

    const openDialog = () => {
      show_dialog.value = true;
      console.log("OPEN DIALOG", show_dialog.value);
    };

    //getAccountInfo(address, signature);
    return {
      connectAccount,
      createDisplay,
      account,
      web3_modal,
      displays,
      signature,
      openDialog,
      show_dialog,
    };
  },
});
</script>
<style>
.controller-card {
  margin: 0 auto;
  max-width: 500px;
  margin-bottom: 15px;
}
</style>