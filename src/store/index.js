import { createStore } from "vuex"

const store = createStore({
    state: {
        address: "",
        tokens: [],
    },
    getters: {
        address: state => {
            return state.address;
        },
        tokens: state => {
            return state.tokens
        },
        token: state => token_id => {
            return state.tokens.find(t => t.token_id === token_id)
        }
    },
    mutations: {
        setAddress(state, address) {
            state.address = address;
        },
        setTokens(state, tokens) {
            state.tokens = tokens;
        }
    }
})

export default store