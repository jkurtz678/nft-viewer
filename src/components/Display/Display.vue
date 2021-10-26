<template>
  <div>
    <loading
      v-if="loading"
      :active="true"
    ></loading>
    <template v-else>
      <template v-if="token">
        <div
          v-if="token.animation_url"
          class="video-container"
        >
          <video
            class="center"
            autoplay
            muted
            loop
            :src="media_url"
          ></video>
        </div>
      </template>
      <template v-else>
        <div
          v-if="display.entity.account_id"
          class="center"
          style="font-size: 40px"
        >Connected - No token selected</div>
        <template v-else>
          <div
            class="center"
            style="font-size: 40px"
          >
            <qrcode-vue
              :value="display_controller_url"
              :size='300'
              :level="'H'"
            />
            <div style="margin-top: 40px">Scan to cast NFTs</div>
          </div>
        </template>

      </template>
    </template>
  </div>
</template>

<script lang="ts">
import { ref, computed } from "vue";
import { defineComponent } from "vue";
import { useRouter } from "vue-router";
import { loadToken, loadArchiveMedia } from "@/api/token";
//mport { loadToken} from "@/api/token";
import { FirestoreDocument, Display, Token } from "@/types/types";
import {
  createDisplayWithListener,
  getDisplayByDisplayIDWithListener,
} from "@/api/display";
import QrcodeVue from "qrcode.vue";
import Loading from "vue-loading-overlay";
import "vue-loading-overlay/dist/vue-loading.css";
import VueViewer, { api as viewerApi, Viewer } from "v-viewer";
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
    const viewer = ref<Viewer>();
    const archive_media_url = ref<string | null>();

    const displayImage = (image_url: string) => {
      viewer.value = viewerApi({
        images: [image_url],
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
          scalable: true,
          transition: true,
          fullscreen: true,
          keyboard: false,
        },
      });
    };

    const initDisplay = async (d: FirestoreDocument<Display>) => {
      console.log("INIT DISPLAY", d);
      router.push({ path: "/display", query: { display_id: d.id } });
      display.value = d;
      window.localStorage.setItem("nft_display_id", d.id);
      if (d.entity.token_id && d.entity.asset_contract_address) {
        const token_resp = await loadToken(
          d.entity.asset_contract_address,
          d.entity.token_id
        );
        token.value = token_resp;


        archive_media_url.value = await loadArchiveMedia(token_resp.token_id + ".mp4")
        console.log("ARCHIVE MEDIA", archive_media_url.value) 

        // if token has no video media, display the image using viewer
        if (token_resp.animation_url) {
          if (viewer.value) {
            viewer.value.hide();
          }
        } else {
          displayImage(token_resp.image_url);
        }
      } else {
        token.value = null;
        if (viewer.value) {
          viewer.value.hide();
        }
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

    const media_url = computed(() => {
      if(archive_media_url.value) {
        return archive_media_url.value
      } 
      if(token.value?.animation_url) {
        return token.value.animation_url
      }
      if(token.value?.image_url) {
        return token.value.image_url;
      }
      return null;
    })

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
    return { display, loading, display_controller_url, token, media_url };
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
.video-container {
  position: absolute;
  top: 0;
  bottom: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
  background-color: black;
}
.video-container video {
  max-height: 100%;
  max-width: 100%;
}
</style>