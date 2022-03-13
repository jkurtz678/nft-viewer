<template>
  <div
    v-if="store.getters.tokens"
    class="p-mt-4"
  >
    <TabView>
      <TabPanel header="My tokens">
        <TokenItem
          v-for="token of store.getters.tokens"
          :key="token.name"
          :token="token"
          :selected="token.token_id == selected_token_id"
          :in_playlist="playlist_tokens.some(t => t.token_id == token.token_id)"
          @click="selectToken(token.token_id)"
          @togglePlaylistToken="togglePlaylistToken(token, $event)"
        ></TokenItem>
      </TabPanel>
      <TabPanel header="Demo tokens">
        <TokenItem
          v-for="token of store.getters.demo_tokens"
          :key="token.name"
          :token="token"
          :selected="token.token_id == selected_token_id"
          :in_playlist="playlist_tokens.some(t => t.token_id == token.token_id)"
          @click="selectToken(token.token_id)"
          @togglePlaylistToken="togglePlaylistToken(token, $event)"
        ></TokenItem>
      </TabPanel>
    </TabView>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useStore } from "vuex";
import {Token } from "@/types/types"
import TokenItem from "@/components/Controller/TokenItem.vue";
export default defineComponent({
  components: { TokenItem },
  props: {
    selected_token_id: String,
    playlist_tokens: {
      type: Array,
      default: () => []
    },
  },
  setup(props, { emit }) {
    const store = useStore();
    const selectToken = (token_id: string) => {
      if (token_id == props.selected_token_id) {
        emit("update:selected_token_id", "");
      } else {
        emit("update:selected_token_id", token_id);
      }
    };
    const togglePlaylistToken = (token: Token, toggle: boolean) => {
      console.log("TOGGLE PLAYLIST", token, toggle)
      let new_playlist: Array<any> = [];
      if (props.playlist_tokens) {
        new_playlist = JSON.parse(JSON.stringify(props.playlist_tokens));
      }

      if(toggle) {
        new_playlist.push({token_id: token.token_id, asset_contract_address: token.asset_contract.address})
      } else {
        new_playlist = new_playlist.filter( t => t.token_id != token.token_id)
      }
      console.log("EMIT PLAYLIST TOKENS", new_playlist)
      emit("update:playlist_tokens", new_playlist)
    }
    return { selectToken, store, togglePlaylistToken};
  },
});
</script>

<style scoped>
.p-tabview :deep(.p-tabview-panels) {
  padding: 10px 0px 0px 0px;
}
</style>