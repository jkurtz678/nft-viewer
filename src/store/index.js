import { createStore } from "vuex"

const store = createStore({
    state: {
        address: "",
        tokens: [],
        demo_tokens: [],
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
        }
    }
})

export default store