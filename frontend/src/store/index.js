import { createStore } from "vuex"

const store = createStore({
    state: {
        address: "",
        account: {},
        tokens: [],
        demo_tokens: [],
        demo_token_metas: [],
        camera_scan_mode: false,
    },
    getters: {
        address: state => {
            return state.address;
        },
        account: state => {
            return state.account;
        },
        demo_token_metas: state => {
            return state.demo_token_metas
        },
        demo_token_meta: state => (asset_contract_address, token_id) =>  {
            return state.demo_token_metas.find(t => t.asset_contract_address == asset_contract_address && t.token_id == token_id)
        },
        tokens: state => {
            return state.tokens;
        },
        token: state => token_id => {
            const token = state.tokens.find(t => t.token_id === token_id);
            if(token){
                return token;
            } else {
                return state.demo_tokens.find(t => t.token_id === token_id);
            }
        },
        demo_tokens: state => {
            return state.demo_tokens; 
        },
        camera_scan_mode: state => {
            return state.camera_scan_mode;
        }
    },
    mutations: {
        setAddress(state, address) {
            state.address = address;
        },
        setAccount(state, account) {
            state.account = account;
        },
        setTokens(state, tokens) {
            state.tokens = tokens;
        },
        setDemoTokens(state, tokens) {
            state.demo_tokens = tokens;
        },
        setDemoTokenMetas(state, token_metas) {
            state.demo_token_metas = token_metas
        }, 
        setCameraScanMode(state, show) {
            state.camera_scan_mode = show;
        }
    }
})

export default store