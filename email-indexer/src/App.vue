<script setup>
  import { ref, watch } from "vue";
  import Search from "./components/Search.vue";
  import Inbox from "./components/Inbox.vue";
  import Email from "./components/Email.vue";

  const search = ref("");
  const testData = ref("");
  const emailId = ref("");
  const email = ref("");
  const showAllData = ref(false);
  const currentPage = ref(0);
  const currentMaxPerPage = ref(50);

  function test() {
    console.log(showAllData.value, "test");
  }

  async function getInbox() {
    const response = await fetch(
      "http://localhost:4080/api/enron_mail/_search",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          authorization: "Basic " + btoa("admin:Complexpass#123"),
        },
        body: JSON.stringify({
          search_type: "alldocuments",
          sort_fields: ["-Date"],
          from: currentPage.value * currentMaxPerPage.value,
          max_results: currentMaxPerPage.value,
          _source: ["From", "To", "Date"],
        }),
      }
    );

    const data = await response.json();
    console.log(data);
    testData.value = data;
  }
  getInbox();

  async function getEmail() {
    const response = await fetch(
      "http://localhost:4080/api/enron_mail/_search",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          authorization: "Basic " + btoa("admin:Complexpass#123"),
        },
        body: JSON.stringify({
          search_type: "term",
          query: {
            term: emailId.value,
            field: "_id",
          },
          _source: [],
        }),
      }
    );

    const data = await response.json();
    email.value = data.hits.hits[0]._source;
  }

  async function getSearch() {
    if (search.value === "") {
      getInbox();
      return;
    }
    const response = await fetch(
      "http://localhost:4080/api/enron_mail/_search",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          authorization: "Basic " + btoa("admin:Complexpass#123"),
        },
        body: JSON.stringify({
          search_type: "querystring",
          query: {
            term: search.value,
          },
          from: currentPage.value * currentMaxPerPage.value,
          max_results: currentMaxPerPage.value,
          _source: ["From", "To", "Date"],
          highlight: {
            pre_tags: ["<span class='bg-yellow-300'>"],
            post_tags: ["</span>"],
            fields: {
              From: {
                pre_tags: [],
                post_tags: [],
              },
              To: {
                pre_tags: [],
                post_tags: [],
              },
              Date: {
                pre_tags: [],
                post_tags: [],
              },
            },
          },
        }),
      }
    );
    if (response.status === 400) {
      testData.value = "";
      return;
    }
    const data = await response.json();
    console.log(data);
    testData.value = data;
  }

  async function getFirstPage() {
    if (currentPage.value === 0) {
      return;
    }
    currentPage.value = 0;
    testData.value = "";
    emailId.value = "";
    email.value = "";
    getSearch();
  }

  async function getPrevPage() {
    if (currentPage.value === 0) {
      return;
    }
    currentPage.value = currentPage.value - 1;
    testData.value = "";
    emailId.value = "";
    email.value = "";
    getSearch();
  }
  async function getNextPage() {
    if (
      currentPage.value ===
      Math.ceil(testData.value.hits.total.value / currentMaxPerPage.value) - 1
    ) {
      return;
    }
    currentPage.value = currentPage.value + 1;
    testData.value = "";
    emailId.value = "";
    email.value = "";
    getSearch();
  }

  watch(emailId, getEmail);
  watch(search, getSearch);
  // watch(showAllData, test);
</script>

<template>
  <div class="flex h-screen flex-col px-20 py-5">
    <Search @newSearch="(newSearch) => (search = newSearch)" />
    <div class="flex justify-between" v-if="testData">
      <div class="flex">
        <svg
          @click="getFirstPage"
          class="cursor-pointer"
          xmlns="http://www.w3.org/2000/svg"
          height="25"
          viewBox="0 -960 960 960"
          width="25"
        >
          <path
            d="M248-251v-462h54v462h-54Zm436-4L458-481l226-226 38 38-188 188 188 188-38 38Z"
          />
        </svg>
        <svg
          @click="getPrevPage"
          class="cursor-pointer"
          xmlns="http://www.w3.org/2000/svg"
          height="25"
          viewBox="0 -960 960 960"
          width="25"
        >
          <path d="M562-257 338-481l224-224 38 38-186 186 186 186-38 38Z" />
        </svg>
        <svg
          @click="getNextPage"
          class="cursor-pointer"
          xmlns="http://www.w3.org/2000/svg"
          height="25"
          viewBox="0 -960 960 960"
          width="25"
        >
          <path d="m376-257-38-38 186-186-186-186 38-38 224 224-224 224Z" />
        </svg>
        <p class="self-center ml-1">
          {{ currentMaxPerPage }} results per page ({{
            testData.hits.total.value
          }}
          total results). Current page:
          {{ currentPage + 1 }} of
          {{ Math.ceil(testData.hits.total.value / currentMaxPerPage) }}
        </p>
      </div>
      <div v-if="email">
        <input type="checkbox" id="checkbox" v-model="showAllData" />
        <label for="checkbox">Show all data</label>
      </div>
    </div>
    <div class="flex flex-col gap-4 md:flex-row flex-grow h-4/5">
      <Inbox
        :data="testData"
        @emailClicked="(clickedEmailId) => (emailId = clickedEmailId)"
        :search="search"
      />
      <Email
        :data="email"
        :show-all-data="showAllData"
        :search="search"
        v-if="email"
      />
    </div>
  </div>
</template>
