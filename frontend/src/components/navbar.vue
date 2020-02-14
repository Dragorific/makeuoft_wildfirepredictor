<template>
    <div class="navcontainer">
        <div style="z-index:7">
            <b-navbar spaced transparent fixed-top style="background-color: #D2E459;">
                <template slot="brand">
                    <b-navbar-item tag="router-link" :to="{ path: '/' }">
                        <img
                            src="https://raw.githubusercontent.com/buefy/buefy/dev/static/img/buefy-logo.png"
                            alt="Lightweight UI components for Vue.js based on Bulma"
                        >
                    </b-navbar-item>
                </template>
                <template slot="start">
                    <b-navbar-item style="background-color: #D2E459;" href="/">
                        <router-link to="/">Home</router-link>
                    </b-navbar-item>
                    <b-navbar-item style="background-color: #D2E459;" href="/">
                        <router-link to="/">About</router-link>
                    </b-navbar-item>
                </template>
            </b-navbar>
        </div>
    </div>
</template>