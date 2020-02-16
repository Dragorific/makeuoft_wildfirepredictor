<template>
    <div>
        <GmapMap
        :center="{lat:55.585901, lng:-105.750596}"
        :zoom="3"
        map-type-id="terrain"
        style="width: 80vw; height: 60vh; margin-left: auto; margin-right: auto;"
        >
        <GmapMarker 
            v-on:click="updateCharts"
            :key="index"
            v-for="(m,index) in Markers"
            :position = getMarker(index)
            :clickable="true"
            :draggable="true"
        />
        </GmapMap>
        <h1>{{Markers}}</h1>

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

    </div>


</template>

<script>
    import LineChart from './LineChart.js'
    export default {
        components: {
            LineChart
        },
        mounted(){
            fetch(process.env.VUE_APP_ENDPOINT + "/api/get-markers")
            .then(response => {
                return response.json();
            })
            .then(json => {
                this.Markers = json.markers
                console.log(json.markers); // eslint-disable-line no-console
            })
            .catch(err => {
                console.log(err); // eslint-disable-line no-console
            });
            this.fillData()
        },
        data(){
            return{
                Place: "nil",
                Markers: [["Main", "70.712891","37.09024"]], 
                tempDataCollection: null,
                humDataCollection: null,
                lightDataCollection: null
            }
        },
        methods: {
            getMarker(i){
                return {lat: parseFloat(this.Markers[i][1]), lng: parseFloat(this.Markers[i][2])}
            },
            updateCharts(){
                fetch(process.env.VUE_APP_ENDPOINT + "/api/get-markers")
                .then(response => {
                    return response.json();
                })
                .then(json => {
                    this.Markers = json.markers
                    console.log(json.markers); // eslint-disable-line no-console
                })
                .catch(err => {
                    console.log(err); // eslint-disable-line no-console
                });
            },

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
                getRandomInt () {
                    return Math.floor(Math.random() * (50 - 5 + 1)) + 5
                }
        }
    }
</script>

<style scoped>
  @import '../assets/css/charts.scss';
</style>