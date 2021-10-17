<template>
  <loading
    v-if="loading"
    active
  />
  <div
    v-else-if="!display"
    class="center"
    style="font-size: 40px;"
  >No display found</div>
  <div
    v-else-if="display && !display.entity.token_id"
    class="center"
    style="font-size: 40px;"
  >No media loaded</div>
  <div v-else style='text-align: left' class="center p-px-6">
    <div style="font-size: 40px;">
      {{token.name}}
    </div>
    <div
      class="p-grid"
    >

      <div
        class="p-col"
        style="min-width: 270px;"
      >

        <div style="margin-bottom: 8px">{{token.creator.user.username}}</div>
        <div>{{token.description}}</div>
      </div>

      <div class="p-col" style="text-align: center;">
        <qrcode-vue
          :value="token.permalink"
          :size='200'
          :level="'M'"
        />
        <div style="font-size: 20px;">View on OpenSea</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, watch } from "vue";
import { loadToken } from "@/api/token";
import { FirestoreDocument, Display, Token } from "@/types/types";
import Loading from "vue-loading-overlay";
import QrcodeVue from "qrcode.vue";
import { getDisplayByDisplayIDWithListener } from "@/api/display";

export default defineComponent({
  props: {
    display_id: String,
  },
  components: {
    QrcodeVue,
    Loading,
  },
  setup() {
    const display = ref<FirestoreDocument<Display> | null>();
    const token = ref<Token | null>();
    const loading = ref(false);
    const display_id = ref(window.localStorage.getItem("nft_display_id"));

    const initDisplay = async (d: FirestoreDocument<Display>) => {
      display.value = d;
      console.log("d", d);
      if (d.entity.token_id && d.entity.asset_contract_address) {
        const token_resp = await loadToken(
          d.entity.asset_contract_address,
          d.entity.token_id
        );
        console.log("TOKEN RESP", token_resp);
        token.value = token_resp;
      }
      loading.value = false;
    };

    watch(
      () => display_id.value,
      async () => {
        console.log("DISPALY ID", display_id.value);
        loading.value = true;
        if (display_id.value) {
          getDisplayByDisplayIDWithListener(display_id.value, initDisplay);
        } else {
          display.value = null;
          loading.value = false;
        }
      },
      { immediate: true }
    );

    window.addEventListener("storage", () => {
      display_id.value = window.localStorage.getItem("nft_display_id");
    });
    return { loading, token, display };
  },
});
</script>

<style scoped>
.center {
  margin: 0;
  position: absolute;
  top: 50%;
  -ms-transform: translate(0%, -50%);
  transform: translate(0%, -50%);
}
</style>