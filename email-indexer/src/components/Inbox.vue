<script setup>
  import InboxItem from "./InboxItem.vue";

  const props = defineProps({
    data: Object,
    search: String,
  });

  const emit = defineEmits(["emailClicked"]);
</script>

<template>
  <div
    v-if="props.data"
    class="md:w-1/2 border border-solid border-slate-300 w-full z-10 overflow-y-auto"
  >
    <div
      class="grid grid-cols-3 w-full h-auto border-b border-slate-200 items-center p-2 font-bold bg-slate-200 sticky top-0"
    >
      <p>From</p>
      <p>To</p>
      <p>Date</p>
    </div>
    <InboxItem
      :search="search"
      @emailClicked="(clickedEmailId) => emit('emailClicked', clickedEmailId)"
      v-for="mail in props.data.hits.hits"
      :key="mail._id"
      :data="mail"
    />
  </div>
  <h1 v-else>Fetching inbox...</h1>
</template>

<style></style>
