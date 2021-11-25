<template>
  <div id="pageTools">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <v-flex sm12>
          <h4>SQL Format</h4>
        </v-flex>
        <div class="select-wrapper">
          <label for="lanugage">Format</label>
          <select id="language">
            <option value="sql">SQL</option>
            <!-- dialects -->
            <option value="redshift">AWS Redshift</option>
            <option value="db2">DB2</option>
            <option value="mariadb">MariaDB</option>
            <option value="mysql">MySQL</option>
            <option value="n1ql">N1QL</option>
            <option value="plsql">PL/SQL</option>
            <option value="postgresql">PostgreSQL</option>
            <option value="spark">Spark</option>
            <option value="tsql">Transact-SQL</option>
          </select>
          |
          <label for="uppercase">Uppercase</label>
          <input type="checkbox" id="uppercase" name="uppercase" checked="" />
        </div>
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
import { format } from "sql-formatter";
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
      console.log(format("SELECT * from tbl"));
      this.output = format(this.input, {
        language: "spark", // Defaults to "sql" (see the above list of supported dialects)
        indent: "    ", // Defaults to two spaces
        uppercase: true, // Defaults to false
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
