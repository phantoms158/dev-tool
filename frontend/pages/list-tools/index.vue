<template>
  <div id="pageTools">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <v-flex :key="index" lg3 sm12 v-for="(item, index) in services">
          <v-card v-bind:to=item.link>
            <v-card-title class="pb-0">
              <span class="material-icons">
                {{ item.icon }}
              </span>
              <h4>{{ item.title }}</h4>
            </v-card-title>
            <v-card-text>
              {{ item.description }}
            </v-card-text>
            <v-card-text class="d-flex justify-center">
              <!-- <v-btn flat small>Link</v-btn> -->
              <button type="button" class="v-btn v-btn--outline v-btn--depressed theme--light info--text" >
                <div class="v-btn__content">Use</div>
              </button>
            </v-card-text>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
import VWidget from "@/components/VWidget";
import { getServices } from "@/api/service";
export default {
  layout: "dashboard",
  components: {
    VWidget,
  },
  data() {
    return {
      show: true,
      cardText:
        "Hey there, I am a very simple card. I am good at containing small bits of information. I am quite convenient because I require little markup to use effectively.",
      services: getServices(),
    };
  },
  async fetch() {
    this.posts = await this.$http.$get("https://api.publicapis.org/entries");
    console.log(this.services);
  },
  computed: {},
  methods: {},
};
</script>
<style lang="stylus" scoped>
.mt-45
  margin-top: -45px

.mt-56
  margin-top: -56px
</style>
