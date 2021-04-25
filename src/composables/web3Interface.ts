import Web3 from "web3";
import Web3Modal from "web3modal";
import { ethers } from "ethers";
import Authereum from "authereum";
import WalletConnectProvider from "@walletconnect/web3-provider";
import firebase from "../firebaseConfig";
const db = firebase.firestore();

export default function web3Interface() {
  let web3: Web3;
  let web3_modal: Web3Modal;
  let address: string = "";
  let signature: string = "";

  const setupWeb3Modal = () => {
    console.log("setup v3 modal...");
    /* if (location.protocol !== "https:") {
      console.error("must use https");
      return;
    } */
    web3_modal = new Web3Modal({
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
  };
  const connectWallet = async () => {
    console.log("connecting wallet...");
    const provider = await web3_modal.connect().catch((err) => {
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

    await setWeb3Account(provider);
  };
  const setWeb3Account = async (provider: any) => {
    console.log("set web3 account...", provider);
    web3 = new Web3(provider);
    const web3_account = (await web3.eth.getAccounts())[0];
    console.log("web3_account", web3_account);
    address = web3_account;
    attemptReverse();
    await getSignature();
  };
  const attemptReverse = async () => {
    console.log("attempt reverse...");
    const provider = new ethers.providers.Web3Provider(
      web3.currentProvider as ethers.providers.ExternalProvider
    );
    provider.lookupAddress(address).then((ensName) => {
      if (ensName != null) {
        address = ensName;
      }
    });
  };
  const getSignature = async () => {
    console.log("get signature...");
    const plain =
      "NFT Viewer - proof of ownership. Please sign this message to prove ownership over your Ethereum account.";
    const msg = web3.utils.asciiToHex(plain);
    /* let hash = this.web3.utils.keccak256(
      "\x19Ethereum Signed Message:\n" + plain.length + plain
    ); */
    signature = await web3.eth.personal.sign(msg, address, "");
    getAccountInfo();
  };
  // create one if it doesn't exist
  const getAccountInfo = async () => {
    db.collection("account")
      .add({ name: "some name" })
      .catch((err: string) => {
        console.error(err);
      });
  };

  return {
    setupWeb3Modal,
    connectWallet,
    signature,
  };
}