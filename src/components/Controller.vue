<template>
  <div :style="{color: web3_modal ? '#4CAF50' : '#F44336'}">Web3 Modal</div>
  <div :style="{color: signature ? '#4CAF50' : '#F44336'}">Metamask</div>
  <div :style="{color: account ? '#4CAF50' : '#F44336'}">Account</div>
  <div :style="{color: displays ? '#4CAF50' : '#F44336'}">Displays</div>
  <button @click="connectAccount">Connect wallet</button>
</template>

<script lang="ts">
import { defineComponent, onMounted} from "vue";
import web3Interface from "@/composables/web3Interface";
import accountManagement from "@/composables/accountManagement"
import displayManagement from "@/composables/displayManagement"
export default defineComponent({
  setup() {
    const {
      connectWallet,
      setupWeb3Modal,
      signature,
      address,
      web3_modal
    } = web3Interface();

    const {account, loadAccount} = accountManagement()
    const {displays, loadDisplays} = displayManagement()

    onMounted(setupWeb3Modal);

    // connect to web3 wallet, and setup firebase account details
    const connectAccount = async () => {
      await connectWallet();
      // signature and account should now be set properly

      await loadAccount(address.value, signature.value);
      if(account.value) {
        await loadDisplays(account.value.id)
      } else {
        alert("Tried to load displays but had no account")
      }
    };

    //getAccountInfo(address, signature);
    return { connectAccount, account, web3_modal, displays, signature };
  },
});
</script>
<style>
</style>