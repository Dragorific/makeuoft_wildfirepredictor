<template>
  <div class="small">
    <line-chart :chart-data="tempDataCollection"></line-chart>
    <line-chart :chart-data="humDataCollection"></line-chart>
    <line-chart :chart-data="lightDataCollection"></line-chart>
    <button @click="fillData()">Randomize</button>
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
        datacollectionOne: null
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
              backgroundColor: '#D2E459',
              //data: [this.getTempData(), this.getTempData()]
              data: [this.getRandomInt(), this.getRandomInt()]
            }
          ]
        }, 
        this.HumDataCollection = {
          labels: [this.getRandomInt(), this.getRandomInt()],
          datasets: [
            {
              label: 'Humidity',
              backgroundColor: '#D2E459',
              //data: [this.getHumData(), this.getHumData()]
              data: [this.getRandomInt(), this.getRandomInt()]
            }
          ]
        }, 
        this.lightDataCollection = {
          labels: [this.getRandomInt(), this.getRandomInt()],
          datasets: [
            {
              label: 'Light Strength',
              backgroundColor: '#D2E459',
              //data: [this.getLightData(), this.getLightData]
              data: [this.getRandomInt(), this.getRandomInt()]
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