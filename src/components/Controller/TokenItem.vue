<template>
  <div>
    <Card style="width: 150px; padding: 0px;">
      <template #content>
        <img :src="token.image_thumbnail_url" />
      </template>
      <template #footer>
          {{token.name}}
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
  },
});
</script>

<style scoped>
.p-card::v-deep .p-card-body {
  padding: 0px;
}
.p-card::v-deep .p-card-footer {
  padding: 10px 0px 10px 0px;
}
.p-card::v-deep .p-card-content{
  padding-bottom: 0px;
}
</style>