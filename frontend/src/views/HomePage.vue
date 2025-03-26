<template>
  <Header />
  <div class="content">
    <div id="left">
      <p>left</p>
    </div>

    <div id="middle">
      <TodoMiddleContent />
    </div>

    <div id="right">
      <p>right</p>
    </div>
  </div>
  <Footer />
</template>

<script setup lang="ts">
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import TodoMiddleContent from "../components/todoMiddleContent.vue";
</script>

<style scoped>
#left{
  grid-area: left;
  background-color: bisque;
}

#middle{
  grid-area: middle;
}

#right{
  grid-area: right;
  background-color: antiquewhite;
}

.content{
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  grid-template-rows: auto;
  grid-template-areas: "left middle right";
}
</style>
