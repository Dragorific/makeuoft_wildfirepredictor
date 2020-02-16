<template>
    <div>
        <GmapMap
        :center="{lat:55.585901, lng:-105.750596}"
        :zoom="3"
        map-type-id="terrain"
        style="width: 80vw; height: 60vh; margin-left: auto; margin-right: auto;"
        >
        <GmapMarker
            :key="index"
            v-for="(m,index) in Markers"
            :position = getMarker(index)
            :clickable="true"
            :draggable="true"
        />
        </GmapMap>
        <h1>{{Markers}}</h1>
    </div>

</template>

<script>
    export default {
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
        },
        data(){
            return{
                Place: "nil",
                Markers: [["Main", "70.712891","37.09024"]]
            }
        },
        methods: {
            getMarker(i){
                return {lat: parseFloat(this.Markers[i][1]), lng: parseFloat(this.Markers[i][2])}
            }
        },
    }
</script>