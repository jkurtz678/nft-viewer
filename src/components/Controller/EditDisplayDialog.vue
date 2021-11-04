<template>
  <Dialog
    header="Header"
    :visible="display_id != null"
    :modal="true"
    @update:visible="$emit('update:display_id', null)"
    style="width: 100%; max-width: 700px;"
  >
    <template #header>New Display</template>
    <div
      class="p-grid"
      style="flex-wrap: nowrap"
    >
      <div class="p-col">
        <span
          class="p-float-label"
          style="margin-top: 20px;"
        >
          <InputText
            id="name"
            v-model="display.entity.name"
            type="text"
            style="margin: 0 auto; max-width: 150px;"
          />
          <label for="name">Name</label>
        </span>
        <Button
          class="p-button-sm p-mt-2"
          label="Open in browser"
          @click="openDisplayInBrowser"
          style="display: block;"
        ></Button>
        <div
          class="p-mt-2"
          style="display: flex; align-items: center;"
        >
          <InputSwitch
            v-model="display.entity.plaque_dark_mode"
            class="p-mr-2"
          ></InputSwitch>
          Plaque dark mode
        </div>
        <div
          class="p-mt-2"
          style="display: flex; align-items: center;"
        >
          <InputSwitch
            v-model="demo_token_playlist"
            class="p-mr-2"
          ></InputSwitch>
          Demo token playlist 
        </div>
        <Button
          class="p-button-danger p-button-sm p-mt-2"
          label="Forget display"
          @click="forgetDisplay"
          style="display: block;"
        ></Button>
      </div>
      <div class="p-col">
        <display-item
          :name="display.entity.name"
          :token_id="display.entity.token_id"
        ></display-item>
      </div>
    </div>
    <TokenList v-model:selected_token_id="display.entity.token_id"></TokenList>
    <template #footer>
      <Button
        label="Save"
        @click="handleSave"
      ></Button>
      <Button
        class="p-button-text"
        label="Cancel"
        @click="$emit('update:display_id', null)"
      ></Button>
    </template>
  </Dialog>
</template>

<script lang="ts">
import { defineComponent, ref, watch, computed } from "vue";
import { useStore } from "vuex";
import { Display, FirestoreDocument, Token } from "../../types/types";
import DisplayItem from "../Controller/DisplayItem.vue";
import { getDisplayByDisplayID, updateDisplay } from "../../api/display";
import TokenList from "../Controller/TokenList.vue";

export default defineComponent({
  components: { DisplayItem, TokenList },
  props: {
    createDisplay: {
      type: Function,
      required: true,
    },
    display_id: {
      type: String,
      required: false,
    },
  },
  setup(props, { emit }) {
    const display = ref<FirestoreDocument<Display>>({
      id: "",
      entity: { name: "" } as Display,
    });
    const store = useStore();

    const handleSave = async () => {
      //props.createDisplay(props.account_id, name.value);
      await updateDisplay(display.value);
      emit("update:display_id", null);
    };
    const forgetDisplay = async () => {
      display.value.entity.account_id = "";
      display.value.entity.token_id = "";
      display.value.entity.asset_contract_address = "";
      await updateDisplay(display.value);
      emit("update:display_id", null);
    };
    const openDisplayInBrowser = async () => {
      window.open(
        `${window.location.origin}${window.location.pathname}/#/display?display_id=${props.display_id}`
      );
    };

    watch(
      props,
      () => {
        if (props.display_id) {
          getDisplayByDisplayID(props.display_id).then((ret_display) => {
            display.value = ret_display;
          });
        } else {
          display.value = { id: "", entity: { name: "" } as Display };
        }
      },
      { immediate: true }
    );

    const token_id = computed(() => {
      return display.value.entity.token_id;
    });
    const demo_token_playlist = computed({
      get(): boolean {
        if(!display.value.entity.playlist_tokens){
          return false;
        }
        return display.value.entity.playlist_tokens.length > 0 
      },
      set(v): void{
        if(v) {
          display.value.entity.playlist_tokens = store.getters.demo_tokens.map((t: Token) => ({asset_contract_address: t?.asset_contract?.address, token_id: t.token_id}))
        } else {
          display.value.entity.playlist_tokens = [];
        }
      }
    })

    watch(token_id, (v) => {
      if (!v) {
        display.value.entity.asset_contract_address = "";
        return;
      }
      const token = store.getters.token(v);
      // token may be null the display is set to the token the user does not have
      if(!token) {
        console.log("TOKEN NOT FOUND FOR token_id = ", v)
        return
      }
      display.value.entity.asset_contract_address =
        token.asset_contract.address;
    });

    return {
      handleSave,
      forgetDisplay,
      display,
      openDisplayInBrowser,
      demo_token_playlist
    };
  },
});
</script>

<style>
</style>