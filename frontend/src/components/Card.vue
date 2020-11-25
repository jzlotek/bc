<template>
  <div class="card">
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
    },
    mine: function () {
      alert("Block #" + this.blockNum + " will be mined.\nmine() is WIP.");
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
</style>
