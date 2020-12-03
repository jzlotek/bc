<template>
  <button class="mineAll" @click="mineall" type="button">Mine All</button>
  <div v-for="card in cards" :key="card.blockNum">
    <div class="card" v-bind:ref=" card.blockNum ">
      <Card
        :blockNum="card.blockNum"
        :data="card.data"
        :nonce="card.nonce"
        :parent="card.parent"
        :hash="card.hash"
        :tampered="card.tampered"
        :time="card.time"
        v-on:changed="checkTampered($event, card)"
      />
      <button v-if="card.tampered==true" @click="mine(card)" type="button">Mine</button>
    </div>
  </div>
</template>

<script>
  import Card from "./Card.vue";
  import axios from "axios";

  export default {
    name: "Timeline",
    components: { Card },
    data: () => ({
              cards: [
                {
                  blockId: 0,
                  blockNum: "0",
                  dataHeld:"Genesis block.",
                  nonce: 0,
                  parent: "0000000000000000000000000000000000000000000000000000000000000000",
                  hash: "0000000000000000000000000000000000000000000000000000000000000000",
                  tampered: true,
                  time: 0,
                  mined: false
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
      },
      mineall: function() {
        var i;
        for (i = 0; i < this.cards.length; i++) {
          var card = this.cards[i];
          if (card.tampered == false) {
            continue;
          }
          if (card.mined == true) {
            this.$refs[card.blockNum].classList.add("cardGood");
            card.tampered = false;
            card.dataMined = card.data;

            this.api(card, true);
          } else {
            this.$refs[card.blockNum].classList.add("cardGood");
            card.tampered = false;
            card.dataMined = card.data;

            this.api(card, false);
          }
        }
      },
      mine: function (card) {
        var i;
        for (i = 0; i < this.cards.length; i++) {
          if (card.mined == true) {
            this.$refs[card.blockNum].classList.add("cardGood");
            card.tampered = false;
            card.dataMined = card.data;

            this.api(card, true);
            return 0;
          }
        }

        this.$refs[card.blockNum].classList.add("cardGood");
        card.tampered = false;
        card.dataMined = card.data;

        this.api(card, false);
      },
      api: function (card, remine) {
        const request = JSON.stringify({
            "start": "000",
            "block": {
              "previous_hash": card.parent,
              "id": card.blockId,
              "nonce": card.nonce,
              "data": card.data
            }
          });
        axios.post("https://bc.dulcimer.live/go/single", request)
                .then(response => {
                  card.hash = response.data.data.hash;
                  card.nonce = response.data.data.nonce;
                  card.time = response.data.data.time;
                  card.mined = true;
                  if (remine === false) {
                    this.cards.push({
                      blockId: card.blockId + 1,
                      blockNum: String(parseInt(card.blockNum) + 1),
                      data: "some",
                      nonce: 0,
                      parent: card.hash,
                      hash: "0000000000000000000000000000000000000000000000000000000000000000",
                      tampered: true,
                      time: 0,
                      mined: false
                    })
                  }
                })
                .then(data => this.postId = data.id)
                .catch(err => console.error(err));
      }
    }
  };
</script>

<style scoped lang="scss">
  .card {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: left;
    color: #2c3e50;
    margin-top: 60px;
    border: 0.5rem solid darkred;
    padding: 1rem;
    border-radius: 0.5rem;
    background: #ffd3cf;
    display: inline-grid;
    width: 700px;
  }

  .cardGood {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: left;
    color: #2c3e50;
    margin-top: 60px;
    border: 0.5rem solid darkgreen;
    padding: 1rem;
    border-radius: 0.5rem;
    background: #c4ffd1;
    display: inline-grid;
    width: 700px;
  }

  button {
    background: white;
    font-size: 20px;
    border-radius: 10px;
    font-family: Chalkboard;
    width: 150px;
  }

</style>