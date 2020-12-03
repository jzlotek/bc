<template>
  <div>
    <Timeline />
  </div>
</template>

<script>
import Timeline from "./components/Timeline.vue";

export default {
  name: "App",
  components: {
    Timeline
  }
};
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
