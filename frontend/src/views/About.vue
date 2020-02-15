<template>
  <div class="about">
    <navbar/>
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
    <div class="members-title">
      <p>Members</p>
    </div>
    <div class="members-container">
      <div class="columns">
        <div class="column is-half">
          <div class="group-member-container">
            <div class="member-title"><p>Faiq Ahmed</p></div>
            <div class="group-member-image-container"><img class="member-img-display" src="../assets/images/Faiq.jpg" width="60%"/></div>
            <div class="group-member-desc-container"><p>I'm faiq and I did nothing in this project</p></div>
          </div>
        </div>
        <div class="column is-half">
          <div class="group-member-container">
            <div class="member-title"><p>Umar Khan</p></div>
            <div class="group-member-image-container"><img class="member-img-display" src="../assets/images/Umar.jpg" width="60%"/></div>
            <div class="group-member-desc-container"><p>I'm umar and I did nothing in this project</p></div>
          </div>
        </div>
      </div>
      <div class="columns">
        <div class="column is-half">
          <div class="group-member-container">
            <div class="member-title"><p>Skylar Liang</p></div>
            <div class="group-member-image-container"><img class="member-img-display" src="../assets/images/Skylar.png" width="60%"/></div>
            <div class="group-member-desc-container"><p>I'm skylar and I did nothing in this project</p></div>
          </div>
        </div>
        <div class="column is-half">
          <div class="group-member-container">
            <div class="member-title"><p>Veerash Palanichamy</p></div>
            <div class="group-member-image-container"><img class="member-img-display" src="../assets/images/Veerash.jpg" width="60%"/></div>
            <div class="group-member-desc-container"><p>I'm veerash and I did nothing in this project</p></div>
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
