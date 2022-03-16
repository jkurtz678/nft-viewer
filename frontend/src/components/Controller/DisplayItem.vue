<template>
  <Card
    class="p-m-2"
    style="width: 130px; padding: 0px; text-align: center;"
    @click="$emit('editDisplay')"
  >
    <template #content>
      <img
        v-if="has_image"
        :src="asset_str"
        style="height: 50px;"
      />
      <div v-else>{{asset_str}}</div>
    </template>
    <template #footer>

      <div class="p-text-bold">{{name}}</div>
    </template>
  </Card>
</template>

<script lang="ts">
import { defineComponent, computed } from "vue";
import { TokenMeta, FirestoreDocument} from "@/types/types";
import { useStore } from "vuex";

export default defineComponent({
  props: {
    name: String,
    asset_contract_address: String,
    token_id: String
  },
  setup(props) {
    const store = useStore();
    const has_image = computed(() => {
      return !!store.getters.token(props.token_id);
    })
    const asset_str = computed(() => {
      const token = store.getters.token(props.token_id);
      if (token) {
        return token.image_thumbnail_url;
      }
      const demo_token: FirestoreDocument<TokenMeta> = store.getters.demo_token_meta(
        props.asset_contract_address,
        props.token_id
      );
      if (demo_token) {
        return `${demo_token.entity.tag} demo token`;
      }
      return "No token";
    });
    return { asset_str, has_image};
  }
});
</script>

<style scoped>
.p-card :deep(.p-card-body) {
  padding: 0px;
  min-height: 130px;
}
.p-card :deep(.p-card-footer) {
  padding: 10px 0px 10px 0px;
}
.p-card :deep(.p-card-content) {
  padding-bottom: 0px;
}
</style>