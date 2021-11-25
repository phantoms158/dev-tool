<template>
  <div id="pageTools">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <v-flex sm12>
          <h4>Diff Checker</h4>
        </v-flex>
        <v-flex lg12 sm12>
          <v-textarea
            outline
            textarea
            label="Input"
            v-model="input1"
            rows="10"
          ></v-textarea>
          <v-textarea
            outline
            textarea
            label="Input"
            v-model="input2"
            rows="10"
          ></v-textarea>
        </v-flex>
        <v-flex class="text-xs-center">
          <v-btn @click="doing()" color="info">Format</v-btn>
        </v-flex>
        <v-flex lg12 sm12 :key="index" v-for="(item, index) in output">
          <div v-html="item"></div>
          <!-- <v-textarea
            v-on:focus="$event.target.select()"
            outline
            textarea
            ref="output1"
            label="Output"
            readonly
            :value="output1"
            rows="10"
          ></v-textarea>
          <v-textarea
            v-on:focus="$event.target.select()"
            outline
            textarea
            ref="output2"
            label="Output"
            readonly
            :value="output2"
            rows="10"
          ></v-textarea> -->
        </v-flex>
        <v-flex class="text-xs-center">
          <v-btn @click="copy()" color="success">Copy</v-btn>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
require("colors");
import VWidget from "@/components/VWidget";
const Diff = require("diff");

export default {
  layout: "dashboard",
  components: {
    VWidget,
  },
  data() {
    return {
      input1: "",
      input2: "",
      output: [],
    };
  },
  computed: {},
  methods: {
    doing() {
      this.output = [];
      const diff = Diff.diffChars(this.input1, this.input2);
      console.log(this.input1);
      console.log(this.input2);
      diff.forEach((part) => {
        // green for additions, red for deletions
        // grey for common parts
        const color = part.added ? "green" : part.removed ? "red" : "grey";
        this.output.push(
          '<span style="color:' + color + '">' + part.value[color] + "</span>"
        );
        console.log(part.value[color]);
      });
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
