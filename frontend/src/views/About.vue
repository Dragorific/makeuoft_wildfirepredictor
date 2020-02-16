<template>
  <div class="about">
    <navbar/>
    <div class="paralax front-transition forest-trees">
      <p>About Us</p>
    </div>
    <div class="group-info-container">
      <div class="columns">
        <div class="column">
          <div class="group-img"><img src="https://www.legalgeek.co/wp-content/uploads/2018/02/Online-Courts-Hackathon-Team-777x440.jpg" /></div>
        </div>
        <div class="column">
          <div class="group-desc">
            <p>"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."</p>
          </div>
        </div>
      </div>
    </div>
    <div class="paralax sub-transition forest-mountains"></div>
    <div class="members-title">
      <p>Members</p>
    </div>
    <div class="members-container">

      <div class="columns">
        <div class="column is-quarter">
          <div class="card">
            <div class="card-image">
              <figure class="image is-1by1">
                <img src="../assets/images/Faiq.jpg" alt="Placeholder image" >
              </figure>
            </div>
            <div class="card-content">
              <div class="media">
                <!--div class="media-left">
                  <figure class="image is-48x48">
                    <img src="../assets/images/Faiq.jpg" alt="Placeholder image">
                  </figure>
                </div-->
                <div class="media-content">
                  <p class="title is-4">Faiq Ahmed</p>
                  <p class="subtitle is-6">@johnsmith</p>
                </div>
              </div>
              <div class="content">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                Phasellus nec iaculis mauris. <a>@bulmaio</a>.
                <a href="#">#css</a> <a href="#">#responsive</a>
                <br>
                <time datetime="2016-1-1">11:09 PM - 1 Jan 2016</time>
              </div>
            </div>
          </div>
        </div>
        <div class="column is-quarter">
          <div class="card">
            <div class="card-image">
              <figure class="image is-1by1">
                <img src="../assets/images/Umar.jpg" alt="Placeholder image" >
              </figure>
            </div>
            <div class="card-content">
              <div class="media">
                <!--div class="media-left">
                  <figure class="image is-48x48">
                    <img src="../assets/images/Umar.jpg" alt="Placeholder image">
                  </figure>
                </div-->
                <div class="media-content">
                  <p class="title is-4">Umar Khan</p>
                  <p class="subtitle is-6">@johnsmith</p>
                </div>
              </div>
              <div class="content">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                Phasellus nec iaculis mauris. <a>@bulmaio</a>.
                <a href="#">#css</a> <a href="#">#responsive</a>
                <br>
                <time datetime="2016-1-1">11:09 PM - 1 Jan 2016</time>
              </div>
            </div>
          </div>
        </div>
    
        <div class="column is-quarter">
          <div class="card">
            <div class="card-image">
              <figure class="image is-1by1">
                <img src="../assets/images/Skylar.png" alt="Placeholder image" >
              </figure>
            </div>
            <div class="card-content">
              <div class="media">
                <!--div class="media-left">
                  <figure class="image is-48x48">
                    <img src="../assets/images/Skylar.png" alt="Placeholder image">
                  </figure>
                </div-->
                <div class="media-content">
                  <p class="title is-4">Skylar Liang</p>
                  <p class="subtitle is-6">@johnsmith</p>
                </div>
              </div>
              <div class="content">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                Phasellus nec iaculis mauris. <a>@bulmaio</a>.
                <a href="#">#css</a> <a href="#">#responsive</a>
                <br>
                <time datetime="2016-1-1">11:09 PM - 1 Jan 2016</time>
              </div>
            </div>
          </div>
        </div>
        <div class="column is-quarter">
          <div class="card">
            <div class="card-image">
              <figure class="image is-1by1">
                <img src="../assets/images/Veerash.jpg" alt="Placeholder image" >
              </figure>
            </div>
            <div class="card-content">
              <div class="media">
                <!--div class="media-left">
                  <figure class="image is-48x48">
                    <img src="../assets/images/Veerash.jpg" alt="Placeholder image">
                  </figure>
                </div-->
                <div class="media-content">
                  <p class="title is-4">Veerash Palanichamy</p>
                  <p class="subtitle is-6">@johnsmith</p>
                </div>
              </div>
              <div class="content">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                Phasellus nec iaculis mauris. <a>@bulmaio</a>.
                <a href="#">#css</a> <a href="#">#responsive</a>
                <br>
                <time datetime="2016-1-1">11:09 PM - 1 Jan 2016</time>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="google-maps-container"></div>

  </div>
</template>

<script>
import navbar from '@/components/navbar.vue'

export default {
  components: {
    navbar
  }
}
</script>


<style scoped>
  @import '../assets/css/AboutPage.scss';
</style>
