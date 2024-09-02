<script lang="ts" setup>
import { reactive } from "vue";
import { Greet } from "@wails/go/main/App";
import { Button } from "@/components/shadcn/button";

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
});

function greet() {
  Greet(data.name).then((result) => {
    data.resultText = result;
  });
}
</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input
        id="name"
        v-model="data.name"
        autocomplete="off"
        class="input"
        type="text"
      />
      <Button class="btn" @click="greet">Greet</Button>
    </div>
  </main>
</template>
