<template>
  <div>
    <!-- This card should be hard-coded at the start but it is editable -->
    <Card
      blockNum="0"
      data="This is the starter block."
      nonce="-1"
      prevHash="0"
      curHash="calculatedHashOfThisBlock"
      tampered=false
    />

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
          blockNum: "1",
          data:"HERE IS SOME DATA",
          nonce: "42069",
          prevHash: "prevHash85F5E10431F69BC2A14046A13AABAEFC660103B6DE7A84F75C4B96181D03F0B5",
          curHash: "curHashC49B876C05FD376922883CEAF5B7E7871AE4AE9F24E79958F3FC70EFE114627D",
          tampered: true
        }
      ]
    }
  ),
};
</script>

<style scoped lang="scss">
</style>
