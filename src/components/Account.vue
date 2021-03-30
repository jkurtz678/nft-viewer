<template>
  <button @click="connectWallet">Connect wallet</button>
</template>

<script lang="ts">

import {defineComponent} from "vue";
import Web3 from "web3";
import Web3Modal from "web3modal";
import { ethers } from "ethers";
import Authereum from "authereum";
import WalletConnectProvider from "@walletconnect/web3-provider";
export default defineComponent({
  data() {
    return {
      web3: {} as Web3,
      web3_modal: {} as Web3Modal, // modal used to authenticate with metamask, etc.
      address: "", // address that is used to sign data
      signature: "",
    };
  },
  mounted() {
    this.setupWeb3Modal();
  },
  methods: {
    setupWeb3Modal() {
      console.log("setup v3 modal...");
      /* if (location.protocol !== "https:") {
        console.error("must use https");
        return;
      } */
      this.web3_modal = new Web3Modal({
        providerOptions: {
          authereum: {
            package: Authereum, // required
          },
          walletconnect: {
            package: WalletConnectProvider.default,
            options: {
              infuraId: "9d5e849c49914b7092581cc71e3c2580",
            },
          },
        },
        cacheProvider: true,
        disableInjectedProvider: false,
      });
    },
    async connectWallet() {
      console.log("connecting wallet...");
      const provider = await this.web3_modal.connect().catch((err) => {
        console.error(err);
      });

      // Subscribe to accounts change
      provider.on("accountsChanged", () => {
        this.setWeb3Account(provider);
      });

      // Subscribe to chainId change
      provider.on("chainChanged", () => {
        this.setWeb3Account(provider);
      });

      // Subscribe to networkId change
      provider.on("networkChanged", () => {
        this.setWeb3Account(provider);
      });

      await this.setWeb3Account(provider);
    },
    async setWeb3Account(provider: any) {
      console.log("set web3 account...", provider);
      this.web3 = new Web3(provider);
      const web3_account = (await this.web3.eth.getAccounts())[0];
      console.log("web3_account", web3_account);
      this.address = web3_account;
      this.attemptReverse();
      await this.getSignature();
    },
    async attemptReverse() {
      console.log("attempt reverse...");
      let provider = new ethers.providers.Web3Provider(
        this.web3.currentProvider as ethers.providers.ExternalProvider
      );
      provider.lookupAddress(this.address).then((ensName) => {
        if (ensName != null) {
          this.address = ensName;
        }
      });
    },
    async getSignature() {
      console.log("get signature...");
      let plain =
        "NFT Viewer - proof of ownership. Please sign this message to prove ownership over your Ethereum account.";
      let msg = this.web3.utils.asciiToHex(plain);
      let hash = this.web3.utils.keccak256(
        "\x19Ethereum Signed Message:\n" + plain.length + plain
      );
      console.log("msg", msg);
      console.log("HASH", hash);
      console.log("this.web3_account", this.address);
      this.signature = await this.web3.eth.personal.sign(msg, this.address, "");
      //GetAccountInfo(true);
    },
  },
});
</script>

<style>
</style>