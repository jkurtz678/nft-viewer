<template>
  <div class="p-mt-4">
    <TabView>
      <TabPanel header="My tokens">
        <template v-if="store.getters.tokens">
          <TokenItem
            v-for="token of user_tokens"
            :key="token.name"
            :token="token"
            :selected="token.asset_contract.address == selected_asset_contract_address && token.token_id == selected_token_id"
            :in_playlist="playlist_tokens.some(t => t.asset_contract_address == token.asset_contract.address && t.token_id == token.token_id)"
            @click="selectToken(token.asset_contract.address, token.token_id)"
            @togglePlaylistToken="togglePlaylistToken(token, $event)"
          ></TokenItem>
        </template>
      </TabPanel>
      <TabPanel header="Demo tokens">
        <template v-if="page_tokens">
          <TokenItem
            v-for="token of page_tokens"
            :key="token.name"
            :token="token"
            :selected="token.asset_contract.address == selected_asset_contract_address && token.token_id == selected_token_id"
            :in_playlist="playlist_tokens.some(t => t.asset_contract_address == token.asset_contract.address && t.token_id == token.token_id)"
            @click="selectToken(token.asset_contract.address, token.token_id)"
            @togglePlaylistToken="togglePlaylistToken(token, $event)"
          ></TokenItem>
        </template>
      </TabPanel>
    </TabView>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed, watch, PropType } from "vue";
import { useStore } from "vuex";
import { OpenseaToken, TokenIDPair } from "@/types/types";
import TokenItem from "@/components/Controller/TokenItem.vue";
import { loadTokensByTokenIDAndAssetContract } from "@/api/token";
export default defineComponent({
  components: { TokenItem },
  props: {
    selected_token_id: String,
    selected_asset_contract_address: String,
    playlist_tokens: {
      type: Array as PropType<Array<TokenIDPair>>,
      default: () => []
    }
  },
  setup(props, { emit }) {
    const store = useStore();
    const page_tokens = ref<Array<OpenseaToken>>([]);
    const limit = 10;
    const page = ref<number>(1);

    const page_token_metas = computed(() => {
      const page_index = (page.value - 1) * limit;
      const token_metas = store.getters.demo_token_metas.slice(
        page_index,
        limit
      );
      return token_metas;
    });

    const user_tokens = computed(
      (): OpenseaToken => {
        return store.getters.tokens;
      }
    );

    watch(
      page_token_metas,
      async v => {
        page_tokens.value = await loadTokensByTokenIDAndAssetContract(v);
      },
      { immediate: true }
    );

    const selectToken = (asset_contract_address: string, token_id: string) => {
      if (
        asset_contract_address == props.selected_asset_contract_address &&
        token_id == props.selected_token_id
      ) {
        emit("update:selected_token_id", "");
        emit("update:selected_asset_contract_address", "");
      } else {
        emit("update:selected_token_id", token_id);
        emit("update:selected_asset_contract_address", asset_contract_address);
      }
    };
    const togglePlaylistToken = (token: OpenseaToken, toggle: boolean) => {
      let new_playlist: Array<TokenIDPair> = [];
      if (props.playlist_tokens) {
        new_playlist = JSON.parse(JSON.stringify(props.playlist_tokens));
      }

      if (toggle) {
        new_playlist.push({
          token_id: token.token_id,
          asset_contract_address: token.asset_contract.address
        });
      } else {
        new_playlist = new_playlist.filter(
          t =>
            t.token_id != token.token_id &&
            t.asset_contract_address != token.asset_contract.address
        );
      }
      emit("update:playlist_tokens", new_playlist);
    };

    return {
      selectToken,
      store,
      togglePlaylistToken,
      page_tokens,
      user_tokens
    };
  }
});
</script>

<style scoped>
.p-tabview :deep(.p-tabview-panels) {
  padding: 10px 0px 0px 0px;
}
</style>