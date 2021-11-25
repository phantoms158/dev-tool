<template>
  <div id="pageTools">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <v-flex sm12>
          <h4>JSON Decode</h4>
        </v-flex>
        <v-flex lg12 sm12>
          <v-textarea
            outline
            textarea
            label="Input"
            v-model="input"
            rows="10"
          ></v-textarea>
        </v-flex>
        <v-flex class="text-xs-center">
          <v-btn @click="doing()" color="info">Encode</v-btn>
        </v-flex>
        <v-flex lg12 sm12>
          <v-textarea
            v-on:focus="$event.target.select()"
            outline
            textarea
            ref="output" 
            label="Output"
            readonly
            :value="output"
            rows="10"
          ></v-textarea>
        </v-flex>
        <v-flex class="text-xs-center">
          <v-btn @click="copy()" color="success">Copy</v-btn>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
import VWidget from "@/components/VWidget";

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
    };
  },
  computed: {
    doing() {
      this.output = jsbeauty(this.input);
    },
    copy() {
      this.$refs.output.focus();
      document.execCommand("copy");
    },
  },
  methods: {},
};
</script>
<style lang="stylus" scoped>
.mt-45
  margin-top: -45px

.mt-56
  margin-top: -56px
</style>
