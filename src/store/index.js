import { createStore } from "vuex"

const store = createStore({
    state: {
        address: "",
        tokens: [],
        demo_tokens: [],
        camera_scan_mode: false,
    },
    getters: {
        address: state => {
            return state.address;
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
        setTokens(state, tokens) {
            state.tokens = tokens;
        },
        setDemoTokens(state, tokens) {
            state.demo_tokens = tokens;
        },
        setCameraScanMode(state, show) {
            state.camera_scan_mode = show;
        }
    }
})

export default store