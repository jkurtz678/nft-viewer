import { ref } from 'vue'
import { FirestoreDocument, Display } from '@/types/types'
import * as display_api from "../api/display";

export default function displayManagement() {
    const displays = ref<Array<FirestoreDocument<Display>>>([])

    // loads saved displays for an account_id and sets the state
    const loadDisplays = async (account_id: string) => {
        const ret = await display_api.getDisplaysByAccountID(account_id).catch(alert)
        if (ret) {
            displays.value = ret;
        }
    }
    const createDisplay = async (account_id: string, name: string, code: string) => {
        const ret = await display_api.createDisplay(account_id, name).catch(alert)
        if (ret) {
            displays.value.push(ret);
        }
    }

    return { displays, loadDisplays, createDisplay }
}