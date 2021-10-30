<template>
  <div v-if="store.getters.camera_scan_mode" style="text-align: right;">
    <Button
      label="Cancel scanning"
      icon="pi pi-ban"
      class="p-m-3"
      @click="store.commit('setCameraScanMode', false)"
    />
    <qrcode-stream @decode="onDecode"></qrcode-stream>
  </div>
  <div v-else>
    <Card class="controller-card">
      <template #header>
        <div class="p-d-flex p-ai-center p-px-3">
          <h2>Wallet</h2>
          <div style="flex-grow: 1"></div>
          <Button
            v-if="!(web3_modal && signature && account)"
            label="Connect"
            :icon="loading_account ? 'pi pi-spin pi-spinner' : 'pi pi-wifi'"
            @click="connectAccount"
          />
          <Button
            v-else
            class="p-button-text"
            icon="pi pi-check"
            label="Connected"
          ></Button>
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
        <div
          v-if="tokens"
          class="p-mt-4"
        >
          {{`${tokens.length} token${tokens.length == 1 ? '' : 's'} loaded from OpenSea account.`}}
        </div>
      </template>
    </Card>
    <DisplayController :account_id="account?.id"></DisplayController>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { ref } from "vue";
import { useStore } from "vuex";
import web3Interface from "@/composables/web3Interface";
import accountManagement from "@/composables/accountManagement";
import DisplayController from "@/components/Controller/DisplayController.vue";
import { QrcodeStream } from "vue-qrcode-reader";
import {
  getDisplayByDisplayID,
  updateDisplay,
  addAccountToDisplay,
} from "@/api/display";

export default defineComponent({
  components: { DisplayController, QrcodeStream },
  props: { display_id: String },
  setup(props) {
    const store = useStore();
    // set refs
    const loading_account = ref(false);
    // load composables
    const { connectWallet, setupWeb3Modal, signature, address, web3_modal } =
      web3Interface();
    const { account, tokens, loadAccount, loadTokens } = accountManagement();
    // connect to web3 wallet, and setup firebase account details
    const connectAccount = async () => {
      loading_account.value = true;
      await setupWeb3Modal();

      await connectWallet();
      // signature and account should now be set properly

      await loadAccount(address.value, signature.value);

      // add display if connected
      if (props.display_id) {
        const display = await getDisplayByDisplayID(props.display_id);
        display.entity.account_id = account.value?.id || "";
        await updateDisplay(display);
      }

      await loadTokens(address.value);

      loading_account.value = false;
    };

    const onDecode = async (qr_code_link: string) => {
      const display_id_index = qr_code_link.indexOf("display_id=");
      if (display_id_index > -1) {
        const display_id = qr_code_link.slice(display_id_index + 11);
        await addAccountToDisplay(display_id, account.value?.id || "");
      }
      store.commit("setCameraScanMode", false);
    };

    connectAccount();

    //getAccountInfo(address, signature);
    return {
      connectAccount,
      account,
      web3_modal,
      signature,
      loading_account,
      tokens,
      onDecode,
      store,
    };
  },
});
</script>
<style>
.controller-card {
  margin: 0 auto;
  max-width: 600px;
  margin-top: 40px;
  margin-bottom: 15px;
  box-shadow: none !important;
}
</style>