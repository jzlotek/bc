<template>
    <div><p>Block:</p> #{{ blockNum }}</div>
    <div v-if="newData == data"><p>Data:</p> {{ data }}</div>
    <div v-else><p>Data:</p> {{ newData }}</div>
    <div><p>Nonce:</p> {{ nonce }}</div>
    <div><p>Parent:</p> {{ parent }}</div>
    <div><p>Hash:</p> {{ hash }}</div>
    <div><p>Enter Data: </p> <input id="newData" v-model="newData" @keypress="emitFunction"></div>
    <div><p>Time:</p> {{ time }}</div>
</template>

<script>
  export default {
    name: "Card",
    props: {
      blockNum: String,
      data: String,
      nonce: Number, // Invalid prop: type check failed for prop
      parent: String,
      hash: String,
      tampered: Boolean,
        time: String,
    },
    methods: {
      emitFunction: function () {
        this.$emit('changed', document.getElementById("newData").value);
      },
    }
  };
</script>

<style scoped lang="scss">
    div {
        padding: 2px;
        font-family: Chalkboard;
    }

    p {
        display: inline;
        font-weight: bolder;
    }

    input {
        background: seashell;
    }
</style>
