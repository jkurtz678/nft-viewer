<template>
  <div style="width: 100%; height: 100%; position: fixed;" :style="dark_mode ? 'background-color: #000000; color: #FFFFFF' : null">
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
    <div
      v-else
      class="center p-px-6"
      style="text-align: left; width: 100%"
    >
      <div style="max-width: 800px; margin-left: auto; margin-right: auto">
        <div style="font-size: 40px;">
          {{token.name}}
        </div>
        <div class="p-grid">
          <div
            class="p-col"
            style="min-width: 350px;"
          >

            <div style="margin-bottom: 8px">{{token?.creator?.user?.username}}</div>
            <div>{{token.description}}</div>
          </div>

          <div
            class="p-col"
            style="text-align: center;"
          >
            <qrcode-vue
              :value="token.permalink"
              :size='200'
              :level="'M'"
            />
            <div style="font-size: 20px;">View on OpenSea</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, watch, computed } from "vue";
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
      if (d.entity.token_id && d.entity.asset_contract_address) {
        const token_resp = await loadToken(
          d.entity.asset_contract_address,
          d.entity.token_id
        );
        token.value = token_resp;
      }
      loading.value = false;
    };

    const dark_mode = computed( () => {
      return display?.value?.entity.plaque_dark_mode
    })

    watch(
      () => display_id.value,
      async () => {
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
    return { loading, token, display, dark_mode };
  },
});
</script>

<style scoped>
.center {
  margin: 0;
  position: absolute;
  top: 50%;
  left: 50%;
  -ms-transform: translate(-50%, -50%);
  transform: translate(-50%, -50%);
}
</style>