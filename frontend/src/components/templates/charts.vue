<template>
  <div class="small">
    <div class="columns">
      <div class="column">
        <line-chart class="graph-display" :chart-data="tempDataCollection"></line-chart>
      </div>
      <div class="column">
        <line-chart class="graph-display" :chart-data="humDataCollection"></line-chart>
      </div>
      <div class="column">
        <line-chart class="graph-display" :chart-data="lightDataCollection"></line-chart>
      </div>
    </div>
    <!--button @click="fillData()">Randomize</button-->
  </div>
</template>

<script>
  import LineChart from './LineChart.js'

  export default {
    components: {
      LineChart
    },
    data () {
      return {
        tempDataCollection: null,
        humDataCollection: null,
        lightDataCollection: null
      }
    },
    mounted () {
      
      this.fillData()
    },
    methods: {
      fillData () {
        this.tempDataCollection = {
          labels: [this.getRandomInt(), this.getRandomInt()],
          datasets: [
            {
              label: 'Temperature',
              backgroundColor: 'red',
              //data: [this.getTempData(), this.getTempData()]
              data: [this.getRandomInt(), this.getRandomInt(), this.getRandomInt(), this.getRandomInt(), this.getRandomInt()]
            }
          ]
        }, 
        this.humDataCollection = {
          labels: [this.getRandomInt(), this.getRandomInt()],
          datasets: [
            {
              label: 'Humidity',
              backgroundColor: 'green',
              //data: [this.getHumData(), this.getHumData()]
              data: [this.getRandomInt(), this.getRandomInt(), this.getRandomInt(), this.getRandomInt(), this.getRandomInt()]
            }
          ]
        }, 
        this.lightDataCollection = {
          labels: [this.getRandomInt(), this.getRandomInt()],
          datasets: [
            {
              label: 'Light Strength',
              backgroundColor: 'blue',
              //data: [this.getLightData(), this.getLightData]
              data: [this.getRandomInt(), this.getRandomInt(), this.getRandomInt(), this.getRandomInt(), this.getRandomInt()]
            }
          ]
        }
      },
      //make a function here that grabs the data from the requests
      /*
      getTempData(){
        return
      },
      getHumData(){
        return
      },
      getLightData(){
        return
      }
      */
      getRandomInt () {
        return Math.floor(Math.random() * (50 - 5 + 1)) + 5
      }
    }
  }
</script>

<style scoped>
  @import '../../assets/css/charts.scss';
</style>