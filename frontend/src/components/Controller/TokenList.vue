<template>
  <div class="p-mt-2">
    <div style="display: flex; justify-content: space-between">
      <Dropdown
        v-if="!user_tokens"
        v-model="tag_filter"
        :options="tag_filter_options"
        optionLabel="name"
      />
      <Paginator
        v-model:first="offset"
        :rows="limit"
        :pageLinkSize="3"
        :totalRecords="total_records"
        :alwaysShow="false"
        class="p-p-0"
      >
        <template #FirstPageLink></template>
      </Paginator>
    </div>
    <TokenItem
      v-for="token of page_tokens"
      :key="token.name"
      :token="token"
      :selected="token.asset_contract.address == selected_asset_contract_address && token.token_id == selected_token_id"
      :in_playlist="playlist_tokens.some(t => t.asset_contract_address == token.asset_contract.address && t.token_id == token.token_id)"
      @click="selectToken(token.asset_contract.address, token.token_id)"
      @togglePlaylistToken="togglePlaylistToken(token, $event)"
    ></TokenItem>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed, watch, PropType } from "vue";
import { useStore } from "vuex";
import {
  OpenseaToken,
  TokenIDPair,
  FirestoreDocument,
  TokenMeta,
  DropdownOption
} from "@/types/types";
import TokenItem from "@/components/Controller/TokenItem.vue";
import { convertTokensToOpenseaFormat } from "@/api/token";
export default defineComponent({
  components: { TokenItem },
  props: {
    selected_token_id: String,
    selected_asset_contract_address: String,
    playlist_tokens: {
      type: Array as PropType<Array<TokenIDPair>>,
      default: () => []
    },
    user_tokens: Boolean
  },
  setup(props, { emit }) {
    const store = useStore();
    const demo_page_tokens = ref<Array<OpenseaToken>>([]);
    const limit = 10;
    const offset = ref(0);
    const tag_filter = ref<DropdownOption>({ name: "All", value: "All" });

    const filtered_token_metas = computed(
      (): Array<FirestoreDocument<TokenMeta>> => {
        let token_metas: Array<FirestoreDocument<TokenMeta>> =
          store.getters.demo_token_metas;
        if (!props.user_tokens && tag_filter.value.value != "All") {
          token_metas = token_metas.filter(m => {
            return m.entity.tag == tag_filter.value.value;
          });
        }
        return token_metas;
      }
    );

    const tag_filter_options = computed(
      (): Array<DropdownOption> => {
        let tag_map = new Map<string, boolean>();
        store.getters.demo_token_metas.forEach(
          (m: FirestoreDocument<TokenMeta>) => {
            if (m.entity.tag) {
              tag_map.set(m.entity.tag, true);
            }
          }
        );
        const tags: Array<DropdownOption> = Array.from(
          tag_map.keys()
        ).map(t => ({ name: t, value: t }));
        return [{ name: "All", value: "All" }, ...tags];
      }
    );

    const page_token_metas = computed(() => {
      const paginated_metas = filtered_token_metas.value.slice(
        offset.value,
        limit + offset.value
      );
      return paginated_metas;
    });

    const user_page_tokens = computed(() => {
      const user_tokens = store.getters.tokens.slice(
        offset.value,
        limit + offset.value
      );
      return user_tokens;
    });

    const page_tokens = computed(
      (): Array<OpenseaToken> => {
        if (props.user_tokens) {
          return user_page_tokens.value;
        }
        return demo_page_tokens.value;
      }
    );

    const total_records = computed(
      (): Number => {
        return props.user_tokens
          ? store.getters.tokens.length
          : filtered_token_metas.value.length;
      }
    );

    watch(
      page_token_metas,
      async v => {
        console.log("PAGE TOKEN METAS", v);
        demo_page_tokens.value = await convertTokensToOpenseaFormat(v);
      },
      { immediate: true, deep: true }
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
      total_records,
      limit,
      offset,
      user_page_tokens,
      tag_filter,
      tag_filter_options
    };
  }
});
</script>

