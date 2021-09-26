import { Token } from "../types/types"

// returns all tokens associated with this web3 account
export const loadTokens = async (web3_account_id: string): Promise<Array<Token>> => {
    const res = await fetch(`https://api.opensea.io/api/v1/assets/?owner=${web3_account_id}`);
    const res_json = await res.json();

    return res_json.assets;
}

export const loadToken = async (asset_contract_address: string, token_id: string): Promise<Token> => {
    const res = await fetch(`https://api.opensea.io/api/v1/assets/?asset_contract_address=${asset_contract_address}&token_id=${token_id}`)
    const res_json = await res.json()
    console.log("LOAD TOKEN RES", res_json)
    return res_json;
}