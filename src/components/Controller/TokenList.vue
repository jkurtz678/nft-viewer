<template>
  <div
    v-if="$store.getters.tokens"
    class="p-mt-4"
  >
    <TabView>
      <TabPanel header="My tokens">
        <TokenItem
          v-for="token of $store.getters.tokens"
          :key="token.name"
          :token="token"
          :selected="token.token_id == selected_token_id"
          @click="selectToken(token.token_id)"
        ></TokenItem>
      </TabPanel>
      <TabPanel header="Demo tokens">
        <TokenItem
          v-for="token of $store.getters.demo_tokens"
          :key="token.name"
          :token="token"
          :selected="token.token_id == selected_token_id"
          @click="selectToken(token.token_id)"
        ></TokenItem>
      </TabPanel>
    </TabView>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import TokenItem from "@/components/Controller/TokenItem.vue";
export default defineComponent({
  components: { TokenItem },
  props: {
    selected_token_id: String,
  },
  setup(props, { emit }) {
    const selectToken = (token_id: string) => {
      console.log("SELECT TOKEN", token_id);
      if (token_id == props.selected_token_id) {
        emit("update:selected_token_id", "");
      } else {
        emit("update:selected_token_id", token_id);
      }
    };
    return { selectToken };
  },
});
</script>

<style>
</style>