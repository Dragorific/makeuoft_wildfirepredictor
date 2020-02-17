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
    <div class="body">
      <div class="carousel">
        <b-carousel
        :autoplay=false>
          <b-carousel-item v-for="(carousel, i) in carousels" :key="i">
              <section class="carouselSlides">
                  <div class="hero-body has-text-centered">
                      <h1>{{carousel.heading}}</h1>
                      <hr>
                      <div class="carouselText">
                        <p>{{carousel.text}}</p>
                      </div>
                  </div>
              </section>
          </b-carousel-item>           
        </b-carousel>
      </div>
      <div class="graph">
        <graph/>
      </div>

      <div>
        <charts/>
      </div>

    </div>
  </div>
</template>

<script>
import navbar from '@/components/navbar.vue'
import graph from '@/components/graph.vue'

export default {
  components: {
    navbar,
    graph
  },
  data(){
        return {
            carousels: [
                { heading: 'Scope and Problem', text: "In the first stage of this project, the plan is to develop a basic infrastructure using a machine learning model in Python using Tensorflow and a Linear Regression training model, thoroughly trained to be able to process live, analog data, or data accessed from the web through a service such as “TheWeatherNetwork”, into an accurate percentage representation of the risk of a wildfire. The machine learning model will be able to process all four variables. Furthermore, devices will be assembled and programmed using a variety of ambient sensors that can read the four key variables, for the use of remote deployment in areas where there is no meteorological coverage. A web interface will be built on Vue.js to display the processed information, and this web app will use Golang as the main programming environment for the backend. The web app will be built using Kubernetes, more specifically Docker, to ensure modularity between components and reliable communication between the front-end and back-end. The front end will be able to create graphical representations of the data as well as be able to display the data geographically in a visually aesthetic manner."},
                { heading: 'Stage 1 | Baseline Detection', text: "In the first stage of this project, the plan is to develop a basic infrastructure using a machine learning model in Python using Tensorflow and a Linear Regression training model, thoroughly trained to be able to process live, analog data, or data accessed from the web through a service such as “TheWeatherNetwork”, into an accurate percentage representation of the risk of a wildfire. The machine learning model will be able to process all four variables. Furthermore, devices will be assembled and programmed using a variety of ambient sensors that can read the four key variables, for the use of remote deployment in areas where there is no meteorological coverage. A web interface will be built on Vue.js to display the processed information, and this web app will use Golang as the main programming environment for the backend. The web app will be built using Kubernetes, more specifically Docker, to ensure modularity between components and reliable communication between the front-end and back-end. The front end will be able to create graphical representations of the data as well as be able to display the data geographically in a visually aesthetic manner."},
                { heading: 'Stage 2 | Taking Action', text: "The next stage, once the foundation is laid, is to add features which will allow users to interact with the web app for live updates, and create a proof-of-concept autonomous solution that will serve as the baseline for developing more effective means of taking autonomous action towards wildfires. The features added in the next stage will take the appearance of email or cell subscriptions that will alert the subscribed user of any dangerous changes in the live data. The proof-of-concept autonomous solution will take the appearance of autonomous drones programmed to go to the site of the wildfire. This will serve as a baseline for drones in the future to be able to deploy immediately in swarms and control the spread of wildfire through fire fighting methods currently used by larger cargo planes during larger wildfires."},
                { heading: 'Structure', text: ""},
            ]
        }
    }
}
</script>

<style scoped>
  @import '../assets/css/HomePage.scss';
</style>
