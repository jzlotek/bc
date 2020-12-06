<template>
  <h1 class="header" id="hOneHead">CryptoSim</h1>
  <p class="header" id="pOneHead">John Zlotek | Kevin Wu | Jainil Patel</p>
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
body {
  background: black;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.header {
  color: oldlace;
  font-family: Mshtakan;
}

#hOneHead {
  font-size: 35px;
}

#pOneHead {
  font-size: 20px;
}
</style>
