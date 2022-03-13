<template>
  <Card
    class="p-m-2"
    style="width: 130px; padding: 0px; text-align: center;"
    @click="$emit('editDisplay')"
  >
    <template #content>
      <img
        v-if="asset_url"
        :src="asset_url"
        style="height: 50px;"
      />
      <div v-else>No token</div>
    </template>
    <template #footer>

      <div class="p-text-bold">{{name}}</div>
    </template>
  </Card>
</template>

<script lang="ts">
import { defineComponent, computed } from "vue";
import { useStore } from "vuex";

export default defineComponent({
  props: {
    name: String,
    token_id: String
  },
  setup(props) {
    const store = useStore();
    const asset_url = computed(() => {
      const token = store.getters.token(props.token_id);
      if (token) {
        return token.image_thumbnail_url;
      }
      return "";
    });
    return { asset_url };
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