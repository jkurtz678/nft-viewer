<template>
  <div>
    <loading
      v-if="loading"
      :active="true"
    ></loading>
    <template v-else>
      <template v-if="token">
        <div
          v-if="media_is_video"
          class="video-container"
        >
          <video
            :class="show_video ? 'show-opacity' : 'hide-opacity'"
            ref="player"
            class="center"
            autoplay
            muted
            :loop="video_should_loop"
            :src="media_url"
          ></video>
        </div>
      </template>
      <template v-if="!loading && !token">
        <div
          v-if="display?.entity?.account_id"
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
import { ref, computed, watch } from "vue";
import { defineComponent } from "vue";
import { useRouter } from "vue-router";
import {
  convertTokenMetaToOpensea,
  loadDemoTokenMeta,
  loadToken
} from "@/api/token";
import { getLocalFileURL, findLocalFile } from "@/api/local-files";
import {
  FirestoreDocument,
  Display,
  OpenseaToken,
  TokenMeta
} from "@/types/types";
import {
  createDisplayWithListener,
  getDisplayByDisplayIDWithListener,
  updateDisplay
} from "@/api/display";

import QrcodeVue from "qrcode.vue";
import Loading from "vue-loading-overlay";
import "vue-loading-overlay/dist/vue-loading.css";
import VueViewer, { api as viewerApi, Viewer } from "v-viewer";
VueViewer.setDefaults({
  zIndex: 2021
});

export default defineComponent({
  props: {
    display_id: String
  },
  components: {
    QrcodeVue,
    Loading
  },
  setup(props) {
    const router = useRouter();
    const display = ref<FirestoreDocument<Display> | null>();
    const token = ref<OpenseaToken | null>();
    const token_meta = ref<FirestoreDocument<TokenMeta> | null>();
    const loading = ref(true);
    const show_video = ref(false);
    const viewer = ref<Viewer>();
    const local_file = ref<string>();
    const player = ref<HTMLVideoElement>();
    const playlist_timer = ref<ReturnType<typeof setTimeout>>();

    /* START COMPUTED */
    // media_is_video determines if token media is video (true) or an image/gif (false)
    const media_is_video = computed((): boolean => {
      const is_opensea_video = !!token.value?.animation_url &&
        !token.value.animation_url.endsWith(".wav")
      const is_alt_video = !!token_meta?.value && !!local_file.value?.endsWith(".mp4")
      return is_opensea_video || is_alt_video;
    });

    // display_controller_url returns the url for the qrcode, which allows people to scan the display with their mobile phone and control it with the controller webapp
    const display_controller_url = computed((): string => {
      if (display.value) {
        return `https://jkurtz678.github.io/nft-viewer/${document.location.pathname}#/controller?display_id=${display.value.id}`;
      }
      return "";
    });

    // media_url decides which url should be used for the media display
    const media_url = computed((): string | null => {
      if (local_file.value && token.value?.token_id) {
        return getLocalFileURL(local_file.value);
      }
      if (media_is_video.value && token.value) {
        return token.value.animation_url;
      }
      if (token.value?.image_url) {
        return token.value.image_url;
      }
      return null;
    });

    // video_should_loop will be true by default, false if any playlist tokens exist
    const video_should_loop = computed((): boolean => {
      const playlist_tokens = display.value?.entity?.playlist_tokens;
      return playlist_tokens == null || playlist_tokens.length == 0 || playlist_timer.value != null;
    });
    /* END COMPUTED */

    /* START WATCHERS */
    // wait until html player loads to add the video ended event listener
    watch(player, v => {
      if (v) {
        /* player.value?.removeEventListener("ended", nextPlaylistToken); // remove any existing listener if it exists
        player.value?.addEventListener("ended", nextPlaylistToken); */
      }
    });
    /* END WATCHERS */

    /* START METHODS */
    // initDisplay will handle changes made to the data of a display, showing/hiding media
    const initDisplay = async (d: FirestoreDocument<Display>) => {
      console.log("initDisplay", d);
      show_video.value = false; // tells display to start fade out
      window.localStorage.setItem("nft_video_loaded", "false"); // use local storage to tell the plaque to fade out text
      if (playlist_timer.value) {
        clearTimeout(playlist_timer.value);
      }
      await new Promise(r => setTimeout(r, 800)); // wait so previous video has time to fade out

      router.push({ path: "/display", query: { display_id: d.id } });
      display.value = d;
      window.localStorage.setItem("nft_display_id", d.id);

      // to load a token, we need both the token id and the contract address. if the display is missing one of these, clear out the token from display
      if (!d.entity.token_id || !d.entity.asset_contract_address) {
        token.value = null;
        local_file.value = "";
        if (viewer.value) {
          viewer.value.hide();
        }
        loading.value = false;
        return;
      }

      // now its safe to load and show the token media
      showToken(d.entity.asset_contract_address, d.entity.token_id);
    };

    // showToken will load a token from opensea and display its associated media. Different behaviors for videos and images/gifs
    const showToken = async (contract_address: string, token_id: string) => {
      try {
        token_meta.value = await loadDemoTokenMeta(contract_address, token_id);
        console.log("loadDemoTokenMeta", token_meta.value);
      } catch {
        console.log("No token meta found");
        token_meta.value = null;
      }
      let token_resp: OpenseaToken;
      if (!token_meta.value || token_meta.value.entity.platform == "opensea") {
        token_resp = await loadToken(contract_address, token_id);
      } else {
        token_resp = convertTokenMetaToOpensea(token_meta.value);
      }
      token.value = token_resp;
      console.log("tokenValue", token.value);

      try {
        local_file.value = await findLocalFile(token.value.token_id);
      } catch (err) {
        local_file.value = "";
        console.log(err);
      }

      // if token has no video media, display the image using viewer
      if (media_is_video.value) {
        if (viewer.value) {
          viewer.value.hide();
        }
        waitToShowVideo();
      } else {
        if (media_url.value) {
          displayImage(media_url.value);
          playlist_timer.value = setTimeout(() => {
            nextPlaylistToken();
          }, 1000 * 45);
        }
      }
      loading.value = false;
    };

    // waitToShowVideo waits until video has started playing before fading it in
    const waitToShowVideo = () => {
      const checkInterval = 50.0; // check every 50 ms to see if video has started playing
      let current_play_pos = 0;

      const interval = setInterval(checkBuffering, checkInterval);
      function checkBuffering() {
        if (!player.value) {
          console.log("ERROR: Display video player not found.");
          return;
        }

        current_play_pos = player.value.currentTime;
        if (current_play_pos > 0) {
          clearInterval(interval);
          show_video.value = true;
          window.localStorage.setItem("nft_video_loaded", "true");

          // handle playlist cycling
          player.value?.removeEventListener("ended", nextPlaylistToken); // remove any existing listener if it exists
          console.log("video duration: ", player.value.duration)
          if (player.value.duration && player.value.duration < 15) {
            playlist_timer.value = setTimeout(() => {
              nextPlaylistToken();
            }, 1000 * 45);
          } else {
            player.value?.addEventListener("ended", nextPlaylistToken);
          }
        }
      }
    };

    const displayImage = (image_url: string) => {
      console.log("displayImage", image_url);
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
          keyboard: false
        }
      });
      window.localStorage.setItem("nft_video_loaded", "true");
    };

    // nextPlaylistToken will update the display to the next token in the playlist
    const nextPlaylistToken = () => {
      console.log("nextPlaylistToken")
      if (!display.value?.entity?.playlist_tokens?.length) {
        return;
      }
      // find new index of next playlist item
      let new_index =
        display.value.entity.playlist_tokens.findIndex(
          t => t.token_id == token.value?.token_id
        ) + 1;
      if (
        new_index == -1 ||
        new_index >= display.value.entity.playlist_tokens.length
      ) {
        new_index = 0;
      }

      display.value.entity.token_id =
        display.value.entity.playlist_tokens[new_index].token_id;
      display.value.entity.asset_contract_address =
        display.value.entity.playlist_tokens[new_index].asset_contract_address;

      updateDisplay(display.value);
    };
    /* END METHODS */

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

    return {
      display,
      loading,
      show_video,
      display_controller_url,
      token,
      media_url,
      player,
      video_should_loop,
      local_file,
      media_is_video
    };
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
.video-container {
  transition: opacity ease-in 0.8s;
  position: absolute;
  top: 0;
  bottom: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
  background-color: black;
}
.hide-opacity {
  opacity: 0;
}
.show-opacity {
  opacity: 1;
}
.video-container video {
  transition: opacity ease-in 0.8s;
  max-height: 100%;
  max-width: 100%;
}
</style>