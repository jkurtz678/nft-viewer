<template>
    <Card class="p-m-2" style="width: 150px;" @click="$emit('editDisplay')">
        <template #footer>
            <div class="p-text-bold">{{name}}</div>
            <div>{{code}}</div>
        </template>
    </Card>
</template>

<script lang="ts">
import { defineComponent} from "vue";
export default defineComponent({
    props: {
        name: String,
        code: String,
        media_url: String,
    },
});
</script>

<style>
</style>