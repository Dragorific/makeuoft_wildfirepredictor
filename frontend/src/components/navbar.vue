<template>
    <div class="navcontainer">
        <div style="z-index:7;">
            <b-navbar spaced fixed-top type="is-dark">
                <template slot="brand">
                    <b-navbar-item tag="router-link" :to="{ path: '/' }">
                        <img
                            src="https://raw.githubusercontent.com/buefy/buefy/dev/static/img/buefy-logo.png"
                            alt="Lightweight UI components for Vue.js based on Bulma"
                        >
                    </b-navbar-item>
                </template>
                <template slot="start">
                    <b-navbar-item href="/">
                        <router-link to="/">Home</router-link>
                    </b-navbar-item>
                    <b-navbar-item href="/">
                        <router-link to="../about">About</router-link>
                    </b-navbar-item>
                </template>
            </b-navbar>
        </div>
    </div>
</template>

<!--<template>
    <div class="navcontainer">
        <div style="z-index:7;">
            <nav role="navigation" aria-label="main navigation" class="navbar is-fixed-top is-spaced" style="background-color:#D2E459;">
                <div class="navbar-brand">
                    <a href="#/" class="navbar-item router-link-exact-active router-link-active">
                        <img src="https://raw.githubusercontent.com/buefy/buefy/dev/static/img/buefy-logo.png" alt="Lightweight UI components for Vue.js based on Bulma">
                    </a>
                    <a role="button" aria-label="menu" class="navbar-burger burger is-active" v-on:click="navOpen = !navOpen">
                        <span aria-hidden="true"></span> <span aria-hidden="true"></span> <span aria-hidden="true"></span>
                    </a>
                </div>
                <div v-if="navOpen" class="navbar-menu is-active" style="background-color:#D2E459;">
                    <div class="navbar-start">
                        <a href="/" class="navbar-item">
                            <a href="#/" class="router-link-exact-active router-link-active">Home</a>
                            <a href="/" class="navbar-item"><a href="#/about" class="">About</a></a>
                        </a>
                    </div>
                <div class="navbar-end"></div>
                <div v-if="!navOpen" class="navbar-menu"><div class="navbar-start"><a href="/" class="navbar-item"><a href="#/" class="router-link-exact-active router-link-active">Home</a></a><a href="/" class="navbar-item"><a href="#/about" class="">About</a></a></div><div class="navbar-end"></div></div>
            </div>
        </nav>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            navOpen: false
        }
    }
}
</script> -->
