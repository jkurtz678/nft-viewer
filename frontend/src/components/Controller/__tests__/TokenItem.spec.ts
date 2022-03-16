import { shallowMount } from '@vue/test-utils';
import TokenItem from "../TokenItem.vue";

describe('TokenItem.vue', () => {
    beforeEach(async () => {

    })
    it("should init", async () => {
        //expect(wrapper.vm).toBeTruthy()
        const wrapper = await shallowMount(TokenItem, {
            props: {
                token: {},
                selected: false,
                in_playlist: false,
            }
        })
    });
})