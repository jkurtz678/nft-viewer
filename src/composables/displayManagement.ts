import { ref } from 'vue'
import { FirestoreDocument, Display } from '@/types/types'
import * as display_api from "../api/display";

export default function displayManagement() {
    const displays = ref<Array<FirestoreDocument<Display>>>()

    // loads saved displays for an account_id and sets the state
    const loadDisplays = async (account_id: string) => {
        const ret = await display_api.getDisplaysByAccountID(account_id).catch(alert)
        if (ret) {
            displays.value = ret;
        }
    }

    return { displays, loadDisplays }
}