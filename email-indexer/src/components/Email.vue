<script setup>
  const props = defineProps({
    data: Object,
    showAllData: Boolean,
    search: String,
  });
  //iterate over data and replace all instances of search with a span with a class of highlight
  function highlightSearch(data) {
    const regex = new RegExp(props.search, "gi");
    return data.replace(
      regex,
      (match) => `<span class="bg-yellow-300">${match}</span>`
    );
  }
</script>

<template>
  <div
    class="md:w-1/2 border border-solid border-slate-300 w-full z-10 overflow-y-auto"
  >
    <div class="flex mb-1">
      <h1 class="text-xl font-bold">Message-ID:&nbsp;</h1>
      <h1 class="text-xl">
        <span v-html="highlightSearch(props.data['Message-ID'])"></span>
      </h1>
    </div>
    <div class="flex mb-1">
      <h1 class="text-xl font-bold">From:&nbsp;</h1>
      <h1 class="text-xl">
        <span v-html="highlightSearch(props.data.From)"></span>
      </h1>
    </div>
    <div class="flex mb-1">
      <h1 class="text-xl font-bold">To:&nbsp;</h1>
      <h1 class="text-xl">
        <span v-html="highlightSearch(props.data.To)"></span>
      </h1>
      <h1 class="text-xl" v-if="!props.data.To">No Reciever</h1>
    </div>
    <div class="flex mb-1">
      <h1 class="text-xl font-bold">Date:&nbsp;</h1>
      <h1 class="text-xl">
        <span v-html="highlightSearch(props.data.Date)"></span>
      </h1>
    </div>
    <div class="flex mb-1">
      <h1 class="text-xl font-bold">Subject:&nbsp;</h1>
      <h1 class="text-xl">
        <span v-html="highlightSearch(props.data.Subject)"></span>
      </h1>
      <h1 class="text-xl" v-if="!props.data.Subject">No Subject</h1>
    </div>
    <div class="flex mb-1" v-if="showAllData">
      <h1 class="text-xl font-bold">Mime-Version:&nbsp;</h1>
      <h1 class="text-xl">{{ props.data["Mime-Version"] }}</h1>
    </div>
    <div class="flex mb-1" v-if="showAllData">
      <h1 class="text-xl font-bold">Content-Type:&nbsp;</h1>
      <h1 class="text-xl">{{ props.data["Content-Type"] }}</h1>
    </div>
    <div class="flex mb-1" v-if="showAllData">
      <h1 class="text-xl font-bold">Content-Transfer-Encoding:&nbsp;</h1>
      <h1 class="text-xl">{{ props.data["Content-Transfer-Encoding"] }}</h1>
    </div>
    <div class="flex mb-1" v-if="showAllData">
      <h1 class="text-xl font-bold">X-From:&nbsp;</h1>
      <h1 class="text-xl">{{ props.data["X-From"] }}</h1>
      <h1 class="text-xl" v-if="!props.data['X-From']">No Field</h1>
    </div>
    <div class="flex mb-1" v-if="showAllData">
      <h1 class="text-xl font-bold">X-To:&nbsp;</h1>
      <h1 class="text-xl">{{ props.data["X-To"] }}</h1>
      <h1 class="text-xl" v-if="!props.data['X-To']">No Field</h1>
    </div>
    <div class="flex mb-1" v-if="showAllData">
      <h1 class="text-xl font-bold">X-Cc:&nbsp;</h1>
      <h1 class="text-xl">{{ props.data["X-cc"] }}</h1>
      <h1 class="text-xl" v-if="!props.data['X-Cc']">No Field</h1>
    </div>
    <div class="flex mb-1" v-if="showAllData">
      <h1 class="text-xl font-bold">X-Bcc:&nbsp;</h1>
      <h1 class="text-xl">{{ props.data["X-Bcc"] }}</h1>
      <h1 class="text-xl" v-if="!props.data['X-Bcc']">No Field</h1>
    </div>
    <div class="flex mb-1" v-if="showAllData">
      <h1 class="text-xl font-bold">X-Folder:&nbsp;</h1>
      <h1 class="text-xl">{{ props.data["X-Folder"] }}</h1>
      <h1 class="text-xl" v-if="!props.data['X-Folder']">No Field</h1>
    </div>
    <div class="flex mb-1" v-if="showAllData">
      <h1 class="text-xl font-bold">X-Origin:&nbsp;</h1>
      <h1 class="text-xl">{{ props.data["X-Origin"] }}</h1>
      <h1 class="text-xl" v-if="!props.data['X-Origin']">No Field</h1>
    </div>
    <div class="flex mb-1" v-if="showAllData">
      <h1 class="text-xl font-bold">X-FileName:&nbsp;</h1>
      <h1 class="text-xl">{{ props.data["X-FileName"] }}</h1>
      <h1 class="text-xl" v-if="!props.data['X-FileName']">No Field</h1>
    </div>
    <br />
    <div class="flex mb-1">
      <h1 class="text-xl font-bold">Message:&nbsp;</h1>
      <h1 class="text-lg">
        <span v-html="highlightSearch(props.data.Message)"></span>
      </h1>
      <h1 class="text-base" v-if="!props.data['Message']">No Message</h1>
    </div>
  </div>
</template>
