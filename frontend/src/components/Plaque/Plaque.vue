<template>
  <div
    style="width: 100%; height: 100%; position: fixed;"
    :style="dark_mode ? 'background-color: #000000; color: #FFFFFF' : null"
  >

    <div
      style="transition: opacity ease-in 0.8s;"
      :class="show_text ? 'show-opacity' : 'hide-opacity'"
    >
      <div
        v-if="!token"
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
import { defineComponent, ref, watch, computed} from "vue";
import { OpenseaToken } from "@/types/types";
import QrcodeVue from "qrcode.vue";

export default defineComponent({
  props: {
    display_id: String
  },
  components: {
    QrcodeVue
  },
  setup() {
    const token = ref<OpenseaToken | null>();
    const show_text = ref(false);

    const dark_mode = computed(() => {
      return true;
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
      () => token.value,
      async () => {
        show_text.value = false;
        await new Promise(r => setTimeout(r, 800));

        show_text.value = true;
      },
      { immediate: true }
    );

    window.addEventListener("storage", () => {
      const token_data = window.localStorage.getItem("nft_token_data");
      if (token_data != null) {
        token.value = JSON.parse(token_data);
      }
      show_text.value =
        window.localStorage.getItem("nft_video_loaded") == "true";
    });
    return { token, dark_mode, show_text, top_bid };
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