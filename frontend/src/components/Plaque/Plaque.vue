<template>
  <div
    style="width: 100%; height: 100%; position: fixed;"
    :style="dark_mode ? 'background-color: #000000; color: #FFFFFF' : null"
  >
    <loading
      v-if="loading"
      active
    />
    <div
      v-else
      style="transition: opacity ease-in 0.8s;"
      :class="show_text ? 'show-opacity' : 'hide-opacity'"
    >
      <div
        v-if="!display"
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
            {{token?.name}}
          </div>
          <div class="p-grid">
            <div
              class="p-col"
              style="min-width: 350px;"
            >
              <div style="margin-bottom: 8px">{{token?.creator?.user?.username}}</div>
              <div
                v-if="top_bid"
                style="margin-bottom: 8px"
              >{{top_bid}}</div>
              <div>{{token?.description}}</div>
            </div>

            <div
              class="p-col"
              style="text-align: center;"
            >
              <qrcode-vue
                :value="token?.permalink"
                :size='200'
                :level="'M'"
              />
              <div style="font-size: 20px;">View on OpenSea</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, watch, computed, onMounted } from "vue";
import { loadToken } from "@/api/token";
import { FirestoreDocument, Display, OpenseaToken } from "@/types/types";
import Loading from "vue-loading-overlay";
import QrcodeVue from "qrcode.vue";
import {
  getDisplayByDisplayIDWithListener,
  getDisplayByDisplayID
} from "@/api/display";

export default defineComponent({
  props: {
    display_id: String
  },
  components: {
    QrcodeVue,
    Loading
  },
  setup() {
    const display = ref<FirestoreDocument<Display> | null>();
    const token = ref<OpenseaToken | null>();
    const loading = ref(false);
    const display_id = ref(window.localStorage.getItem("nft_display_id"));
    const show_text = ref(false);

    onMounted(() => {
      if (display_id.value) {
        show_text.value = true;
      }
    });

    const initDisplay = async (d: FirestoreDocument<Display>) => {
      show_text.value = false;
      await new Promise(r => setTimeout(r, 800));

      display.value = d;
      if (d.entity.token_id && d.entity.asset_contract_address) {
        const token_resp = await loadToken(
          d.entity.asset_contract_address,
          d.entity.token_id
        );
        console.log("token_resp", token_resp);
        token.value = token_resp;
      }
      loading.value = false;
    };

    const dark_mode = computed(() => {
      return display?.value?.entity.plaque_dark_mode;
    });

    const top_bid = computed(() => {
      const orders = token?.value?.orders;
      if (orders == null || orders?.length == 0) {
        return "";
      }
      let highest_order = orders[0];
      orders.forEach(o => {
        if (o.base_price > highest_order.base_price) {
          highest_order = o;
        }
      });
      //highest_order = orders[orders.length - 1];
      const bid =
        highest_order.base_price /
        Math.pow(10, highest_order.payment_token_contract.decimals);

      return `Top bid: ${bid.toFixed(3)} ${
        highest_order.payment_token_contract.symbol
      }`;
    });

    watch(
      () => display_id.value,
      async () => {
        loading.value = true;
        if (!window.navigator.onLine && display_id.value) {
          getDisplayByDisplayID(display_id.value).then(r => {
            initDisplay(r);
          });
        } else if (display_id.value) {
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
      show_text.value =
        window.localStorage.getItem("nft_video_loaded") == "true";
    });
    return { loading, token, display, dark_mode, show_text, top_bid };
  }
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

.hide-opacity {
  opacity: 0;
}
.show-opacity {
  opacity: 1;
}
</style>