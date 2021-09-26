import { ref } from 'vue'
import { FirestoreDocument, Account, Token } from "../types/types";
import * as account_api from "../api/account";
import * as token_api from "../api/token";
import { useStore } from "vuex";

export default function manageAccounts() {
    const loading_account = ref(false)
    const account = ref<FirestoreDocument<Account> | null>()
    const tokens = ref<Array<Token>>()

    const store = useStore();
    // gets account for address, creating if it does not exist
    const loadAccount = async (address: string, signature: string) => {
        loading_account.value = true
        const account_info = await account_api.getAccountByAddress(address)

        if (account_info.length > 1) {
            alert("Error! Multiple accounts for given address")
        } else if (account_info.length == 0) {
            // create new account if one was not found in the database
            console.log("creating account since none exist in the database...");
            account.value = await account_api.createAccount(address, signature);
        } else {
            account.value = account_info[0]
        }

        loading_account.value = false
    };

    const loadTokens = async (address: string) => {
        tokens.value = await token_api.loadTokens(address)
        store.commit("setTokens", tokens.value)
    }

    return { account, loading_account, tokens, loadAccount, loadTokens }
}