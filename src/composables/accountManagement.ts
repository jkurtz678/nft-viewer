import { ref } from 'vue'
import { FirestoreDocument, Account } from "../types/types";
import * as account_api from "../api/account";

export default function manageAccounts() {
    const loading_account = ref(false)
    const account = ref<FirestoreDocument<Account> | null>()

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

    return { account, loading_account, loadAccount }
}