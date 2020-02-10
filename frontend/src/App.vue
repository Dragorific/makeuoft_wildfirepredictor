

<template>
  <div id="app">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
    <link href="https://fonts.googleapis.com/css?family=Odibee+Sans&display=swap" rel="stylesheet">
    <!--div id="nav">
      <router-link to="/">Home</router-link>
      <router-link to="/about">About</router-link>
      <router-link to="/helloworld">DontClickMe</router-link>
    </div-->
    <router-view/>
  </div>
</template>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}


.front-page{
  height:100vh;
}
.project-title{
  position:absolute;
  top:50%;
  left:50%;
  transform:translate(-50%, -50%);
  font-family: 'Odibee Sans', cursive;
  font-size:70px;
  color: #2F66A9;
}
.project-content{
  background-color:#74D5DD;
  padding:15px;
}
.intro{
  height:100vh;
  width:100%;
  display:flex;
  align-items: center;
  justify-content: center;
}
.row{
  width:100%;
}
.intro-desc{
  min-height:400px;
  background-color:#2F66A9;
  border: 1px solid;
  box-shadow: 5px 10px black;
}
.intro-img{
  min-height:400px;
}
.info{
  min-height:600px;
  background-color:#398AD7;
  border:1px solid;
}

p{
  color:black ;
}
</style>
