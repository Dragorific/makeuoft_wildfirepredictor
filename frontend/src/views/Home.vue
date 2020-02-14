<template>
  <div class="home">
    <navbar/>
    <div class="front-page">
      <div class="paralax">
        <div class = "overlay">
        </div>
        <div class="project-title">
          <h1>Wildfire Predictor Model 2.0</h1>
        </div>
      </div>
    </div>
    <div class="project-content">
      <div class="intro">
        <div class="container-fluid">
          <div class="row">
            <div class="col-12 col-md-6">
              <div class="intro-img"><p>image of something?</p></div>
            </div>
            <div class="col-12 col-md-6">
              <div class="intro-desc"><p>description of shit</p></div>
            </div>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-12">
          <div class="info"><p>idk what goes here</p></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import navbar from '@/components/navbar.vue'

export default {
  name: 'home',
  components: {
    navbar
  }
}
</script>

<style scoped>
  @import '../assets/css/HomePage.css'
</style>
