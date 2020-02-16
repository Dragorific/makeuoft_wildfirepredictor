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
            :position = "m.position"
            :clickable="true"
            :draggable="true"
            @click="Place=m.name"
        />
        </GmapMap>
        <h1>{{Place}}</h1>
    </div>

</template>

<script>
    export default {
        data(){
            return{
                Place: "nil",
                Markers: [{
                    name: "Main",
                    position: {lat:55.585901, lng:-105.750596},
                }]
            }
        }
    }
</script>