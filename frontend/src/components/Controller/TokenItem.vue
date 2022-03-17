<template>
  <div>
    <Card :style="selected ? 'border-style: solid; -webkit-box-sizing: border-box;' : ''">
      <template #content>
        <div class="p-grid">
          <div
            class="p-col-6"
            style="margin: auto; text-align: left; padding-left: 20px"
          >
            <div>{{token?.name}}</div>
            <div>{{token_meta?.entity?.tag}}</div>
          </div>
          <div class="p-col-4">
            <template v-if="token_meta?.entity && token_meta?.entity?.platform != 'opensea'">Local only</template>
            <img
              v-else
              :src="token?.image_thumbnail_url"
              style="height: 50px;"
            />
          </div>
          <div
            class="p-col-2"
            style="margin: auto"
          >
            <Button
              v-if="in_playlist"
              class="p-button-text"
              icon="pi pi-list"
              :style="{color: '#000000'}"
              @click.stop='$emit("togglePlaylistToken", false)'
            ></Button>
            <Button
              v-else
              class="p-button-text"
              icon="pi pi-plus"
              @click.stop='$emit("togglePlaylistToken", true)'
            >
            </Button>
          </div>
        </div>
      </template>
    </Card>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from "vue";
import { useStore } from "vuex";
import { OpenseaToken, TokenMeta, FirestoreDocument } from "../../types/types";
export default defineComponent({
  props: {
    token: Object as () => OpenseaToken,
    selected: Boolean,
    in_playlist: Boolean
  },
  setup(props) {
    const store = useStore();
    const token_meta = computed(
      (): FirestoreDocument<TokenMeta> => {
        return store.getters.demo_token_meta(
          props.token?.asset_contract.address,
          props.token?.token_id
        );
      }
    );
    return { token_meta };
  }
});
</script>

<style scoped>
.p-card {
  width: 100%;
  margin-bottom: 5px;
  text-align: center;
  box-shadow: none;
}
.p-card :deep(.p-card-body) {
  padding: 0px;
}
.p-card :deep(.p-card-footer) {
  padding: 10px 0px 10px 0px;
}
.p-card :deep(.p-card-content) {
  padding-bottom: 0px;
}
</style>