<template>
  <div>
    <Card
      :style="selected ? 'border-style: solid;' : ''"
    >
      <template #content>
        <div class="p-grid">
          <div
            class="p-col"
            style="margin: auto; text-align: left; padding-left: 40px"
          >{{token.name}}</div>
          <div class="p-col">
            <img
              :src="token.image_thumbnail_url"
              style="height: 100px;"
            />
          </div>
        </div>
      </template>
    </Card>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Token } from "../../types/types";
export default defineComponent({
  props: {
    token: Object as () => Token,
    selected: Boolean,
  },
});
</script>

<style scoped>
.p-card {
  width: 100%;
  margin-bottom: 10px;
  text-align: center;
  box-shadow: none;
}
.p-card::v-deep .p-card-body {
  padding: 0px;
}
.p-card::v-deep .p-card-footer {
  padding: 10px 0px 10px 0px;
}
.p-card::v-deep .p-card-content {
  padding-bottom: 0px;
}
</style>