<template>
  <button @click="connectAccount">Connect wallet</button>
  <div>Address: {{account.address}}</div>
</template>

<script lang="ts">
import { defineComponent, ref} from "vue";
import { onMounted } from "vue";
import web3Interface from "@/composables/web3Interface";
import * as firebase_api from "../api/firebase";
import { EntityPayload, Account } from "../types/types";

export default defineComponent({
  setup() {
    const account = ref<EntityPayload<Account> | null>() 
    const {
      connectWallet,
      setupWeb3Modal,
      signature,
      address,
    } = web3Interface();

    onMounted(setupWeb3Modal);

    // connect to web3 wallet, and setup firebase account details
    const connectAccount = async () => {
      await connectWallet();
      // signature and account should now be set properly

      account.value = await getAccountDetails();
    };

    // gets account for address, creating if it does not exist
    const getAccountDetails = async (): Promise<EntityPayload<Account>> => {
      const account_info = await firebase_api.getAccount(address.value)

      if (account_info.length > 1) {
        throw "Error! Multiple accounts for given address";
      }
      // create new account if one was not found in the database
      if (account_info.length == 0) {
        console.log("creating account since none exist in the database...");
        return await firebase_api.createAccount(address.value, signature.value);
      }

      return account_info[0]
    };

    //getAccountInfo(address, signature);
    return { connectAccount, account };
  },
});
</script>
<style>
</style>