<template>
  <div id="pageTools">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <v-flex sm12>
          <h4>HTML Format</h4>
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
          <v-btn @click="doing()" color="info">Format</v-btn>
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
var beautify_html = require('js-beautify').html
import VWidget from "@/components/VWidget";

export default {
  layout: "dashboard",
  components: {
    VWidget,
  },
  data() {
    return {
      input: "",
      output: "",
    };
  },
  computed: {},
  methods: {
    doing() {
      this.output = beautify_html(this.input);
    },
    copy() {
      this.$refs.output.focus();
      document.execCommand("copy");
    },
  },
};
</script>
<style lang="stylus" scoped>
.mt-45
  margin-top: -45px

.mt-56
  margin-top: -56px
</style>
