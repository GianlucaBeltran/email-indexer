<script setup>
  const props = defineProps({
    data: Object,
    search: String,
  });
  const emit = defineEmits(["emailClicked"]);

  function emailClicked(emailId) {
    emit("emailClicked", emailId);
  }
</script>

<template>
  <div
    @click="emailClicked(props.data._id)"
    v-if="props.data.highlight"
    class="grid grid-cols-3 w-full h-auto border-b border-slate-200 items-center p-2 hover:shadow-[#ced4dc] hover:shadow-sm cursor-pointer"
  >
    <p class="truncate">
      <span
        v-if="props.data.highlight.From"
        v-html="props.data.highlight.From[0]"
      ></span>
      <span v-else>{{ props.data._source.From }}</span>
    </p>
    <p class="truncate" v-if="props.data._source.To.includes(',')">
      {{ props.data._source.To.split(", ")[0] }}
      +{{ props.data._source.To.split(", ").length - 1 }} more
    </p>
    <p class="truncate" v-else-if="props.data._source.To === ''">No Reciever</p>
    <p class="truncate" v-else>
      <span
        v-if="props.data.highlight.To"
        v-html="props.data.highlight.To[0]"
      ></span>
      <span v-else>{{ props.data._source.To }}</span>
    </p>
    <p class="truncate">
      <span
        v-if="props.data.highlight.Date"
        v-html="props.data.highlight.Date[0]"
      ></span>
      <span v-else>{{ props.data._source.Date }}</span>
    </p>
  </div>
  <div
    @click="emailClicked(props.data._id)"
    v-else
    class="grid grid-cols-3 w-full h-auto border-b border-slate-200 items-center p-2 hover:shadow-[#ced4dc] hover:shadow-sm cursor-pointer"
  >
    <p class="truncate">{{ props.data._source.From }}</p>
    <p class="truncate" v-if="props.data._source.To.includes(',')">
      {{ props.data._source.To.split(", ")[0] }}
      +{{ props.data._source.To.split(", ").length - 1 }} more
    </p>
    <p class="truncate" v-else-if="props.data._source.To === ''">No Reciever</p>
    <p class="truncate" v-else>{{ props.data._source.To }}</p>
    <p class="truncate">{{ props.data._source.Date }}</p>
  </div>
</template>
