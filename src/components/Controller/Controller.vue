<template>
  <div>
    <Card class="controller-card" >
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
        <div v-if="tokens" class="p-mt-4">
          <TokenItem
            v-for="token of tokens"
            :key="token.name"
            :token="token"
          ></TokenItem>
        </div>
      </template>
    </Card>
    <DisplayController :account_id="account?.id"></DisplayController>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { ref } from "vue";
import web3Interface from "@/composables/web3Interface";
import accountManagement from "@/composables/accountManagement";
import DisplayController from "@/components/Controller/DisplayController.vue";
import TokenItem from "@/components/Controller/TokenItem.vue";

export default defineComponent({
  components: { DisplayController, TokenItem },
  setup() {
    // set refs
    const loading_account = ref(false);
    // load composables
    const {
      connectWallet,
      setupWeb3Modal,
      signature,
      address,
      web3_modal,
    } = web3Interface();
    const { account, tokens, loadAccount, loadTokens } = accountManagement();

    // connect to web3 wallet, and setup firebase account details
    const connectAccount = async () => {
      loading_account.value = true;
      await setupWeb3Modal();

      await connectWallet();
      // signature and account should now be set properly

      await loadAccount(address.value, signature.value);

      await loadTokens(address.value);

      loading_account.value = false;
    };

    //getAccountInfo(address, signature);
    return {
      connectAccount,
      account,
      web3_modal,
      signature,
      loading_account,
      tokens,
    };
  },
});
</script>
<style>
.controller-card {
  margin: 0 auto;
  max-width: 600px;
  margin-bottom: 15px;
  box-shadow: none !important;
}
</style>