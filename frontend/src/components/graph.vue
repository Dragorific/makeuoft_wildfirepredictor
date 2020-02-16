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
            v-for="(m, index) in markers"
            :position="m.position"
            :clickable="true"
            :draggable="true"
            @click="center=m.position"
        />
        </GmapMap>
    </div>

</template>

<script>
    export default {
        
    }
</script>