
<template>
  <div id="app">
    <link rel="stylesheet" href="https://cdn.materialdesignicons.com/2.5.94/css/materialdesignicons.min.css">
    <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.2.0/css/all.css">
    <link href="https://fonts.googleapis.com/css?family=Odibee+Sans&display=swap" rel="stylesheet">
    <div id="nav">
      <router-link to="/">Project</router-link>
      <router-link to="/about">About</router-link>
    </div>
    <router-view/>
  </div>
</template>

<style>
    @import './assets/css/AboutPage.css';
    @import './assets/css/ProjectPage.css';
    @import './assets/css/HomePage.css';
</style>

