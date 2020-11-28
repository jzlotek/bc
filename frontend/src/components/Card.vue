<template>
  <Timeline ></Timeline>
    <b>Bind these text values to their props.</b><br>
    <b>Convert these text values into input fields.</b><br>
    <b>On change in input fields, display the block as red.</b><br>
    <b>Show the <em>Mine!</em> only if the block has been tampered:</b> <br>
    <div>Block: #{{ blockNum }}: </div>
    <div>Data: {{ data }}</div>
    <div>Nonce: {{ nonce }}</div>
    <div>prevHash: {{ prevHash }}</div>
    <div>curHash: {{ curHash }}</div>
</template>

<script>
  import Timeline from "./Timeline.vue";
  export default {
    name: "Card",
    components: {
      'timeline': { Timeline },
    },
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
        console.log('d')
        this.$refs[this.blockNum].classList += " cardGood";
        //this.$emit('mine', this.$refs[this.blockNum]);
      },
    }
  };
</script>

<style scoped lang="scss">
</style>
