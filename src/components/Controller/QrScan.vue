<template>
  <div style="text-align: right; background-color: #000000; width: 100%; height: 100%; position: fixed;">
    <Button
      label="Cancel scanning"
      icon="pi pi-ban"
      class="p-m-3"
      @click="store.commit('setCameraScanMode', false)"
    />
    <loading
      :active="camera_loading"
      :opacity="0"
      color="#FFFFFF"
    ></loading>
    <div>
      <qrcode-stream
        style="width: 100%;"
        @decode="onDecode"
        @init="onInit"
      ></qrcode-stream>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { useStore } from "vuex";
import { QrcodeStream } from "vue-qrcode-reader";
import { addAccountToDisplay } from "@/api/display";
import Loading from "vue-loading-overlay";

export default defineComponent({
  components: { QrcodeStream, Loading },
  props: { account_id: String },
  setup(props) {
    const store = useStore();
    const camera_loading = ref(false);
    const onDecode = async (qr_code_link: string) => {
      const display_id_index = qr_code_link.indexOf("display_id=");
      if (display_id_index > -1) {
        const display_id = qr_code_link.slice(display_id_index + 11);
        await addAccountToDisplay(display_id, props.account_id || "");
      }
      store.commit("setCameraScanMode", false);
    };

    const onInit = async (promise: Promise<any>) => {
      camera_loading.value = true;

      try {
        await promise;
        // successfully initialized
      } catch (error: any) {
        if (error.name === "NotAllowedError") {
          // user denied camera access permisson
        } else if (error.name === "NotFoundError") {
          // no suitable camera device installed
        } else if (error.name === "NotSupportedError") {
          // page is not served over HTTPS (or localhost)
        } else if (error.name === "NotReadableError") {
          // maybe camera is already in use
        } else if (error.name === "OverconstrainedError") {
          // did you requested the front camera although there is none?
        } else if (error.name === "StreamApiNotSupportedError") {
          // browser seems to be lacking features
        }
      } finally {
        camera_loading.value = false;
      }
    };

    return { store, onDecode, onInit, camera_loading };
  },
});
</script>
