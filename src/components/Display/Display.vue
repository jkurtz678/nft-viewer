<template>
  <div>
    <loading
      v-if="loading"
      :active="true"
    ></loading>
    <template v-else>
      <template v-if="token">
        <!--  <viewer :images="[token.image_url]" :options="{fullscreen: true, show: true}">
          
        </viewer> -->
      </template>
      <template v-else>
        <div
          v-if="display.entity.account_id"
          class="center"
          style="font-size: 40px"
        >No token selected</div>
        <qrcode-vue
          v-else
          :value="display_controller_url"
          :size='300'
          :level="'H'"
        />
      </template>
    </template>
  </div>
</template>

<script lang="ts">
import { ref, computed } from "vue";
import { defineComponent } from "vue";
import { useRouter } from "vue-router";
import { loadToken } from "@/api/token";
import { FirestoreDocument, Display, Token } from "@/types/types";
import {
  createDisplayWithListener,
  getDisplayByDisplayIDWithListener,
} from "@/api/display";
import QrcodeVue from "qrcode.vue";
import Loading from "vue-loading-overlay";
import "vue-loading-overlay/dist/vue-loading.css";
import VueViewer, { api as viewerApi } from "v-viewer";
VueViewer.setDefaults({
  zIndex: 2021,
});

export default defineComponent({
  props: {
    display_id: String,
  },
  components: {
    QrcodeVue,
    Loading,
  },
  setup(props) {
    const router = useRouter();
    const display = ref<FirestoreDocument<Display> | null>();
    const token = ref<Token | null>();
    const loading = ref(true);

    const initDisplay = async (d: FirestoreDocument<Display>) => {
      console.log("INIT DISPLAY", d);
      router.push({ path: "/display", query: { display_id: d.id } });
      display.value = d;
      window.localStorage.setItem("nft_display_id", d.id);
      if (d.entity.token_id && d.entity.asset_contract_address) {
        token.value = await loadToken(
          d.entity.asset_contract_address,
          d.entity.token_id
        );
        const $viewer = viewerApi({
          //images: [token.value.image_url],
          images: ["https://cdn.wallpapersafari.com/99/95/EYC9Zn.jpg"],
          options: {
            inline: false,
            button: false,
            navbar: false,
            title: false,
            toolbar: false,
            tooltip: false,
            movable: false,
            zoomable: false,
            rotatable: false,
            scalable: false,
            transition: true,
            fullscreen: true,
            keyboard: false,
          },
        });
        console.log("VIEWER", $viewer)
      } else {
        token.value = null;
      }
      loading.value = false;
    };

    const display_controller_url = computed(() => {
      console.log("document", document);
      if (display.value) {
        return `${document.location.origin}${document.location.pathname}#/controller?display_id=${display.value.id}`;
      }
      return "";
    });

    if (props.display_id) {
      getDisplayByDisplayIDWithListener(props.display_id, initDisplay);
    } else {
      const nft_display_id = localStorage.getItem("nft_display_id");
      if (nft_display_id) {
        getDisplayByDisplayIDWithListener(nft_display_id, initDisplay);
      } else {
        createDisplayWithListener("", "", "", initDisplay);
      }
    }

    return { display, loading, display_controller_url, token };
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