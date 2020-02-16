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
        <!--h1>{{Markers}}</h1-->
        <br/>
        <br/>
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
            setInterval(this.updateCharts, 2000)
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
                fetch(process.env.VUE_APP_ENDPOINT + "/api/getData-usa")
                .then(response => {
                    return response.json();
                })
                .then(json => {
                    this.Temp = json.Temp
                    this.Humidity = json.Humidity 
                    this.Light = json.Light 
                    console.log(json.markers); // eslint-disable-line no-console
                })
                .catch(err => {
                    console.log(err); // eslint-disable-line no-console
                });

                this.fillData()
            },

            fillData () {
                this.tempDataCollection = {
                    labels: [0, 100],
                    datasets: [
                        {
                        label: 'Temperature',
                        backgroundColor: 'red',
                        data: this.Temp, 
                        
                        }
                    ], 
                    scales: {
                        yAxes: [{
                            ticks: {
                                stepSize: 50,
                                maxTicksLimit: 3
                            }
                        }]
                    }
                }, 
                
                this.humDataCollection = {
                    labels: [0, 100],
                    datasets: [
                        {
                        label: 'Humidity',
                        backgroundColor: 'green',
                        data: this.Humidity
                        }
                    ]
                }, 
                this.lightDataCollection = {
                    labels: [0, 30],
                    datasets: [
                        {
                        label: 'Light Strength',
                        backgroundColor: 'blue',
                        data: this.Light
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