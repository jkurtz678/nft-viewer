import { Token } from "../types/types"

// returns all tokens associated with this web3 account
export const loadTokens = async (web3_account_id: string): Promise<Array<Token>> => {
    const res = await fetch(`https://api.opensea.io/api/v1/assets/?owner=${web3_account_id}`);
    const res_json = await res.json();
    return res_json.assets;
}