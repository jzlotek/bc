<template>
    <b>Bind these text values to their props.</b><br>
    <b>Convert these text values into input fields.</b><br>
    <b>On change in input fields, display the block as red.</b><br>
    <b>Show the <em>Mine!</em> only if the block has been tampered:</b> <br>
    <div>Block: #{{ blockNum }}: </div>
    <div v-if="newData == data">Data: {{ data }}</div>
    <div v-else>Data: {{ newData }}</div>
    <div>Nonce: {{ nonce }}</div>
    <div>Parent: {{ parent }}</div>
    <div>Hash: {{ hash }}</div>
    <div>Enter Data: <input id="newData" v-model="newData" @keypress="emitFunction"></div>
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
    },
    methods: {
      emitFunction: function () {
        console.log("here");
        this.$emit('changed', document.getElementById("newData").value);
      },
    }
  };
</script>

<style scoped lang="scss">
</style>
