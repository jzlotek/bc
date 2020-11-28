<template>
  <div v-for="card in cards" :key="card.blockNum">
    <div class="card" v-bind:ref=" card.blockNum ">
      <Card
        :blockNum="card.blockNum"
        :data="card.data"
        :nonce="card.nonce"
        :prevHash="card.prevHash"
        :curHash="card.curHash"
        :tampered="card.tampered"
        :func="this.remine"
      />
      <button @click="remine(card.blockNum)" type="button">Mine!</button>
    </div>
  </div>
</template>

<script>
import Card from "./Card.vue";

export default {
  name: "Timeline",
  components: { Card },
  data: () => ({
    cards: [
        {
          blockNum: "0",
          data:"This is the starter block.",
          nonce: "-1",
          prevHash: "0",
          curHash: "calculatedHashOfThisBlock",
          tampered: true
       }
      ]
    }
  ),
  methods: {
    nextBlock: function() {
      this.$click('mine', console.log('s'));
    },
    remine: function (block) {
      console.log(block);
      this.$refs[block].classList += " cardGood";
      this.cards.push({
        blockNum: String(parseInt(block) + 1),
        data: "some",
        nonce: "-1",
        prevHash: "0",
        curHash: "new",
        tampered: true,
      })
    },
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