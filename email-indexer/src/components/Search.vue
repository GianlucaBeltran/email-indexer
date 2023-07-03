<script setup>
  import { ref, watch } from "vue";
  const emit = defineEmits(["newSearch"]);

  const search = ref("");

  function deleteSearch() {
    search.value = "";
  }

  watch(search, () => emit("newSearch", search.value));
</script>

<template>
  <label class="relative block md:w-1/3 mb-8 w-full">
    <span class="sr-only">Search</span>
    <span class="absolute inset-y-0 left-0 flex items-center pl-2">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        height="25"
        viewBox="0 -960 960 960"
        width="25"
      >
        <path
          d="m779.833-165.833-251.766-252q-29.471 27.051-69.101 40.942Q419.335-363 382.158-363q-92.061 0-155.193-63.163-63.132-63.163-63.132-154.5t62.708-154.67q62.708-63.334 154.348-63.334 91.641 0 155.459 63.22 63.819 63.22 63.819 154.498 0 40.616-15.334 79.949-15.333 39.333-40.166 66.89l252 250.777-16.834 17.5ZM382-386.167q81.889 0 138.445-55.931Q577-498.029 577-580.931t-56.555-138.736Q463.889-775.5 381.583-775.5q-82.305 0-138.444 55.931Q187-663.637 187-580.735T243.208-442Q299.417-386.167 382-386.167Z"
        />
      </svg>
    </span>
    <input
      class="placeholder:italic placeholder:text-slate-400 block bg-white w-full border border-slate-300 rounded-md py-2 pl-9 pr-3 shadow-sm focus:outline-none focus:border-sky-500 focus:ring-sky-500 focus:ring-1 md:text-sm"
      placeholder="Search mail"
      type="text"
      name="search"
      v-model="search"
    />
    <button
      class="absolute inset-y-1 right-0 flex items-center"
      v-if="search"
      @click="deleteSearch"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        height="25"
        viewBox="0 -960 960 960"
        width="25"
      >
        <path
          d="M255.167-238.333 238.5-255.167 463.167-480 238.5-704.833l16.667-16.834L480-496.833l224.833-224.834 16.667 16.834L496.833-480 721.5-255.167l-16.667 16.834L480-463.167 255.167-238.333Z"
        />
      </svg>
    </button>
  </label>
</template>
