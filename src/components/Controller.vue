<template>
  <div>
    <Card class="controller-card">
      <template #header>
        <div class="p-d-flex p-ai-center p-px-3">
          <h2>Account</h2>
          <div style="flex-grow: 1"></div>
          <Button
            v-if="!(web3_modal && signature && account)"
            label="Connect"
            :icon="loading_account ? 'pi pi-spin pi-spinner' : 'pi pi-wifi'"
            @click="connectAccount"
          />
          <Button v-else class="p-button-text" icon="pi pi-check" label="Connected"></Button>
        </div>
      </template>
      <template #content>
        <div class="p-d-flex p-jc-around">
          <Chip
            label="Web3"
            :icon="web3_modal ? 'pi pi-check' : 'pi pi-times'"
            :style="{color: web3_modal ? '#4CAF50' : '#F44336'}"
          ></Chip>
          <Chip
            label="Metamask"
            :icon="signature ? 'pi pi-check' : 'pi pi-times'"
            :style="{color: signature ? '#4CAF50' : '#F44336'}"
          ></Chip>
          <Chip
            label="Firebase"
            :icon="account ? 'pi pi-check' : 'pi pi-times'"
            :style="{color: account ? '#4CAF50' : '#F44336'}"
          ></Chip>
        </div>
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
import { defineComponent } from "vue";
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
    const loading_account = ref(false);
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

    // connect to web3 wallet, and setup firebase account details
    const connectAccount = async () => {
      loading_account.value = true;
      await setupWeb3Modal();

      await connectWallet();
      // signature and account should now be set properly

      await loadAccount(address.value, signature.value);

      loading_account.value = false;
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
      loading_account
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