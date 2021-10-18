import { Token } from "../types/types"
import firebase from "../firebaseConfig";

// returns all tokens associated with this web3 account
export const loadTokens = async (web3_account_id: string): Promise<Array<Token>> => {
    const res = await fetch(`https://api.opensea.io/api/v1/assets/?owner=${web3_account_id}`);
    const res_json = await res.json();

    return res_json.assets;
}

export const loadToken = async (asset_contract_address: string, token_id: string): Promise<Token> => {
    const res = await fetch(`https://api.opensea.io/api/v1/asset/${asset_contract_address}/${token_id}`)
    const res_json = await res.json()
    return res_json;
}

// loadArchiveMedia loads file url from firebase storage, returning null if not found
export const loadArchiveMedia = async (file_name: string): Promise<string | null> => {
    console.log("LOAD ARCHIVE", file_name)
    return firebase.storage().ref().child(file_name).getDownloadURL().catch(() => {
        return null
    })
}