<template>
  <div v-for="card in cards" :key="card.blockNum">
    <div class="card" v-bind:ref=" card.blockNum ">
      <Card
        :blockNum="card.blockNum"
        :data="card.data"
        :nonce="card.nonce"
        :parent="card.parent"
        :hash="card.hash"
        :tampered="card.tampered"
        :func="this.remine"
        v-on:changed="checkTampered($event, card)"
      />
      <button v-if="card.tampered==true" @click="remine(card)" type="button">Mine!</button>
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
          dataHeld:"This is the starter block.",
          nonce: "-1",
          parent: "0",
          hash: "calculatedHashOfThisBlock",
          tampered: true,
          dataMined: ""
       }
      ]
    }
  ),
  methods: {
    checkTampered: function(data, card) {
      var i;
      if (card.dataMined != data && card.tampered == false) {
        for (i = card.blockNum; i < this.cards.length; i++) {
          this.$refs[i].classList.remove("cardGood");
          this.cards[i].tampered = true;
        }
      }
      else {
        console.log("d")
      }
    },
    remine: function (card) {
      this.$refs[card.blockNum].classList.add("cardGood");
      card.tampered = false;
      card.dataMined = card.data;

      this.cards.push({
        blockNum: String(parseInt(card.blockNum) + 1),
        data: "some",
        nonce: "-1",
        parent: String(this.cards[card.blockNum].hash),
        hash: "new",
        tampered: true,
        dataMined: ""
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