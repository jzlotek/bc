<template>
  <div class="card" v-bind:ref=" blockNum ">
    <b>Bind these text values to their props.</b><br>
    <b>Convert these text values into input fields.</b><br>
    <b>On change in input fields, display the block as red.</b><br>
    <b>Show the <em>Mine!</em> only if the block has been tampered:</b> <br>
    <div>Block: #{{ blockNum }}: </div>
    <div>Data: {{ data }}</div>
    <div>Nonce: {{ nonce }}</div>
    <div>prevHash: {{ prevHash }}</div>
    <div>curHash: {{ curHash }}</div>
    <div v-if="tampered === true">
      <button @click="remine()" type="button">Mine!</button>
    </div>
  </div>
</template>

<script>
  import { Timeline } from "@/components/Timeline.vue";

  export default {
    name: "Card",
    props: {
      blockNum: String,
      data: String,
      nonce: Number, // Invalid prop: type check failed for prop
      prevHash: String,
      curHash: String,
      tampered: Boolean,
    },
    methods: {
      remine: function () {
        alert("Block #" + this.blockNum + " will be remined.\nremine() is WIP.");
        this.$refs[this.blockNum].classList += " cardGood";
        this.nextBlock();
      },
      mine: function () {
        alert("Block #" + this.blockNum + " will be mined.\nmine() is WIP.");
      },
      nextBlock: function () {
        console.log(Timeline);
      }
    }
  };
</script>

<style scoped lang="scss">
.card {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  border: 0.5rem solid red;
  padding: 1rem;
  border-radius: 0.5rem;
}

.cardGood {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  border: 0.5rem solid green;
  padding: 1rem;
  border-radius: 0.5rem;
}
</style>
