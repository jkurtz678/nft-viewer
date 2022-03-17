import { TokenMeta, FirestoreDocument, OpenseaToken } from "../types/types"
import firebase from "../firebaseConfig";

const db = firebase.firestore();

export const loadToken = async (asset_contract_address: string, token_id: string): Promise<OpenseaToken> => {
    const res = await fetch(`https://api.opensea.io/api/v1/asset/${asset_contract_address}/${token_id}?include_orders=true`)
    const res_json = await res.json()
    return res_json;
}

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
    return metas.filter(m => m.entity.token_id && m.entity.asset_contract_address)
}

// loadDemoTokenMeta returns a single demo token meta for a given asset_contract and token id
export const loadDemoTokenMeta = async (asset_contract_address: string, token_id: string): Promise<FirestoreDocument<TokenMeta>> => {
    const query_snapshot = await db.collection("demo_tokens").where("asset_contract_address", "==", asset_contract_address).where("token_id", "==", token_id).get()
    if(query_snapshot.docs.length != 1) {
        throw "Error - got duplicate token metas"
    }
    const doc = query_snapshot.docs[0]
    return {
        id: doc.id, entity: doc.data() as TokenMeta,
    }
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

// loadAndConvertTokenMeta will take a list of token metas, fetching their data from respective apis (opensea), otherwise using local data on the meta
export const loadAndConvertTokenMeta = async (tokens: Array<FirestoreDocument<TokenMeta>>): Promise<Array<OpenseaToken>> => {
    const opensea_only_tokens: Array<FirestoreDocument<TokenMeta>> = [];
    for (let i = 0; i < tokens.length; i++) {
        if (tokens[i].entity.platform == "opensea") {
            opensea_only_tokens.push(tokens[i])
        }
    }
    const opensea_res = await loadTokensByTokenIDAndAssetContract(opensea_only_tokens)
    const converted_tokens: Array<OpenseaToken> = [];
    for (let i = 0; i < tokens.length; i++) {
        const t = tokens[i]
        // if token is opensea, add from api response
        if (tokens[i].entity.platform == "opensea") {
            const pop = opensea_res.shift()
            if (pop) {
                converted_tokens.push(pop)
            }
        } else { // otherwise convert token meta to a similar format
            const converted: OpenseaToken = convertTokenMetaToOpensea(t)
            converted_tokens.push(converted)
        }
    }
    return converted_tokens
}

// convertTokenMetaToOpensea will convert a token meta to the OpenseaToken format
export const convertTokenMetaToOpensea = (t: FirestoreDocument<TokenMeta>): OpenseaToken => {
    return {
        name: t.entity.name,
        asset_contract: { address: t.entity.asset_contract_address },
        token_id: t.entity.token_id,
        description: t.entity.description,
        creator: { user: { username: t.entity.tag } },
        image_url: "",
        animation_url: "",
        image_thumbnail_url: "",
        orders: [],
        permalink: ""
    }
}



// loadArchiveMedia loads file url from firebase storage, returning null if not found
export const loadArchiveMedia = async (file_name: string): Promise<string | null> => {
    return firebase.storage().ref().child(file_name).getDownloadURL().catch(() => {
        return null
    })
}