import { ref } from 'vue'
import { FirestoreDocument, Display } from '@/types/types'
import * as display_api from "../api/display";
import { customAlphabet } from 'nanoid';

export default function displayManagement() {
    const displays = ref<Array<FirestoreDocument<Display>>>([])
    const displays_loaded = ref(false)

    // loads saved displays for an account_id and sets the state
    const loadDisplays = async (account_id: string) => {
        const ret = await display_api.getDisplaysByAccountID(account_id).catch(alert)
        if (ret) {
            displays.value = ret;
            displays_loaded.value = true;
        }
    }
    const createDisplay = async (account_id: string, name: string) => {
        const nanoid = customAlphabet("1234567890", 6)
        const code = nanoid()
        const ret = await display_api.createDisplay(account_id, name, code).catch(alert)
        if (ret) {
            displays.value.push(ret);
        }
    }

    return { displays, displays_loaded, loadDisplays, createDisplay }
}