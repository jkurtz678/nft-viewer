<template>
  <div>
    <loading
      v-if="loading"
      :active="true"
    ></loading>
    <template v-else>
      <template v-if="display.entity.media_url">
        <img :src="display.entity.media_url" />
      </template>
      <template v-else>
        <qrcode-vue
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
import { FirestoreDocument, Display } from "@/types/types";
import {
  createDisplayWithListener,
  getDisplayByDisplayIDWithListener,
} from "@/api/display";
import QrcodeVue from "qrcode.vue";
import Loading from "vue-loading-overlay";
import "vue-loading-overlay/dist/vue-loading.css";

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
    const loading = ref(true);
    const initDisplay = (d: FirestoreDocument<Display>) => {
      router.push({ path: "/display", query: { display_id: d.id } });
      display.value = d;
      window.localStorage.setItem("nft_display_id", d.id);
      loading.value = false;
    };

    const display_controller_url = computed(() => {
      if (display.value) {
        return `http://localhost:8080/#/controller?display_id=${display.value.id}`;
      }
      return "";
    });
    
    console.log("PROPS", props.display_id)

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

    return { display, loading, display_controller_url };
  },
});
</script>