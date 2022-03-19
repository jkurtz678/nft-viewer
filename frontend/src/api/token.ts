import { TokenMeta, FirestoreDocument, OpenseaToken } from "../types/types"
import firebase from "../firebaseConfig";

const db = firebase.firestore();

// returns all tokens associated with this web3 account
export const loadTokens = async (web3_account_id: string): Promise<Array<OpenseaToken>> => {
    const res = await fetch(`https://api.opensea.io/api/v1/assets/?owner=${web3_account_id}`);
    const res_json = await res.json();

    return res_json.assets;
}

// loadDemoTokenMetas returns a list of demo tokens which all controllers have access to
export const loadDemoTokenMetas = async (): Promise<Array<FirestoreDocument<TokenMeta>>> => {
    const query_snapshot = await db.collection("demo_tokens").get();
    const metas = query_snapshot.docs.map(s => ({
        id: s.id, entity: s.data() as TokenMeta,
    }))
    return metas.filter(m => m.entity.platform == "opensea")
}

export const loadTokensByTokenIDAndAssetContract = async (tokens: Array<FirestoreDocument<TokenMeta>>): Promise<Array<OpenseaToken>> => {
    let url = `https://api.opensea.io/api/v1/assets/?`
    tokens.forEach(t => {
        url += `asset_contract_addresses=${t.entity.asset_contract_address}&token_ids=${t.entity.token_id}&`
    })

    url = url.substring(0, url.length - 1)

    const res = await fetch(url);
    const res_json = await res.json();

    return res_json.assets;
}

export const loadToken = async (asset_contract_address: string, token_id: string): Promise<OpenseaToken> => {
    const req_url = `https://api.opensea.io/api/v1/asset/${asset_contract_address}/${token_id}?include_orders=true`;

    try {
        const cache_res = await fetch(`http://localhost:8081/api/cache?req=${encodeURIComponent(req_url)}`)
        if (cache_res.status < 400) {
            const json_cache_res = await cache_res.json()
            return json_cache_res.data
        }
    } catch (err) {
        console.log("no cached data found")
    }

    const res = await fetch(req_url)
    const res_json = await res.json()

    // write cached data
    try {
        await fetch(`http://localhost:8081/api/cache`, {
            method: "POST",
            body: JSON.stringify({
                req: req_url,
                res: res_json
            })
        })
    } catch (err) {
        console.error(err)
    }

    return res_json;
}

// loadArchiveMedia loads file url from firebase storage, returning null if not found
export const loadArchiveMedia = async (file_name: string): Promise<string | null> => {
    return firebase.storage().ref().child(file_name).getDownloadURL().catch(() => {
        return null
    })
}