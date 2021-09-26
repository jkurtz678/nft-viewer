import Web3 from "web3";
import Web3Modal from "web3modal";
import { ethers } from "ethers";
import Authereum from "authereum";
import WalletConnectProvider from "@walletconnect/web3-provider";
import { ref } from 'vue'
import { useStore } from "vuex";

export default function web3Interface() {
  let web3: Web3;
  const web3_modal = ref<Web3Modal>();
  const address = ref("");
  const signature = ref("");
  const store = useStore();

  const setupWeb3Modal = () => {
    console.log("setup v3 modal...");
    /* if (location.protocol !== "https:") {
      console.error("must use https");
      return;
    } */
    return web3_modal.value = new Web3Modal({
      providerOptions: {
        authereum: {
          package: Authereum, // required
        },
        walletconnect: {
          package: WalletConnectProvider.default,
          options: {
            infuraId: "e132974b42d54791bd631e7bcd88572b", // infura.io Mainnet endpoint
          },
        },
      },
      cacheProvider: true,
      disableInjectedProvider: false,
    });
  };
  const connectWallet = async () => {
    console.log("connecting wallet...");
    if (!web3_modal.value) {
      alert("Tried to connect wallet but web3 modal not loaded!")
      return
    }
    const provider = await web3_modal?.value?.connect().catch((err) => {
      console.error(err);
    });

    // Subscribe to accounts change
    provider.on("accountsChanged", () => {
      setWeb3Account(provider);
    });

    // Subscribe to chainId change
    provider.on("chainChanged", () => {
      setWeb3Account(provider);
    });

    // Subscribe to networkId change
    provider.on("networkChanged", () => {
      setWeb3Account(provider);
    });

    return setWeb3Account(provider);
  };
  const setWeb3Account = async (provider: any) => {
    web3 = new Web3(provider);
    const web3_account = (await web3.eth.getAccounts())[0];
    address.value = web3_account;
    store.commit("setAddress", web3_account);
    attemptReverse();
    return getSignature();
  };
  const attemptReverse = async () => {
    const provider = new ethers.providers.Web3Provider(
      web3.currentProvider as ethers.providers.ExternalProvider
    );
    provider.lookupAddress(address.value).then((ensName) => {
      if (ensName != null) {
        address.value = ensName;
      }
    });
  };
  const getSignature = async () => {
    const local_signature = window.localStorage.getItem("web3_signature")
    if (local_signature) {
      signature.value = local_signature;
      return;
    }
    const plain =
      "NFT Viewer - proof of ownership. Please sign this message to prove ownership over your Ethereum account.";
    const msg = web3.utils.asciiToHex(plain);

    signature.value = await web3.eth.personal.sign(msg, address.value, "");
    window.localStorage.setItem("web3_signature", signature.value);
  };


  return {
    setupWeb3Modal,
    connectWallet,
    signature,
    address,
    web3_modal
  };
}