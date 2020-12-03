<template>
    <div>Block: #{{ blockNum }}</div>
    <div v-if="newData == data">Data: {{ data }}</div>
    <div v-else>Data: {{ newData }}</div>
    <div>Nonce: {{ nonce }}</div>
    <div>Parent: {{ parent }}</div>
    <div>Hash: {{ hash }}</div>
    <div>Enter Data: <input id="newData" v-model="newData" @keypress="emitFunction"></div>
    <div>Time: {{ time }}</div>
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
</style>
