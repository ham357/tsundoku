<template>
  <div id="app">
    <TheHeader />

    <div class="columns">
      <div class="column content">
        <nuxt />
      </div>
    </div>
  </div>
</template>

<script>
import TheHeader from '~/components/TheHeader';
export default {
    components: {
        TheHeader
    }
};
</script>

<style scoped>
  #app {
    min-height: 100vh;
    background: #FAFAFA;
  }
  .columns {
    min-height: 100vh;
  }
</style>
