<template>
  <button @click="connectWallet">Connect wallet</button>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { onMounted } from "vue";
import web3Interface from "@/composables/web3Interface"
export default defineComponent({
  setup() {
    const { connectWallet, setupWeb3Modal } = web3Interface();
    onMounted(setupWeb3Modal)
    return {connectWallet}
  },
});
</script>
<style>
</style>