<template>
  <div>
    <!-- This card should be hard-coded at the start but it is editable -->
    <Card
      v-for="card in cards" :key="card.blockNum"
      :blockNum="card.blockNum"
      :data="card.data"
      :nonce="card.nonce"
      :prevHash="card.prevHash"
      :curHash="card.curHash"
      :tampered="card.tampered"
    />
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
};
</script>

<style scoped lang="scss">
</style>
